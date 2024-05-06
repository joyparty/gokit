package gokit

import (
	"context"
	"time"
)

// Retry 在首次执行失败之后，根据指定的次数重试执行，每次重试之间等待一段时间
func Retry(ctx context.Context, count int, wait time.Duration, fn func() error) error {
	return retry(ctx, count, wait, false, fn)
}

// BackoffRetry 在首次执行失败之后，根据指定的次数重试执行，每次重试之间等待一段时间，等待时间每次翻倍
func BackoffRetry(ctx context.Context, count int, wait time.Duration, fn func() error) error {
	return retry(ctx, count, wait, true, fn)
}

func retry(ctx context.Context, count int, wait time.Duration, backoff bool, fn func() error) (err error) {
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

	for i := 0; i < count; i++ {
		if err = tryExecute(wait); err == nil {
			return
		}

		if backoff {
			wait = wait * 2
		}
	}
	return
}
