package gokit

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	// Test case 1: Retry succeeds on first attempt
	count := 3
	wait := 100 * time.Millisecond
	err := Retry(context.Background(), count, wait, func() error {
		return nil
	})
	if err != nil {
		t.Errorf("Retry failed unexpectedly, error: %v", err)
	}

	// Test case 2: Retry fails after specified attempts
	count = 3
	wait = 100 * time.Millisecond
	expectedErr := errors.New("retry failed")
	err = Retry(context.Background(), count, wait, func() error {
		return expectedErr
	})
	if err != expectedErr {
		t.Errorf("Retry did not return expected error, expected: %v, got: %v", expectedErr, err)
	}

	// Test case 3: Retry succeeds after multiple attempts
	count = 3
	wait = 100 * time.Millisecond
	attempts := 0
	err = Retry(context.Background(), count, wait, func() error {
		attempts++
		if attempts < count {
			return errors.New("retry failed")
		}
		return nil
	})
	if err != nil {
		t.Errorf("Retry failed unexpectedly, error: %v", err)
	}
}
