package main

import (
	"github.com/zeroallox/go-lib-polygon-io-uo/polyconst"
	"github.com/zeroallox/go-lib-polygon-io-uo/polymodels"
	"github.com/zeroallox/go-lib-polygon-io-uo/polypsv"
	"github.com/zeroallox/go-lib-polygon-io-uo/polyrest"
	"github.com/zeroallox/go-lib-polygon-io-uo/polyutils"
	"log"
	"os"
)

// This example downloads all trades for AAPL and outputs a flat file.

func main() {

	// Get your Polygon.io API key from envar
	var apiKey = os.Getenv("POLY_API_KEY")

	// Get the scratch / working dir from envar
	var scratchDir = os.Getenv("SCRATCH_DIR")

	// Make sure we actually got them :)
	if len(apiKey) == 0 || len(scratchDir) == 0 {
		panic("need to set the envars!")
	}

	// Get "2020-10-14" as a time.Time in NYC's time zone
	tm, err := polyutils.TimeFromStringDate("2020-10-14", polyconst.NYCTime)
	if err != nil {
		panic(err)
	}

	// Initialize a PSV file with the correct params for the data
	// we'll be writing.
	var psvFile = polypsv.NewFileInfo(polyconst.LOC_USA, polyconst.MKT_Stocks, polyconst.DT_Trades, tm)

	// Create a local file in the scratch dir.
	hFile, err := polypsv.CreateLocalPSVFile(scratchDir, psvFile, true)
	if err != nil {
		panic(err)
	}
	defer hFile.Close()

	// Create a new PSV Writer
	writer, err := polypsv.NewWriter(psvFile, hFile, true)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	// Generate a header for a HistoricEquityTrade
	var header = polymodels.MakePSVHeader(polymodels.HistoricEquityTrade{})

	// Set the header
	writer.SetHeader(header)

	// Download all the trades for AAPL from 2020-10-14
	trades, ar, err := polyrest.GetAllStockTrades(apiKey, tm, "GUSH")
	if err != nil {
		log.Println(ar.Error())
		panic(err)
	}

	// Write each of the trades pulled from the API to the PSV file.
	for _, cTrade := range trades {
		err = writer.WriteObject(cTrade)
		if err != nil {
			panic(err)
		}
	}
}
