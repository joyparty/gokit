package gokit

import (
	"errors"
	"testing"
)

func TestMustReturn(t *testing.T) {
	// No Error
	foo := func() (int, error) {
		return 42, nil
	}

	result1, _ := foo()
	actual1 := MustReturn(foo())
	if result1 != actual1 {
		t.Errorf("expected %d, got %d", result1, actual1)
	}

	// return error
	fooError := func() (int, error) {
		return 0, errors.New("some error")
	}
	defer func() {
		if v := recover(); v == nil {
			t.Error("MustReturn did not panic")
		}
	}()
	_ = MustReturn(fooError())
}

func TestMustReturn2(t *testing.T) {
	// No Error
	foo := func() (int, string, error) {
		return 42, "hello", nil
	}

	result1, result2, _ := foo()
	actual1, actual2 := MustReturn2(foo())
	if result1 != actual1 || result2 != actual2 {
		t.Errorf("expected %d, %s, got %d, %s", result1, result2, actual1, actual2)
	}

	// return error
	fooError := func() (int, string, error) {
		return 0, "", errors.New("some error")
	}
	defer func() {
		if v := recover(); v == nil {
			t.Error("MustReturn2 did not panic")
		}
	}()
	_, _ = MustReturn2(fooError())
}
