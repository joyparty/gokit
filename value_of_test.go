package gokit

import (
	"testing"
)

func TestValueOf(t *testing.T) {
	type TestData struct {
		Value    int
		Expected int
	}

	// Create a new ValueOf instance
	vo := NewValueOf[int]()

	// Test Load() method
	loadData := TestData{
		Value:    42,
		Expected: 42,
	}
	vo.Store(loadData.Value)
	if val := vo.Load(); val != loadData.Expected {
		t.Errorf("Load() failed, expected %d but got %d", loadData.Expected, val)
	}

	// Test Swap() method
	swapData := TestData{
		Value:    77,
		Expected: 77,
	}
	oldVal := vo.Swap(swapData.Value)
	if oldVal != loadData.Expected {
		t.Errorf("Swap() failed, expected old value %d but got %d", loadData.Expected, oldVal)
	}
	if val := vo.Load(); val != swapData.Expected {
		t.Errorf("Swap() failed, expected new value %d but got %d", swapData.Expected, val)
	}

	// Test CompareAndSwap() method
	compareAndSwapData := TestData{
		Value:    55,
		Expected: 55,
	}
	swapped := vo.CompareAndSwap(swapData.Expected, compareAndSwapData.Value)
	if !swapped {
		t.Errorf("CompareAndSwap() failed, expected swap to be successful")
	}
	if val := vo.Load(); val != compareAndSwapData.Expected {
		t.Errorf("CompareAndSwap() failed, expected new value %d but got %d", compareAndSwapData.Expected, val)
	}
}
