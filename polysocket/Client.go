package polysocket

import (
    "bytes"
    "context"
    "errors"
    jsoniter "github.com/json-iterator/go"
    log "github.com/sirupsen/logrus"
    "github.com/zeroallox/go-lib-polygon-io-uo/polymodels"
    "github.com/zeroallox/go-lib-polygon-io-uo/polysocket/internal"
    "io"
    "nhooyr.io/websocket"
    "sync"
    "time"
)

const readBufferSize = 65536
const defaultConnectionInterval = time.Second * 5

type Client struct {
    ws            *websocket.Conn
    opt           Options
    autoReconnect bool
    mtx           sync.Mutex
    msgQueue      []jsoniter.RawMessage
    cond          sync.Cond
    stop          bool
    state         State

    onStateChanged OnConnectionStateChangedFunc
    onDataReceived OnDataReceivedFunc
}

type OnDataReceivedFunc func(item any)

func (cli *Client) SetOnDataReceivedHandler(onDataReceived OnDataReceivedFunc) {
    cli.onDataReceived = onDataReceived
}

// NewClient initializes a new Client configured with Options.
func NewClient(options Options) (*Client, error) {

    var n = new(Client)
    n.cond = sync.Cond{L: &sync.Mutex{}}
    n.state = STUnconnected

    var err = validateOptions(&options)
    if err != nil {
        return nil, err
    }

    n.opt = options

    go n.beginProcessMessage()

    return n, nil
}

type OnConnectionStateChangedFunc func(state State)

// State returns the Clients current State
func (cli *Client) State() State {
    cli.mtx.Lock()
    defer cli.mtx.Unlock()

    return cli.state
}

// setState sets the Clients current State and calls
// the onStateChangedHandler if one was set by the user.
func (cli *Client) setState(cs State) {
    cli.mtx.Lock()
    {
        if cli.state == cs {
            cli.mtx.Unlock()
            return
        }
        cli.state = cs
    }
    cli.mtx.Unlock()

    if cli.onStateChanged != nil {
        cli.onStateChanged(cs)
    }
}

func (cli *Client) SetOnStateChangedHandler(onStateChanged OnConnectionStateChangedFunc) {
    cli.onStateChanged = onStateChanged
}

// Connect opens a connection to the websocket server.
func (cli *Client) Connect() error {

    if cli.State() == STClosed {
        return ErrClientClosed
    }

    var err = cli.closeWSConnection()
    if err != nil {
        return err
    }

    cli.setState(STConnecting)
    return cli.connect()
}

// connect opens the raw websocket connection and starts the
// message read thread.
func (cli *Client) connect() error {

    ctx, cancel := context.WithTimeout(context.Background(), cli.opt.AutoReconnectInterval)
    defer cancel()

    var ep = cli.opt.ClusterType.endpoint()

    conn, res, err := websocket.Dial(ctx, ep, nil)
    if err != nil {
        return err
    }

    conn.SetReadLimit(readBufferSize)

    if res.StatusCode != 101 { // Switching Protocols
        return errors.New("not switching protocols")
    }

    cli.ws = conn
    cli.msgQueue = cli.msgQueue[:0]

    go cli.beginReading()
    go cli.beginPinging()

    return nil
}

// Disconnect gracefully closes the websocket connection but does
// NOT kill the message dispatch thread. This is useful if you'd
// like to reuse the client later and simply want to stop receiving
// messages.
//
// Note: AutoReconnect will be disabled.
func (cli *Client) Disconnect() error {
    cli.autoReconnect = false
    cli.setState(STDisconnected)

    return cli.closeWSConnection()
}

// Close gracefully closes the underlying websocket connection AND kills
// the message processor thread. Call this when you are completely done
// with the Client. Closed clients may not be reused.
func (cli *Client) Close() error {
    cli.shutdownProcessor()
    cli.setState(STClosed)

    return cli.closeWSConnection()
}

