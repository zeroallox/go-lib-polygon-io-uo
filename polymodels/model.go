package polymodels

import "sync"

type model interface {
    register(self any, owner *sync.Pool)
    setReferenceCount(rc uint64)
}
