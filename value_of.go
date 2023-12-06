package gokit

import "sync/atomic"

// ValueOf 泛型实现的atomic.Value
//
// 如果值是指针类型，应该使用标准库提供的atomic.Pointer
type ValueOf[T any] interface {
	Load() T
	Store(val T)
	Swap(new T) (old T)
	CompareAndSwap(old, new T) (swapped bool)
}

type valueOf[T any] struct {
	value *atomic.Value
}

// NewValueOf 构造泛型atomic.Value
func NewValueOf[T any]() ValueOf[T] {
	v := &valueOf[T]{
		value: &atomic.Value{},
	}

	// atomic.Value的初始值为nil
	// 但实际上使用者会期望这个初始值为T的零值
	// 所以初始化直接把T零值放入，避免误会
	var value T
	v.value.Store(value)
	return v
}

// CompareAndSwap 比较并替换
func (vo *valueOf[T]) CompareAndSwap(old, new T) (swapped bool) {
	return vo.value.CompareAndSwap(old, new)
}

// Load 获取保存的值
func (vo *valueOf[T]) Load() T {
	if v, ok := vo.value.Load().(T); ok {
		return v
	}

	var result T
	return result
}

// Store 保存值
func (vo *valueOf[T]) Store(val T) {
	vo.value.Store(val)
}

// Swap 替换值
func (vo *valueOf[T]) Swap(new T) (old T) {
	if v, ok := vo.value.Swap(new).(T); ok {
		return v
	}

	var result T
	return result
}
