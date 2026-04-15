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

func TestExecRetry_ShouldRetry(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	err3 := errors.New("error 3")

	t.Run("ExecRetryCheck - Retry Always", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			return err1
		}, 3, time.Nanosecond, ExecRetryCheck(func(err error) bool {
			return true
		}))
		assert.ErrorIs(t, err, err1)
		assert.Equal(t, 4, count)
	})

	t.Run("ExecRetryCheck - Retry Never", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			return err1
		}, 3, time.Nanosecond, ExecRetryCheck(func(err error) bool {
			return false
		}))
		assert.ErrorIs(t, err, err1)
		assert.Equal(t, 1, count)
	})

	t.Run("ExecRetryIfErrorIs - Matches", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			return err1
		}, 3, time.Nanosecond, ExecRetryIfErrorIs(err1))
		assert.ErrorIs(t, err, err1)
		assert.Equal(t, 4, count)
	})

	t.Run("ExecRetryIfErrorIs - Not Matches", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			return err2
		}, 3, time.Nanosecond, ExecRetryIfErrorIs(err1))
		assert.ErrorIs(t, err, err2)
		assert.Equal(t, 1, count)
	})

	t.Run("ExecRetryIfErrorIs - Multiple Errors", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			if count == 1 {
				return err1
			}
			return err2
		}, 3, time.Nanosecond, ExecRetryIfErrorIs(err1, err2))
		assert.ErrorIs(t, err, err2)
		assert.Equal(t, 4, count)
	})

	t.Run("ExecRetryIfErrorIsNot - Matches (stops)", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			return err1
		}, 3, time.Nanosecond, ExecRetryIfErrorIsNot(err1))
		assert.ErrorIs(t, err, err1)
		assert.Equal(t, 1, count)
	})

	t.Run("ExecRetryIfErrorIsNot - Not Matches (continues)", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			return err2
		}, 3, time.Nanosecond, ExecRetryIfErrorIsNot(err1))
		assert.ErrorIs(t, err, err2)
		assert.Equal(t, 4, count)
	})

	t.Run("ExecRetryIfErrorIsNot - Multiple Errors", func(t *testing.T) {
		count := 0
		err := ExecRetry(func() error {
			count++
			if count == 1 {
				return err1
			}
			return err3
		}, 3, time.Nanosecond, ExecRetryIfErrorIsNot(err2))
		assert.ErrorIs(t, err, err3)
		assert.Equal(t, 4, count)
	})
}