// Subscribe subscribes to the specified topic and symbols.
func (cli *Client) Subscribe(topic Topic, symbols ...string) error {
    return cli.modSubscription(internal.SASubscribe, topic, symbols...)
}

// Unsubscribe unsubscribes from the specified topic and symbols.
func (cli *Client) Unsubscribe(topic Topic, symbols ...string) error {
    return cli.modSubscription(internal.SAUnsubscribe, topic, symbols...)
}

// modSubscription validates the clients state, topic, action, and sends the
// resulting message to the server.
func (cli *Client) modSubscription(action internal.SubAction, topic Topic, symbols ...string) error {

    if cli.State() != STReady {
        return ErrClientNotReady
    }

    if cli.opt.ClusterType.supportsTopic(topic) == false {
        return ErrUnsupportedTopic
    }

    if len(symbols) == 0 {
        return ErrNoSymbols
    }

    var msg internal.Message
    var err = internal.ConfigureModSubMessage(&msg, action, topic.subscriptionPrefix(), symbols)
    if err != nil {
        return err
    }

    jData, err := json.Marshal(&msg)
    if err != nil {
        return err
    }

    return cli.writeMessage(jData)
}

// Gracefully closes the websocket connection.
func (cli *Client) closeWSConnection() error {
    cli.mtx.Lock()
    defer cli.mtx.Unlock()

    if cli.ws == nil {
        return nil
    }

    return cli.ws.Close(websocket.StatusNormalClosure, "")
}

// shutdownProcessor wakes up and kills the processing thread.
func (cli *Client) shutdownProcessor() {
    cli.cond.L.Lock()
    defer cli.cond.L.Unlock()

    cli.stop = true
    cli.cond.Signal()
}

func (cli *Client) isStopped() bool {
    cli.cond.L.Lock()
    defer cli.cond.L.Unlock()

    return cli.stop
}

// beginReading reads from the websocket and submits messages
// to the processing queue.
func (cli *Client) beginReading() {

    defer func() {
        if cli.autoReconnect == true {
            time.Sleep(cli.opt.AutoReconnectInterval)
            go cli.reconnect()
        }
    }()

    var buff [readBufferSize]byte

    var ctx = context.Background()

    for {

        msgType, reader, err := cli.ws.Reader(ctx)
        if err != nil {
            log.WithError(err).Error("get reader failed")
            return
        }

        if msgType != websocket.MessageText {
            err = errors.New("expected text message")
            return
        }

        var cursor = 0
        for {
            bytesRead, err := reader.Read(buff[cursor : readBufferSize-cursor])
            if err != nil {
                if err == io.EOF {
                    break
                }
                log.WithError(err).Error("reader error")
                return
            }

            cursor = cursor + bytesRead
        }

        var data = buff[0:cursor]

        var msgs []jsoniter.RawMessage
        err = json.Unmarshal(data, &msgs)
        if err != nil {
            //log.Println(string(data))
            //log.WithError(err).Error("unmarshal failed")
            //log.Println("LEN", len(data))
            return
        }

        cli.cond.L.Lock()
        cli.msgQueue = append(cli.msgQueue, msgs...)
        cli.cond.L.Unlock()

        cli.cond.Signal()
    }

}

func (cli *Client) beginPinging() {

    for {

        if cli.isStopped() == true {
            return
        }

        var err = cli.ws.Ping(context.Background())
        if err != nil {
            log.WithError(err).Debug("pinging failed")
            return
        }

        time.Sleep(3 * time.Second)
    }
}

// reconnect starts the reconnection cycle.
func (cli *Client) reconnect() {
    cli.setState(STReconnecting)

    var err = cli.closeWSConnection()
    if err != nil {
        cli.setState(STError)
    }

    err = cli.connect()
    if err != nil {
        cli.setState(STError)
    }
}

