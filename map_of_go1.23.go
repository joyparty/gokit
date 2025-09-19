//go:build go1.23
// +build go1.23

package gokit

// Clear 清空
func (mo *MapOf[K, V]) Clear() {
	mo.values.Clear()
}
