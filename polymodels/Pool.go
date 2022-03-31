package polymodels

import (
    "sync"
)

type Pool[T any] struct {
    pool sync.Pool
}

func (this *Pool[T]) Release(item ...T) {
    for _, cItem := range item {
        this.pool.Put(cItem)
    }
}

func (this *Pool[T]) Acquire() T {
    return this.pool.Get().(T)
}