func TestExecRetry2(t *testing.T) {
	v1, err := ExecRetry2(func() (int, error) {
		return 1, nil
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.NoError(t, err)
	assert.Equal(t, 1, v1)

	v1, err = ExecRetry2(func() (int, error) {
		return 1, errors.New("fail")
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.Error(t, err)
	assert.Equal(t, 1, v1)

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	count := 0
	v, err := ExecRetry2(func() (int, error) {
		count++
		if count < 2 {
			return 1, err1
		}
		return 2, nil
	}, 3, time.Nanosecond, ExecRetryIfErrorIs(err1))
	assert.NoError(t, err)
	assert.Equal(t, 2, v)
	assert.Equal(t, 2, count)

	count = 0
	v, err = ExecRetry2(func() (int, error) {
		count++
		if count < 2 {
			return 1, err1
		}
		return 2, err2
	}, 3, time.Nanosecond, ExecRetryIfErrorIs(err1))
	assert.ErrorIs(t, err, err2)
	assert.Equal(t, 2, v)
	assert.Equal(t, 2, count)
}

func TestExecRetry3(t *testing.T) {
	v1, v2, err := ExecRetry3(func() (int, string, error) {
		return 1, "a", nil
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.NoError(t, err)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "a", v2)

	v1, v2, err = ExecRetry3(func() (int, string, error) {
		return 1, "a", errors.New("fail")
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.Error(t, err)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "a", v2)

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	count := 0
	v1, v2, err = ExecRetry3(func() (int, string, error) {
		count++
		if count < 2 {
			return 1, "a", err1
		}
		return 2, "b", nil
	}, 3, time.Nanosecond, ExecRetryIfErrorIs(err1))
	assert.NoError(t, err)
	assert.Equal(t, 2, v1)
	assert.Equal(t, "b", v2)
	assert.Equal(t, 2, count)

	count = 0
	v1, v2, err = ExecRetry3(func() (int, string, error) {
		count++
		if count < 2 {
			return 1, "a", err1
		}
		return 2, "b", err2
	}, 3, time.Nanosecond, ExecRetryIfErrorIs(err1))
	assert.ErrorIs(t, err, err2)
	assert.Equal(t, 2, v1)
	assert.Equal(t, "b", v2)
	assert.Equal(t, 2, count)
}

func TestExecRetry4(t *testing.T) {
	v1, v2, v3, err := ExecRetry4(func() (int, string, bool, error) {
		return 1, "a", true, nil
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.NoError(t, err)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "a", v2)
	assert.Equal(t, true, v3)

	v1, v2, v3, err = ExecRetry4(func() (int, string, bool, error) {
		return 1, "a", true, errors.New("fail")
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.Error(t, err)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "a", v2)
	assert.Equal(t, true, v3)

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	count := 0
	v1, v2, v3, err = ExecRetry4(func() (int, string, bool, error) {
		count++
		if count < 2 {
			return 1, "a", false, err1
		}
		return 2, "b", true, nil
	}, 3, time.Nanosecond, ExecRetryIfErrorIs(err1))
	assert.NoError(t, err)
	assert.Equal(t, 2, v1)
	assert.Equal(t, "b", v2)
	assert.Equal(t, true, v3)
	assert.Equal(t, 2, count)

	count = 0
	v1, v2, v3, err = ExecRetry4(func() (int, string, bool, error) {
		count++
		if count < 2 {
			return 1, "a", false, err1
		}
		return 2, "b", true, err2
	}, 3, time.Nanosecond, ExecRetryIfErrorIs(err1))
	assert.ErrorIs(t, err, err2)
	assert.Equal(t, 2, v1)
	assert.Equal(t, "b", v2)
	assert.Equal(t, true, v3)
	assert.Equal(t, 2, count)
}

func TestExecRetry5(t *testing.T) {
	v1, v2, v3, v4, err := ExecRetry5(func() (int, string, bool, float64, error) {
		return 1, "a", true, 2.5, nil
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.NoError(t, err)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "a", v2)
	assert.Equal(t, true, v3)
	assert.Equal(t, 2.5, v4)

	v1, v2, v3, v4, err = ExecRetry5(func() (int, string, bool, float64, error) {
		return 1, "a", true, 2.5, errors.New("fail")
	}, 1, 0, ExecRetryDelayMax(time.Millisecond))
	assert.Error(t, err)
	assert.Equal(t, 1, v1)
	assert.Equal(t, "a", v2)
	assert.Equal(t, true, v3)
	assert.Equal(t, 2.5, v4)

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	count := 0
	v1, v2, v3, v4, err = ExecRetry5(func() (int, string, bool, float64, error) {
		count++
		if count < 2 {
			return 1, "a", false, 1.2, err1
		}
		return 2, "b", true, 2.3, nil
	}, 3, time.Nanosecond, ExecRetryIfErrorIs(err1))
	assert.NoError(t, err)
	assert.Equal(t, 2, v1)
	assert.Equal(t, "b", v2)
	assert.Equal(t, true, v3)
	assert.Equal(t, 2.3, v4)
	assert.Equal(t, 2, count)

	count = 0
	v1, v2, v3, v4, err = ExecRetry5(func() (int, string, bool, float64, error) {
		count++
		if count < 2 {
			return 1, "a", false, 1.2, err1
		}
		return 2, "b", true, 2.3, err2
	}, 3, time.Nanosecond, ExecRetryIfErrorIs(err1))
	assert.ErrorIs(t, err, err2)
	assert.Equal(t, 2, v1)
	assert.Equal(t, "b", v2)
	assert.Equal(t, true, v3)
	assert.Equal(t, 2.3, v4)
	assert.Equal(t, 2, count)
}
