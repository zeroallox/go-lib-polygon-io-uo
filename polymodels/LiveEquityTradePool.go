package polymodels

import "sync"

type LiveEquityTradePool struct {
	p *sync.Pool
}

var DefaultLiveEquityTradePool = newLiveEquityTradePool()

func newLiveEquityTradePool() *LiveEquityTradePool {
	var n = new(LiveEquityTradePool)
	n.p = new(sync.Pool)

	n.p.New = func() interface{} {
		var q = new(LiveEquityTrade)
		q.model.owner = n.p
		return q
	}

	return n
}

func (qp *LiveEquityTradePool) Acquire() *LiveEquityTrade {
	return qp.p.Get().(*LiveEquityTrade)
}
