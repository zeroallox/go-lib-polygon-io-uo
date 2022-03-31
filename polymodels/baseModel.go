package polymodels

import (
    "sync"
    "sync/atomic"
)

const tagSize = 2

type baseModel struct {
    self     any
    owner    *sync.Pool
    refCount int64
    tags     [tagSize]int64
}

func (m *baseModel) SetTag(idx uint8, value int64) {
    if idx > tagSize-1 {
        panic("bad tag index")
    }
    atomic.StoreInt64(&m.tags[idx], value)
}

func (m *baseModel) GetTag(idx uint8) int64 {
    if idx > tagSize-1 {
        panic("bad tag index")
    }
    return atomic.LoadInt64(&m.tags[idx])
}

func (m *baseModel) AutoRelease() {
    switch atomic.AddInt64(&m.refCount, -1) {
    case 0:
        m.owner.Put(m.self)
        return
    case -1:
        panic("baseModel: AutoRelease called when model not AutoAcquired")
    }
}

func (m *baseModel) Release() {
    m.refCount = 0
    m.owner.Put(m.self)
}

func (m *baseModel) register(self any, owner *sync.Pool) {
    m.owner = owner
    m.self = self
}

func (m *baseModel) setReferenceCount(rc uint64) {
    atomic.StoreInt64(&m.refCount, int64(rc))
}
