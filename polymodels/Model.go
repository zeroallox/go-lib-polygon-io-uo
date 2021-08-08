package polymodels

import "sync"

type model struct {
	owner *sync.Pool
}

func (m *model) Release() {
	if m.owner != nil {
		m.owner.Put(m)
	}
}
