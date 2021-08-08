package polyrest

import (
	"github.com/valyala/fasthttp"
	"strconv"
)

type ReferenceTickerParams struct {
	Ticker    string
	TickerLT  string
	TickerLTE string
	TickerGT  string
	TickerGTE string
	Type      string
	Market    string
	Exchange  string
	CUSIP     string
	CIK       string
	Date      string
	Search    string
	Active    bool
	Sort      string
	Limit     uint
}

func (p *ReferenceTickerParams) configureRequest(req *fasthttp.Request) error {

	var args = req.URI().QueryArgs()
	args.Set("ticker", p.Ticker)
	args.Set("ticker.lt", p.TickerLT)
	args.Set("ticker.lte", p.TickerLTE)
	args.Set("ticker.gt", p.TickerGT)
	args.Set("ticker.gte", p.TickerGTE)
	args.Set("type", p.Type)
	args.Set("market", p.Market)
	args.Set("exchange", p.Exchange)
	args.Set("cusip", p.CUSIP)
	args.Set("cik", p.CIK)
	args.Set("date", p.Date)
	args.Set("search", p.Search)
	args.Set("active", strconv.FormatBool(p.Active))
	args.Set("sort", p.Sort)

	if p.Limit != 0 {
		// Max results is 1,000 according to docs
		args.SetUint("limit", int(p.Limit))
	}

	return nil
}
