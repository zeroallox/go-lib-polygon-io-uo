package polyrest

import (
	"github.com/valyala/fasthttp"
	"github.com/zeroallox/go-lib-polygon-io-uo/polymodels"
)

// GetAllReferenceTickers returns ALL tickers both ACTIVE and INACTIVE.
//A *APIResponse will only be returned if one of the calls made resulted in an API error.
func GetAllReferenceTickers(apiKey string) ([]*polymodels.Ticker, *APIResponse, error) {

	var req = fasthttp.AcquireRequest()
	req.SetRequestURI("https://api.polygon.io/v3/reference/tickers")

	defer fasthttp.ReleaseRequest(req)

	var params ReferenceTickerParams
	params.Limit = 1000
	params.Active = false

	var allTickers []*polymodels.Ticker

	for {

		var err = params.configureRequest(req)
		if err != nil {
			return nil, nil, err
		}

		tickers, rez, err := getReferenceTickers(apiKey, req)
		if err != nil {
			return nil, rez, err
		}

		allTickers = append(allTickers, tickers...)

		if len(tickers) < int(params.Limit) {

			// First we got all the inactive tickers.
			// Now we request all the active ones starting
			// from zero.
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
func GetReferenceTickers(apiKey string, params *ReferenceTickerParams) ([]*polymodels.Ticker, *APIResponse, error) {
	var req = fasthttp.AcquireRequest()
	req.SetRequestURI("https://api.polygon.io/v3/reference/tickers")

	defer fasthttp.ReleaseRequest(req)

	var err = params.configureRequest(req)
	if err != nil {
		return nil, nil, err
	}

	return getReferenceTickers(apiKey, req)
}

func getReferenceTickers(apiKey string, req *fasthttp.Request) ([]*polymodels.Ticker, *APIResponse, error) {

	var resp = fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	var tickers []*polymodels.Ticker
	ar, err := do(apiKey, req, resp, &tickers, false)
	if err != nil {
		return nil, ar, err
	}

	return tickers, ar, nil
}
