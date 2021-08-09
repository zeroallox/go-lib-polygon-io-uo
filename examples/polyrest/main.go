package main

import (
	"github.com/zeroallox/go-lib-polygon-io-uo/polyconst"
	"github.com/zeroallox/go-lib-polygon-io-uo/polyrest"
	"github.com/zeroallox/go-lib-polygon-io-uo/polyutils"
	"log"
	"os"
)

func main() {

	var apiKey = os.Getenv("POLY_API_KEY")
	//
	//tickers, res, err := polyrest.GetAllReferenceTickers(apiKey)
	//if err != nil {
	//	log.Println(res)
	//	panic(err)
	//}
	//
	//for _, cTicker := range tickers {
	//	log.Println(cTicker)
	//}
	//
	//var p polyrest.StocksTradeParams
	//p.Ticker = "AAPL"
	//p.Date = "2020-10-14"
	//p.Limit = 1

	polyrest.EnableDebug()

	var ds = "2021-01-02"

	tm, err := polyutils.TimeFromStringDate(ds, polyconst.NYCTime)

	trades, ar, err := polyrest.GetAllStockTrades(apiKey, tm, "999")
	if err != nil {
		log.Println("Err", err)
		if ar != nil {
			log.Println("APIResponse", ar.String())
		}
	}

	log.Println("Trades", len(trades))

	for _, cTrade := range trades {
		continue
		log.Println(cTrade)
	}

}