// beginProcessMessage handles the decoding and dispatching of
// received messages.
func (cli *Client) beginProcessMessage() {

    log.Debug("Client: Start Processing Thread")

    defer func() {
        log.Debug("Client: Exit Processing Thread")
    }()

    var localQueue []jsoniter.RawMessage

    for {

        cli.cond.L.Lock()
        {
            for len(cli.msgQueue) == 0 {

                if cli.stop == true {
                    cli.cond.L.Unlock()
                    return
                }

                cli.cond.Wait()
            }

            localQueue = append(localQueue, cli.msgQueue...)
            cli.msgQueue = cli.msgQueue[:0]
        }
        cli.cond.L.Unlock()

        cli.handleMessages(localQueue)

        localQueue = localQueue[:0]
    }
}

var eqQuotePrefix = []byte("{\"ev\":\"Q\"")
var eqTradePrefix = []byte("{\"ev\":\"T\"")
var statusPrefix = []byte("{\"ev\":\"status\"")

// handleMessages routes each message to the correct handler. We specifically
// do not break and return on error as it is possible for a single message
// to be malformed.
func (cli *Client) handleMessages(msgs []jsoniter.RawMessage) {

    var err error

    for _, cMessage := range msgs {

        switch {
        case bytes.HasPrefix(cMessage, eqTradePrefix):
            err = cli.handleLiveEquityTrade(cMessage)
            break
        case bytes.HasPrefix(cMessage, eqQuotePrefix):
            err = cli.handleLiveEquityQuote(cMessage)
            break
        case bytes.HasPrefix(cMessage, statusPrefix):
            err = cli.handleStatusMessage(cMessage)
            break
        }

        if err != nil {
            log.WithError(err).Error(string(cMessage))
        }
    }
}

func (cli *Client) handleStatusMessage(msg jsoniter.RawMessage) error {

    var sm internal.Message

    var err = json.Unmarshal(msg, &sm)
    if err != nil {
        return err
    }

    switch sm.Status {
    case "connected":
        cli.setState(STConnected)
        return cli.sendAuthRequest()
    case "auth_success":
        cli.setState(STReady)
        return nil
    case "auth_failed":
        cli.setState(STError)
        _ = cli.closeWSConnection()
        return ErrAuthenticationFailed
    case "success":
        // For sub confirmations.
        // AFAIK there can never be a failure on subscribe.
        // Server does not validate ticker symbols.
        return nil
    default:
        cli.setState(STError)
        _ = cli.closeWSConnection()
        return ErrUnhandledStatusMessage
    }
}

// handleLiveEquityQuote decodes and dispatches Quote messages
func (cli *Client) handleLiveEquityQuote(msg jsoniter.RawMessage) error {

    if cli.onDataReceived == nil {
        return nil
    }

    var quote = polymodels.DefaultLiveEquityQuotePool.Acquire()
    var err = json.Unmarshal(msg, &quote)
    if err != nil {
        return err
    }

    cli.onDataReceived(quote)

    return nil
}

// handleLiveEquityTrade decodes and dispatches Trade messages
func (cli *Client) handleLiveEquityTrade(msg jsoniter.RawMessage) error {

    if cli.onDataReceived == nil {
        return nil
    }

    var trade = polymodels.DefaultLiveEquityTradePool.Acquire()
    var err = json.Unmarshal(msg, &trade)
    if err != nil {
        return err
    }

    cli.onDataReceived(trade)

    return nil
}

// sendAuthRequest creates an auth message and sends to the websocket server
func (cli *Client) sendAuthRequest() error {

    var msg = internal.MakeAuthMessage(cli.opt.APIKey)

    jData, err := json.Marshal(&msg)
    if err != nil {
        return err
    }

    return cli.writeMessage(jData)
}

// writeMessage sends the message data to the websocket server.
// A mutex is used to prevent concurrent writes.
func (cli *Client) writeMessage(data []byte) error {
    cli.mtx.Lock()
    defer cli.mtx.Unlock()

    if cli.state != STReady && cli.state != STConnected {
        return ErrClientNotReady
    }

    return cli.ws.Write(context.Background(), websocket.MessageText, data)
}
