package gokit

import (
	"context"
	"errors"
	"time"
)

// Retry 重试执行，最多不超过指定次数，每次间隔固定时长
func Retry(ctx context.Context, count int, wait time.Duration, fn func() error) error {
	return retry(ctx, count, wait, false, fn)
}

// BackoffRetry 重试执行，最多不超过指定次数，每次间隔时长翻倍
func BackoffRetry(ctx context.Context, count int, wait time.Duration, fn func() error) error {
	return retry(ctx, count, wait, true, fn)
}

func retry(ctx context.Context, count int, wait time.Duration, backoff bool, fn func() error) (err error) {
	if count <= 0 {
		return errors.New("invalid retry count")
	}

	tryExecute := func(wait time.Duration) (stop bool, err error) {
		select {
		case <-ctx.Done():
			return true, ctx.Err()
		default:
			if wait == 0 {
				return false, fn()
			}

			select {
			case <-ctx.Done():
				return true, ctx.Err()
			case <-time.After(wait):
				return false, fn()
			}
		}
	}

	var stop bool
	if stop, err = tryExecute(0); err == nil || stop {
		return
	}

	for i, c := 0, count-1; i < c; i++ {
		if stop, err = tryExecute(wait); err == nil || stop {
			return
		}

		if backoff {
			wait = wait * 2
		}
	}

	return
}
