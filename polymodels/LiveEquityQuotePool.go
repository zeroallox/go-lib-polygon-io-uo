package polymodels

import "sync"

type LiveEquityQuotePool struct {
	p *sync.Pool
}

var DefaultLiveEquityQuotePool = newLiveEquityQuotePool()

func newLiveEquityQuotePool() *LiveEquityQuotePool {
	var n = new(LiveEquityQuotePool)
	n.p = new(sync.Pool)

	n.p.New = func() interface{} {
		var q = new(LiveEquityQuote)
		q.model.owner = n.p
		return q
	}

	return n
}

func (qp *LiveEquityQuotePool) Acquire() *LiveEquityQuote {
	return qp.p.Get().(*LiveEquityQuote)
}
