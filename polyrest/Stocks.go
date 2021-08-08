package polyrest

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/zeroallox/go-lib-polygon-io-uo/polymodels"
	"github.com/zeroallox/go-lib-polygon-io-uo/polyutils"
	"time"
)

// GetAllStockTrades gets all historic stock trades for the given date and ticker.
// A *Result will only be returned if one of the calls made resulted in an API error.
func GetAllStockTrades(apiKey string, date time.Time, ticker string) ([]*polymodels.HistoricEquityTrade, *Result, error) {

	var req = fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	if date.IsZero() == true {
		return nil, nil, ErrInvalidDateParam
	}

	var ds = polyutils.TimeToStringDate(date)
	var uri = fmt.Sprintf("https://api.polygon.io/v2/ticks/stocks/trades/%v/%v", ticker, ds)
	req.SetRequestURI(uri)

	var params StocksTradeParams
	params.Limit = 10000

	var allTrades []*polymodels.HistoricEquityTrade

	for {

		var err = params.configureRequest(req)
		if err != nil {
			return nil, nil, err
		}

		cTrades, rez, err := getStockTrades(apiKey, req, ticker)
		if err != nil {
			return nil, rez, err
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
func GetStockTrades(apiKey string, params *StocksTradeParams) ([]*polymodels.HistoricEquityTrade, *Result, error) {
	var req = fasthttp.AcquireRequest()

	var uri = fmt.Sprintf("https://api.polygon.io/v2/ticks/stocks/trades/%v/%v", params.Ticker, params.Date)
	req.SetRequestURI(uri)

	defer fasthttp.ReleaseRequest(req)

	var err = params.configureRequest(req)
	if err != nil {
		return nil, nil, err
	}

	return getStockTrades(apiKey, req, params.Ticker)
}

func getStockTrades(apiKey string, req *fasthttp.Request, ticker string) ([]*polymodels.HistoricEquityTrade, *Result, error) {

	var resp = fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	rez, err := do(apiKey, req, resp)
	if err != nil {
		return nil, rez, err
	}

	var trades []*polymodels.HistoricEquityTrade
	err = json.Unmarshal(rez.takeResultData(), &trades)
	if err != nil {
		return nil, rez, err
	}

	// Set the ticker on each trade as it's not returned
	// as part of the model json. It is returned as part of
	// Result but I have a feeling this will change.
	for _, cTrade := range trades {
		cTrade.Ticker = ticker
	}

	return trades, rez, nil
}
