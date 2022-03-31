package polyrest

type ReferenceTickerParams struct {
	Ticker    string `url:"ticker,omitempty"`
	TickerLT  string `url:"ticker.lt,omitempty"`
	TickerLTE string `url:"ticker.lte,omitempty"`
	TickerGT  string `url:"ticker.gt,omitempty"`
	TickerGTE string `url:"ticker.gte,omitempty"`
	Type      string `url:"type,omitempty"`
	Market    string `url:"market,omitempty"`
	Exchange  string `url:"exchange,omitempty"`
	CUSIP     string `url:"cusip,omitempty"`
	CIK       string `url:"cik,omitempty"`
	Date      string `url:"date,omitempty"`
	Search    string `url:"search,omitempty"`
	Active    bool   `url:"active,omitempty"`
	Sort      string `url:"sort,omitempty"`
	Limit     uint   `url:"limit,omitempty"`
}
