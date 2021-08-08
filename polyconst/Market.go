package polyconst

import "strings"

type Market uint8

const (
	MKT_Invalid Market = 0
	MKT_Stocks  Market = 1
	MKT_Crypto  Market = 2
	MKT_Forex   Market = 3
	mkt_max     Market = 4
)

func (mkt Market) Code() string {
	switch mkt {
	case MKT_Stocks:
		return "stocks"
	case MKT_Crypto:
		return "crypto"
	case MKT_Forex:
		return "fx"
	default:
		return "_INVALID_Market_"
	}
}

func MarketFromString(str string) Market {

	str = strings.ToLower(str)

	for i := 0; i < int(mkt_max); i++ {
		var cMarket = Market(i)
		if cMarket.Code() == str {
			return cMarket
		}
	}

	return MKT_Invalid
}
