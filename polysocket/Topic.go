package polysocket

type Topic uint8

const (
	tpEquityMin        Topic = 10
	TPEquityTrades     Topic = 11
	TPEquityQuotes     Topic = 12
	TPEquityAggMin     Topic = 13
	TPEquityAggSec     Topic = 14
	TPEquityLULD       Topic = 15
	TPEquityImbalances Topic = 16
	tpEquityMax        Topic = 17

	tpForexMin    Topic = 50
	TPForexTrades Topic = 51
	TPForexAggMin Topic = 52
	tpForexMax    Topic = 53

	tpCryptoMin    Topic = 60
	TPCryptoTrades Topic = 61
	TPCryptoQuotes Topic = 62
	TPCryptoBook   Topic = 63
	TPCryptoAggMin Topic = 64
	tpCryptoMax    Topic = 65
)

func (tp Topic) subscriptionPrefix() string {
	switch tp {
	case TPEquityTrades:
		return "T"
	case TPEquityQuotes:
		return "Q"
	case TPEquityAggMin:
		return "AM"
	case TPEquityAggSec:
		return "A"
	case TPEquityLULD:
		return "LULD"
	case TPEquityImbalances:
		return "NOI"
	case TPForexTrades:
		return "C"
	case TPForexAggMin:
		return "CA"
	case TPCryptoTrades:
		return "XT"
	case TPCryptoQuotes:
		return "XQ"
	case TPCryptoBook:
		return "XL2"
	case TPCryptoAggMin:
		return "XA"
	default:
		return ""
	}
}
