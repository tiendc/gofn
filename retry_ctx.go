package gofn

import (
	"context"
	"time"
)

func ExecRetryCtx(
	ctx context.Context,
	fn func() error,
	maxRetries int,
	delay time.Duration,
	options ...ExecRetryOption,
) error {
	cfg := &ExecRetryConfig{
		kind:  execRetryFixedDelay,
		delay: delay,
	}
	for _, option := range options {
		option(cfg)
	}

	retry := 0
	nextDelay := cfg.delay
	for {
		err := fn()
		if err == nil {
			return nil
		}
		if maxRetries >= 0 && retry >= maxRetries {
			return err
		}
		if cfg.shouldRetry != nil && !cfg.shouldRetry(err) {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err() // nolint: wrapcheck
		case <-time.After(nextDelay):
		}
		retry++
		nextDelay = cfg.nextDelay(retry)
	}
}

func ExecRetryCtx2[T any](
	ctx context.Context,
	fn func() (T, error),
	maxRetries int,
	delay time.Duration,
	options ...ExecRetryOption,
) (T, error) {
	cfg := &ExecRetryConfig{
		kind:  execRetryFixedDelay,
		delay: delay,
	}
	for _, option := range options {
		option(cfg)
	}

	retry := 0
	nextDelay := cfg.delay
	for {
		v, err := fn()
		if err == nil {
			return v, nil
		}
		if maxRetries >= 0 && retry >= maxRetries {
			return v, err
		}
		if cfg.shouldRetry != nil && !cfg.shouldRetry(err) {
			return v, err
		}
		select {
		case <-ctx.Done():
			return v, ctx.Err() // nolint: wrapcheck
		case <-time.After(nextDelay):
		}
		retry++
		nextDelay = cfg.nextDelay(retry)
	}
}

func ExecRetryCtx3[T1, T2 any](
	ctx context.Context,
	fn func() (T1, T2, error),
	maxRetries int,
	delay time.Duration,
	options ...ExecRetryOption,
) (T1, T2, error) {
	cfg := &ExecRetryConfig{
		kind:  execRetryFixedDelay,
		delay: delay,
	}
	for _, option := range options {
		option(cfg)
	}

	retry := 0
	nextDelay := cfg.delay
	for {
		v1, v2, err := fn()
		if err == nil {
			return v1, v2, nil
		}
		if maxRetries >= 0 && retry >= maxRetries {
			return v1, v2, err
		}
		if cfg.shouldRetry != nil && !cfg.shouldRetry(err) {
			return v1, v2, err
		}
		select {
		case <-ctx.Done():
			return v1, v2, ctx.Err() // nolint: wrapcheck
		case <-time.After(nextDelay):
		}
		retry++
		nextDelay = cfg.nextDelay(retry)
	}
}

func ExecRetryCtx4[T1, T2, T3 any](
	ctx context.Context,
	fn func() (T1, T2, T3, error),
	maxRetries int,
	delay time.Duration,
	options ...ExecRetryOption,
) (T1, T2, T3, error) {
	cfg := &ExecRetryConfig{
		kind:  execRetryFixedDelay,
		delay: delay,
	}
	for _, option := range options {
		option(cfg)
	}

	retry := 0
	nextDelay := cfg.delay
	for {
		v1, v2, v3, err := fn()
		if err == nil {
			return v1, v2, v3, nil
		}
		if maxRetries >= 0 && retry >= maxRetries {
			return v1, v2, v3, err
		}
		if cfg.shouldRetry != nil && !cfg.shouldRetry(err) {
			return v1, v2, v3, err
		}
		select {
		case <-ctx.Done():
			return v1, v2, v3, ctx.Err() // nolint: wrapcheck
		case <-time.After(nextDelay):
		}
		retry++
		nextDelay = cfg.nextDelay(retry)
	}
}

func ExecRetryCtx5[T1, T2, T3, T4 any](
	ctx context.Context,
	fn func() (T1, T2, T3, T4, error),
	maxRetries int,
	delay time.Duration,
	options ...ExecRetryOption,
) (T1, T2, T3, T4, error) {
	cfg := &ExecRetryConfig{
		kind:  execRetryFixedDelay,
		delay: delay,
	}
	for _, option := range options {
		option(cfg)
	}

	retry := 0
	nextDelay := cfg.delay
	for {
		v1, v2, v3, v4, err := fn()
		if err == nil {
			return v1, v2, v3, v4, nil
		}
		if maxRetries >= 0 && retry >= maxRetries {
			return v1, v2, v3, v4, err
		}
		if cfg.shouldRetry != nil && !cfg.shouldRetry(err) {
			return v1, v2, v3, v4, err
		}
		select {
		case <-ctx.Done():
			return v1, v2, v3, v4, ctx.Err() // nolint: wrapcheck
		case <-time.After(nextDelay):
		}
		retry++
		nextDelay = cfg.nextDelay(retry)
	}
}
