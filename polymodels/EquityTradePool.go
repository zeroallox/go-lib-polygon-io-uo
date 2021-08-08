package polymodels
//
//import "sync"
//
//type EquityTradePool struct {
//	pool sync.Pool
//}
//
//var DefaultEquityTradePool = newEquityTradePool()
//
//func newEquityTradePool() *EquityTradePool {
//	var n = new(EquityTradePool)
//	n.pool.New = func() interface{} {
//		var trade = new(HistoricEquityTrade)
//		trade.model.owner = &n.pool
//		return trade
//	}
//
//	return n
//}
