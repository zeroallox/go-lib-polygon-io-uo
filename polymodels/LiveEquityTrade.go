package polymodels

type LiveEquityTrade struct {
    baseModel

    Ticker     string  `json:"sym"`
    TradeID    string  `json:"i"`
    Exchange   int     `json:"x"`
    Price      float64 `json:"p"`
    Volume     int     `json:"s"`
    Conditions []int   `json:"c"`
    Timestamp  int64   `json:"t"`
    Sequence   int     `json:"q"`
    Tape       int     `json:"z"`
}
