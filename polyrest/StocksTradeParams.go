package polyrest

type StocksTradeParams struct {
    Ticker         string `url:"-"`
    Date           string `url:"-"`
    Timestamp      int64  `url:"timestamp,omitempty"`
    TimestampLimit int64  `url:"timestampLimit,omitempty"`
    Reverse        bool   `url:"reverse,omitempty"`
    Limit          uint   `url:"limit,omitempty"`
}
