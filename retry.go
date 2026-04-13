package gofn

import (
	"math"
	"math/rand"
	"time"
)

type execRetryKind int8

const (
	execRetryFixedDelay execRetryKind = iota
	execRetryIncrementalDelay
	execRetryExponentialBackoff
)

type ExecRetryConfig struct {
	kind             execRetryKind
	delay            time.Duration
	maxDelay         time.Duration
	incremental      time.Duration
	expBackoffJitter time.Duration
}

func (cfg *ExecRetryConfig) nextDelay(retry int) time.Duration {
	if cfg.kind == execRetryFixedDelay {
		return cfg.delay
	}

	if cfg.kind == execRetryIncrementalDelay {
		delay := cfg.delay
		if retry > 0 {
			delay = cfg.delay + time.Duration(retry)*cfg.incremental
		}
		if cfg.maxDelay > 0 && delay > cfg.maxDelay {
			return cfg.maxDelay
		}
		return delay
	}

	// Expo backoff
	jitter := time.Duration(0)
	if cfg.expBackoffJitter > 0 {
		jitter = time.Duration(rand.Int63n(int64(cfg.expBackoffJitter))) //nolint:gosec
	}

	exp := 1.0
	if retry > 0 {
		exp = math.Pow(2, float64(retry)) //nolint:mnd
	}
	delay := time.Duration(exp*float64(cfg.delay)) + jitter
	if cfg.maxDelay > 0 && delay > cfg.maxDelay {
		return cfg.maxDelay
	}
	return delay
}

type ExecRetryOption func(*ExecRetryConfig)

func ExecRetryDelayMax(maxDelay time.Duration) ExecRetryOption {
	return func(config *ExecRetryConfig) {
		config.maxDelay = maxDelay
	}
}

func ExecRetryDelayIncr(incremental time.Duration) ExecRetryOption {
	return func(config *ExecRetryConfig) {
		config.kind = execRetryIncrementalDelay
		config.incremental = incremental
	}
}

func ExecRetryDelayExpoBackoff(jitter time.Duration) ExecRetryOption {
	return func(config *ExecRetryConfig) {
		config.kind = execRetryExponentialBackoff
		config.expBackoffJitter = jitter
	}
}

func ExecRetry(
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
		time.Sleep(nextDelay)
		retry++
		nextDelay = cfg.nextDelay(retry)
	}
}

func ExecRetry2[T any](
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
		time.Sleep(nextDelay)
		retry++
		nextDelay = cfg.nextDelay(retry)
	}
}

func ExecRetry3[T1, T2 any](
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
		time.Sleep(nextDelay)
		retry++
		nextDelay = cfg.nextDelay(retry)
	}
}

func ExecRetry4[T1, T2, T3 any](
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
		time.Sleep(nextDelay)
		retry++
		nextDelay = cfg.nextDelay(retry)
	}
}

func ExecRetry5[T1, T2, T3, T4 any](
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
		time.Sleep(nextDelay)
		retry++
		nextDelay = cfg.nextDelay(retry)
	}
}
