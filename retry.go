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

	tryExecute := func(wait time.Duration) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if wait == 0 {
				return fn()
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(wait):
				return fn()
			}
		}
	}

	if err = tryExecute(0); err == nil {
		return
	}

	for i, c := 0, count-1; i < c; i++ {
		if err = tryExecute(wait); err == nil {
			return
		}

		if backoff {
			wait = wait * 2
		}
	}
	return
}
