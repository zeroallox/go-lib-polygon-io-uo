package polyrest

import (
    "github.com/valyala/fasthttp"
    "github.com/zeroallox/go-lib-polygon-io-uo/polymodels"
)

// GetAllReferenceTickers returns ALL tickers both ACTIVE and INACTIVE.
//A *Response will only be returned if one of the calls made resulted in an API error.
func GetAllReferenceTickers(apiKey string) ([]*polymodels.Ticker, *Response, error) {

    var params ReferenceTickerParams
    params.Limit = 1000
    params.Active = false

    var allTickers []*polymodels.Ticker

    for {

        tickers, rez, err := GetReferenceTickers(apiKey, &params)
        if err != nil {
            return nil, rez, err
        }

        allTickers = append(allTickers, tickers...)

        if len(tickers) < int(params.Limit) {

            // If the number of results returned is less than the limit
            // we got all the results.
            // Set Active to true and clear the GTE filter
            // and loop again starting from the first active ticker.
            if params.Active == false {
                params.Active = true
                params.TickerGTE = ""
                continue
            }

            break
        }

        var lastTicker = tickers[len(tickers)-1]
        params.TickerGTE = lastTicker.Ticker
    }

    // Dedupe the results
    var lookup = map[uint64]bool{}
    var final []*polymodels.Ticker

    for _, cTicker := range allTickers {
        var hash = cTicker.Hash()
        if lookup[hash] == true {
            continue
        }

        lookup[hash] = true
        final = append(final, cTicker)
    }

    return final, nil, nil
}

// GetReferenceTickers fetches tickers based on ReferenceTickerParams
func GetReferenceTickers(apiKey string, params *ReferenceTickerParams) ([]*polymodels.Ticker, *Response, error) {

    var req = fasthttp.AcquireRequest()
    defer fasthttp.ReleaseRequest(req)

    const uri = "https://api.polygon.io/v3/reference/tickers"
    var err = buildRequest(req, uri, params)
    if err != nil {
        return nil, nil, err
    }

    var resp = fasthttp.AcquireResponse()
    defer fasthttp.ReleaseResponse(resp)

    var tickers []*polymodels.Ticker
    ar, err := do(apiKey, req, resp, false, &tickers)
    if err != nil {
        return nil, ar, err
    }

    return tickers, ar, nil
}
