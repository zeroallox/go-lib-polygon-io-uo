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
	polyrest.EnableAutoRetry()

	var ds = "2021-01-04"
	tm, err := polyutils.TimeFromStringDate(ds, polyconst.NYCTime)
	if err != nil {
		panic(err)
	}

	trades, ar, err := polyrest.GetAllStockTrades(apiKey, tm, "AAPL")
	if err != nil {
		log.Println("lib-level error", err)
		return
	}

	if ar != nil && ar.IsError() == true {
		log.Println("api-level error", ar.Error())
	}

	log.Println("Got Trades:", len(trades))
	for _, cTrade := range trades {
		log.Println(cTrade)
	}

}
