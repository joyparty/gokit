//go:build go1.20
// +build go1.20

package gokit

// CompareAndDelete 比较并删除
func (mo *MapOf[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return mo.values.CompareAndDelete(key, old)
}

// CompareAndSwap 比较并替换
func (mo *MapOf[K, V]) CompareAndSwap(key K, old, new V) bool {
	return mo.values.CompareAndSwap(key, old, new)
}

// Swap 替换值
func (mo *MapOf[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	if v, ok := mo.values.Swap(key, value); ok {
		return v.(V), true
	}
	return
}
