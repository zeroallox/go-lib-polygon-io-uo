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

	//tickers, err := polyrest.GetAllReferenceTickers(apiKey)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, cTicker := range tickers {
	//	log.Println(cTicker)
	//}

	var p polyrest.StocksTradeParams
	p.Ticker = "AAPL"
	p.Date = "2020-10-14"
	p.Limit = 1

	tm, err := polyutils.TimeFromStringDate("2020-10-14", polyconst.NYCTime)

	trades, rez, err := polyrest.GetAllStockTrades(apiKey, tm, "GUSH")
	if err != nil {
		log.Println(rez.String())
		panic(err)
	}

	for _, cTrade := range trades {
		continue
		log.Println(cTrade)
	}

}
