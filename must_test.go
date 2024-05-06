package gokit

import (
	"errors"
	"testing"
)

func TestMustReturn(t *testing.T) {
	// Test case 1: No error
	value1 := 42
	err1 := error(nil)
	result1 := MustReturn(value1, err1)
	if result1 != value1 {
		t.Errorf("MustReturn failed, expected: %v, got: %v", value1, result1)
	}

	// Test case 2: With error
	value2 := 0
	err2 := errors.New("some error")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustReturn did not panic")
		}
	}()
	MustReturn(value2, err2)
}

func TestMustReturn2(t *testing.T) {
	// Test case 1: No error
	value1 := 42
	value2 := "hello"
	err1 := error(nil)
	result1, result2 := MustReturn2(value1, value2, err1)
	if result1 != value1 || result2 != value2 {
		t.Errorf("MustReturn2 failed, expected: (%v, %v), got: (%v, %v)", value1, value2, result1, result2)
	}

	// Test case 2: With error
	value3 := 0
	value4 := ""
	err2 := errors.New("some error")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustReturn2 did not panic")
		}
	}()
	MustReturn2(value3, value4, err2)
}
