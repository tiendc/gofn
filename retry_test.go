package gofn

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecRetryConfig_nextDelay(t *testing.T) {
	cfg1 := &ExecRetryConfig{
		kind:  execRetryFixedDelay,
		delay: 10 * time.Millisecond,
	}
	assert.Equal(t, 10*time.Millisecond, cfg1.nextDelay(0))
	assert.Equal(t, 10*time.Millisecond, cfg1.nextDelay(5))

	cfg2 := &ExecRetryConfig{
		kind:        execRetryIncrementalDelay,
		delay:       10 * time.Millisecond,
		incremental: 2 * time.Millisecond,
		maxDelay:    13 * time.Millisecond,
	}
	assert.Equal(t, 10*time.Millisecond, cfg2.nextDelay(0))
	assert.Equal(t, 12*time.Millisecond, cfg2.nextDelay(1))
	assert.Equal(t, 13*time.Millisecond, cfg2.nextDelay(2))
	assert.Equal(t, 13*time.Millisecond, cfg2.nextDelay(3))
	assert.Equal(t, 13*time.Millisecond, cfg2.nextDelay(4))

	cfg3 := &ExecRetryConfig{
		kind:             execRetryExponentialBackoff,
		delay:            5 * time.Millisecond,
		maxDelay:         20 * time.Millisecond,
		expBackoffJitter: 0,
	}
	assert.Equal(t, 5*time.Millisecond, cfg3.nextDelay(0))
	assert.Equal(t, 10*time.Millisecond, cfg3.nextDelay(1))
	assert.Equal(t, 20*time.Millisecond, cfg3.nextDelay(2)) // 2^2 * 5 = 20
	assert.Equal(t, 20*time.Millisecond, cfg3.nextDelay(3)) // 2^3 * 5 = 40 -> maxDelay

	cfg4 := &ExecRetryConfig{
		kind:             execRetryExponentialBackoff,
		delay:            5 * time.Millisecond,
		maxDelay:         20 * time.Millisecond,
		expBackoffJitter: 2 * time.Millisecond,
	}
	delayWithJitter := cfg4.nextDelay(0)
	assert.True(t, delayWithJitter >= 5*time.Millisecond && delayWithJitter < 7*time.Millisecond)
}

func TestExecRetry(t *testing.T) {
	t.Run("Immediate Success", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			return nil
		}, 3, time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 1, count)
	})

	t.Run("Success After Retries", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			if count < 3 {
				return errors.New("fail")
			}
			return nil
		}, 3, time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 3, count)
	})

	t.Run("Exhausted maxRetries", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			return errors.New("fail")
		}, 3, time.Millisecond)
		assert.Error(t, err)
		assert.Equal(t, "fail", err.Error())
		assert.Equal(t, 4, count) // initial call + 3 retries
	})

	t.Run("Exhausted infinite maxRetries", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			if count == 5 {
				return nil
			}
			return errors.New("fail")
		}, -1, time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 5, count)
	})

	t.Run("Zero maxRetries means no retry", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			return errors.New("fail")
		}, 0, time.Millisecond)
		assert.Error(t, err)
		assert.Equal(t, 1, count)
	})

	t.Run("Options - Delay Max", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			return errors.New("fail")
		}, 3, 5*time.Millisecond, ExecRetryDelayMax(10*time.Millisecond), ExecRetryDelayIncr(20*time.Millisecond))
		assert.Error(t, err)
		assert.Equal(t, 4, count)
	})

	t.Run("Options - Expo Backoff", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			return errors.New("fail")
		}, 2, 5*time.Millisecond, ExecRetryDelayExpoBackoff(time.Millisecond))
		assert.Error(t, err)
		assert.Equal(t, 3, count)
	})
}

func TestExecRetryVariants(t *testing.T) {
	v1, err := ExecRetry2(func() (int, error) {
		return 1, nil
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.NoError(t, err)
	assert.Equal(t, 1, v1)

	v1_3, v2_3, err := ExecRetry3(func() (int, string, error) {
		return 1, "a", nil
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.NoError(t, err)
	assert.Equal(t, 1, v1_3)
	assert.Equal(t, "a", v2_3)

	v1_4, v2_4, v3_4, err := ExecRetry4(func() (int, string, bool, error) {
		return 1, "a", true, nil
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.NoError(t, err)
	assert.Equal(t, 1, v1_4)
	assert.Equal(t, "a", v2_4)
	assert.Equal(t, true, v3_4)

	v1_5, v2_5, v3_5, v4_5, err := ExecRetry5(func() (int, string, bool, float32, error) {
		return 1, "a", true, 2.5, nil
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.NoError(t, err)
	assert.Equal(t, 1, v1_5)
	assert.Equal(t, "a", v2_5)
	assert.Equal(t, true, v3_5)
	assert.Equal(t, float32(2.5), v4_5)

	// Exhausted variants
	v1, err = ExecRetry2(func() (int, error) {
		return 1, errors.New("fail")
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.Error(t, err)
	assert.Equal(t, 1, v1)

	v1_3, v2_3, err = ExecRetry3(func() (int, string, error) {
		return 1, "a", errors.New("fail")
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.Error(t, err)
	assert.Equal(t, 1, v1_3)
	assert.Equal(t, "a", v2_3)

	v1_4, v2_4, v3_4, err = ExecRetry4(func() (int, string, bool, error) {
		return 1, "a", true, errors.New("fail")
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.Error(t, err)
	assert.Equal(t, 1, v1_4)
	assert.Equal(t, "a", v2_4)
	assert.Equal(t, true, v3_4)

	v1_5, v2_5, v3_5, v4_5, err = ExecRetry5(func() (int, string, bool, float32, error) {
		return 1, "a", true, 2.5, errors.New("fail")
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.Error(t, err)
	assert.Equal(t, 1, v1_5)
	assert.Equal(t, "a", v2_5)
	assert.Equal(t, true, v3_5)
	assert.Equal(t, float32(2.5), v4_5)
}
