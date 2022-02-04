package polyrest

import (
    "fmt"
    "github.com/valyala/fasthttp"
    "github.com/zeroallox/go-lib-polygon-io-uo/polymodels"
    "github.com/zeroallox/go-lib-polygon-io-uo/polyutils"
    "time"
)

// GetAllStockTrades gets all historic stock trades for the given date and ticker.
// An *Response will only be returned if one of the calls made resulted in an API error.
func GetAllStockTrades(apiKey string, date time.Time, ticker string) ([]*polymodels.HistoricEquityTrade, *Response, error) {

    if date.IsZero() == true {
        return nil, nil, ErrInvalidDateParam
    }

    var params StocksTradeParams
    params.Limit = 5000
    params.Ticker = ticker
    params.Date = polyutils.TimeToStringDate(date)

    var allTrades []*polymodels.HistoricEquityTrade

    for {

        cTrades, ar, err := GetStockTrades(apiKey, &params)
        if err != nil {
            return nil, ar, err
        }

        allTrades = append(allTrades, cTrades...)

        if len(cTrades) < int(params.Limit) {
            break
        }

        var lastTrade = allTrades[len(allTrades)-1]
        params.Timestamp = lastTrade.SIPTimestamp
    }

    // Dedupe Trades
    var lookup = map[uint64]bool{}
    var trades []*polymodels.HistoricEquityTrade

    for _, cTrade := range allTrades {
        var hash = cTrade.Hash()
        if lookup[hash] == true {
            continue
        }

        lookup[hash] = true
        trades = append(trades, cTrade)
    }

    return trades, nil, nil
}

// GetStockTrades fetches historic trades based on StocksTradeParams
func GetStockTrades(apiKey string, params *StocksTradeParams) ([]*polymodels.HistoricEquityTrade, *Response, error) {

    var req = fasthttp.AcquireRequest()
    defer fasthttp.ReleaseRequest(req)

    var uri = fmt.Sprintf("https://api.polygon.io/v2/ticks/stocks/trades/%s/%s", params.Ticker, params.Date)
    var err = buildRequest(req, uri, params)
    if err != nil {
        return nil, nil, err
    }

    var resp = fasthttp.AcquireResponse()
    defer fasthttp.ReleaseResponse(resp)

    var trades []*polymodels.HistoricEquityTrade
    ar, err := do(apiKey, req, resp, true, &trades)
    if err != nil {
        return nil, ar, err
    }

    // Set the ticker on each trade.
    for _, cTrade := range trades {
        cTrade.Ticker = params.Ticker
    }

    return trades, ar, nil
}
