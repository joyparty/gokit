package gokit

import "sync"

// MapOf 泛型实现的sync.Map
type MapOf[K comparable, V any] struct {
	values *sync.Map
}

// NewMapOf 构造函数
func NewMapOf[K comparable, V any]() *MapOf[K, V] {
	return &MapOf[K, V]{
		values: &sync.Map{},
	}
}

// Load 根据key查询结果
func (mo *MapOf[K, V]) Load(key K) (value V, found bool) {
	if v, ok := mo.values.Load(key); ok {
		return v.(V), true
	}

	return
}

// LoadAndDelete 查询并删除
func (mo *MapOf[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	if v, ok := mo.values.LoadAndDelete(key); ok {
		return v.(V), true
	}
	return
}

// LoadOrStore 查询已经存在的数据，不存在则存储
func (mo *MapOf[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	if v, ok := mo.values.LoadOrStore(key, value); ok {
		return v.(V), true
	}

	return value, false
}

// Store 保存
func (mo *MapOf[K, V]) Store(key K, value V) {
	mo.values.Store(key, value)
}

// Delete 删除
func (mo *MapOf[K, V]) Delete(key K) {
	mo.values.Delete(key)
}

// Range 遍历
func (mo *MapOf[K, V]) Range(f func(key K, value V) bool) {
	mo.values.Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}

// ToMap to map
func (mo *MapOf[K, V]) ToMap() map[K]V {
	result := make(map[K]V)
	mo.Range(func(k K, v V) bool {
		result[k] = v
		return true
	})

	return result
}

// Count 总数
func (mo *MapOf[K, V]) Count() (count int64) {
	mo.values.Range(func(k, v any) bool {
		count++
		return true
	})
	return
}
