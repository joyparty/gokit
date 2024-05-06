package gokit

import (
	"sync"
	"testing"
)

func TestPoolOf(t *testing.T) {
	pool := NewPoolOf[int](nil)

	// Test Get and Put
	value := pool.Get()
	pool.Put(42)

	// Test concurrency
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value := pool.Get()
			pool.Put(value)
		}()
	}
	wg.Wait()

	// Test custom factory function
	pool = NewPoolOf(func() int {
		return 0
	})
	value = pool.Get()
	if value != 0 {
		t.Errorf("Expected value to be 0, got %d", value)
	}
}
