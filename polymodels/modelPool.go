package polymodels

type ModelPool[T model] struct {
    Pool[T]
}

type NewModelFunc[T model]func() T

func newModelPool[T model](nf NewModelFunc[T]) *ModelPool[T] {

    var n = new(ModelPool[T])
    n.pool.New = func() any {
        var o = nf()
        o.register(o, &n.pool)
        return o
    }

    return n
}

func (this *ModelPool[T]) AutoAcquire(rc uint64) T {
    var n = this.pool.Get().(T)
    n.setReferenceCount(rc)
    return n
}
