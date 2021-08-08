package main

import (
	"github.com/zeroallox/go-lib-polygon-io-uo/polymodels"
	"github.com/zeroallox/go-lib-polygon-io-uo/polysocket"
	"log"
	"os"
	"sync"
	"time"
)

func main() {

	// Configure client options
	var opt polysocket.Options
	opt.APIKey = os.Getenv("POLY_API_KEY")
	opt.ClusterType = polysocket.CTStocks
	opt.AutoReconnect = true
	opt.AutoReconnectInterval = time.Second * 5

	// Initialize a new Client
	client, err := polysocket.NewClient(opt)
	if err != nil {
		panic(err)
	}

	// Set the handler to detect changes to the client state.
	// Connections, Errors, Reconnections, etc
	client.SetOnStateChangedHandler(func(state polysocket.State) {

		if state == polysocket.STReady {
			_ = client.Subscribe(polysocket.TPEquityTrades, "*")
			_ = client.Subscribe(polysocket.TPEquityQuotes, "*")
		}

		log.Println("State", state)
	})

	// Set the handler to process incoming data.
	client.SetOnDataReceivedHandler(func(data polymodels.PooledItem) {

		switch data.(type) {
		case *polymodels.LiveEquityQuote:
			var quote = data.(*polymodels.LiveEquityQuote)
			log.Println(quote)
			break
		case *polymodels.LiveEquityTrade:
			var trade = data.(*polymodels.LiveEquityTrade)
			log.Println(trade)
			break
		}

	})

	// Connect to Polygon
	err = client.Connect()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
