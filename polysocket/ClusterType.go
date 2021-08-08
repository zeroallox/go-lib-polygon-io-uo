package polysocket

type ClusterType uint8

const (
	CTInvalid       ClusterType = 0
	CTStocks        ClusterType = 1
	CTStocksDelayed ClusterType = 2
	CTForex         ClusterType = 3
	CTCrypto        ClusterType = 4
)

func (ct ClusterType) endpoint() string {
	switch ct {
	case CTStocks:
		return "wss://socket.polygon.io/stocks"
	case CTStocksDelayed:
		return "wss://delayed.polygon.io/stocks"
	case CTForex:
		return "wss://socket.polygon.io/forex"
	case CTCrypto:
		return "wss://socket.polygon.io/crypto"
	default:
		panic("should never happen")
	}
}

func (ct ClusterType) supportsTopic(topic Topic) bool {

	switch ct {
	case CTStocks, CTStocksDelayed:
		if topic > tpEquityMin && topic < tpEquityMax {
			return true
		}
	case CTForex:
		if topic > tpForexMin && topic < tpForexMax {
			return true
		}
	case CTCrypto:
		if topic > tpCryptoMin && topic < tpCryptoMax {
			return true
		}
	}

	return false
}
