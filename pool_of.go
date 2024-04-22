package gokit

import "sync"

// PoolOf sync.Pool泛型
type PoolOf[T any] struct {
	noCopy noCopy
	pool   *sync.Pool
}

// NewPoolOf 构造函数
func NewPoolOf[T any](factory func() T) *PoolOf[T] {
	p := &sync.Pool{}
	if factory != nil {
		p.New = func() any {
			return factory()
		}
	}

	return &PoolOf[T]{pool: p}
}

// Get 从池内获取
func (p *PoolOf[T]) Get() T {
	return p.pool.Get().(T)
}

// Put 放回池内
func (p *PoolOf[T]) Put(x T) {
	p.pool.Put(x)
}

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}
