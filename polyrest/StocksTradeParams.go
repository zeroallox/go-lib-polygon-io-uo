package polyrest

import (
	"github.com/valyala/fasthttp"
	"strconv"
)

type StocksTradeParams struct {
	Ticker         string
	Date           string
	Timestamp      int64
	TimestampLimit int64
	Reverse        bool
	Limit          uint
}

func (p *StocksTradeParams) configureRequest(req *fasthttp.Request) error {

	var args = req.URI().QueryArgs()
	args.Set("ticker", p.Ticker)
	args.Set("date", p.Date)
	args.SetUint("timestamp", int(p.Timestamp))
	args.SetUint("timestampLimit", int(p.TimestampLimit))
	args.Set("reverse", strconv.FormatBool(p.Reverse))

	if p.Limit != 0 {
		// Max results is 50,000 according to docs
		args.SetUint("limit", int(p.Limit))
	}

	return nil
}
