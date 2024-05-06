package gokit

import (
	"testing"
)

func TestMapOf(t *testing.T) {
	m := NewMapOf[int, string]()

	// Test LoadOrStore
	value, loaded := m.LoadOrStore(1, "one")
	if loaded {
		t.Errorf("Expected LoadOrStore to store a new value, but it loaded an existing value")
	}
	if value != "one" {
		t.Errorf("Expected LoadOrStore to return 'one', but got '%s'", value)
	}

	// Test Load
	value, found := m.Load(1)
	if !found {
		t.Errorf("Expected Load to find a value, but it didn't")
	}
	if value != "one" {
		t.Errorf("Expected Load to return 'one', but got '%s'", value)
	}

	// Test Store
	m.Store(2, "two")

	// Test LoadAndDelete
	value, loaded = m.LoadAndDelete(2)
	if !loaded {
		t.Errorf("Expected LoadAndDelete to find and delete a value, but it didn't")
	}
	if value != "two" {
		t.Errorf("Expected LoadAndDelete to return 'two', but got '%s'", value)
	}

	// Test Delete
	m.Delete(1)

	// Test Range
	m.Store(3, "three")
	m.Store(4, "four")
	count := 0
	m.Range(func(key int, value string) bool {
		count++
		return true
	})
	if count != 2 {
		t.Errorf("Expected Range to iterate over 2 items, but got %d", count)
	}

	// Test Count
	count = int(m.Count())
	if count != 2 {
		t.Errorf("Expected Count to return 2, but got %d", count)
	}
}
