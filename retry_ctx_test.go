package gofn

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecRetryCtx(t *testing.T) {
	t.Run("Context Canceled Immediately", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		count := 0
		err := ExecRetryCtx(ctx, func() error {
			count++
			return errors.New("fail")
		}, 3, time.Second, ExecRetryDelayMax(time.Millisecond))
		assert.Error(t, err)
		assert.Equal(t, context.Canceled, err)
		assert.Equal(t, 1, count)
	})

	t.Run("Context Canceled Midway", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
		defer cancel()

		count := 0
		err := ExecRetryCtx(ctx, func() error {
			count++
			return errors.New("fail")
		}, 3, 30*time.Millisecond)
		assert.Error(t, err)
		assert.Equal(t, context.DeadlineExceeded, err)
		assert.Equal(t, 1, count)
	})

	t.Run("Success Before Cancellation", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()

		count := 0
		err := ExecRetryCtx(ctx, func() error {
			count++
			if count < 2 {
				return errors.New("fail")
			}
			return nil
		}, 3, time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 2, count)
	})

	t.Run("Zero maxRetries means no retry", func(t *testing.T) {
		ctx := context.Background()
		count := 0
		err := ExecRetryCtx(ctx, func() error {
			count++
			return errors.New("fail")
		}, 0, time.Millisecond)
		assert.Error(t, err)
		assert.Equal(t, 1, count)
	})
}

func TestExecRetryCtxVariants(t *testing.T) {
	ctx := context.Background()

	t.Run("ExecRetryCtx2", func(t *testing.T) {
		v, err := ExecRetryCtx2(ctx, func() (int, error) {
			return 1, nil
		}, 1, 0, ExecRetryDelayMax(time.Millisecond))
		assert.NoError(t, err)
		assert.Equal(t, 1, v)

		v, err = ExecRetryCtx2(ctx, func() (int, error) {
			return 1, errors.New("fail")
		}, 0, 0, ExecRetryDelayMax(time.Millisecond))
		assert.Error(t, err)
		assert.Equal(t, 1, v)

		cancelCtx, cancel := context.WithCancel(context.Background())
		cancel()
		v, err = ExecRetryCtx2(cancelCtx, func() (int, error) {
			return 2, errors.New("fail")
		}, 1, time.Second)
		assert.Equal(t, context.Canceled, err)
		assert.Equal(t, 2, v)

		count2 := 0
		v, err = ExecRetryCtx2(ctx, func() (int, error) {
			count2++
			if count2 < 2 {
				return 2, errors.New("fail")
			}
			return 2, nil
		}, 2, time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 2, v)
	})

	t.Run("ExecRetryCtx3", func(t *testing.T) {
		v1, v2, err := ExecRetryCtx3(ctx, func() (int, string, error) {
			return 1, "a", nil
		}, 1, 0, ExecRetryDelayMax(time.Millisecond))
		assert.NoError(t, err)
		assert.Equal(t, 1, v1)
		assert.Equal(t, "a", v2)

		v1, v2, err = ExecRetryCtx3(ctx, func() (int, string, error) {
			return 1, "a", errors.New("fail")
		}, 0, 0, ExecRetryDelayMax(time.Millisecond))
		assert.Error(t, err)
		assert.Equal(t, 1, v1)
		assert.Equal(t, "a", v2)

		cancelCtx, cancel := context.WithCancel(context.Background())
		cancel()
		v1, v2, err = ExecRetryCtx3(cancelCtx, func() (int, string, error) {
			return 2, "b", errors.New("fail")
		}, 1, time.Second)
		assert.Equal(t, context.Canceled, err)
		assert.Equal(t, 2, v1)
		assert.Equal(t, "b", v2)

		count3 := 0
		v1, v2, err = ExecRetryCtx3(ctx, func() (int, string, error) {
			count3++
			if count3 < 2 {
				return 2, "b", errors.New("fail")
			}
			return 2, "b", nil
		}, 2, time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 2, v1)
		assert.Equal(t, "b", v2)
	})

	t.Run("ExecRetryCtx4", func(t *testing.T) {
		v1, v2, v3, err := ExecRetryCtx4(ctx, func() (int, string, bool, error) {
			return 1, "a", true, nil
		}, 1, 0, ExecRetryDelayMax(time.Millisecond))
		assert.NoError(t, err)
		assert.Equal(t, 1, v1)
		assert.Equal(t, "a", v2)
		assert.Equal(t, true, v3)

		v1, v2, v3, err = ExecRetryCtx4(ctx, func() (int, string, bool, error) {
			return 1, "a", true, errors.New("fail")
		}, 0, 0, ExecRetryDelayMax(time.Millisecond))
		assert.Error(t, err)
		assert.Equal(t, 1, v1)
		assert.Equal(t, "a", v2)
		assert.Equal(t, true, v3)

		cancelCtx, cancel := context.WithCancel(context.Background())
		cancel()
		v1, v2, v3, err = ExecRetryCtx4(cancelCtx, func() (int, string, bool, error) {
			return 2, "b", false, errors.New("fail")
		}, 1, time.Second)
		assert.Equal(t, context.Canceled, err)
		assert.Equal(t, 2, v1)
		assert.Equal(t, "b", v2)
		assert.Equal(t, false, v3)

		count4 := 0
		v1, v2, v3, err = ExecRetryCtx4(ctx, func() (int, string, bool, error) {
			count4++
			if count4 < 2 {
				return 2, "b", false, errors.New("fail")
			}
			return 2, "b", false, nil
		}, 2, time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 2, v1)
		assert.Equal(t, "b", v2)
		assert.Equal(t, false, v3)
	})

	t.Run("ExecRetryCtx5", func(t *testing.T) {
		v1, v2, v3, v4, err := ExecRetryCtx5(ctx, func() (int, string, bool, float32, error) {
			return 1, "a", true, 2.5, nil
		}, 1, 0, ExecRetryDelayMax(time.Millisecond))
		assert.NoError(t, err)
		assert.Equal(t, 1, v1)
		assert.Equal(t, "a", v2)
		assert.Equal(t, true, v3)
		assert.Equal(t, float32(2.5), v4)

		v1, v2, v3, v4, err = ExecRetryCtx5(ctx, func() (int, string, bool, float32, error) {
			return 1, "a", true, 2.5, errors.New("fail")
		}, 0, 0, ExecRetryDelayMax(time.Millisecond))
		assert.Error(t, err)
		assert.Equal(t, 1, v1)
		assert.Equal(t, "a", v2)
		assert.Equal(t, true, v3)
		assert.Equal(t, float32(2.5), v4)

		cancelCtx, cancel := context.WithCancel(context.Background())
		cancel()
		v1, v2, v3, v4, err = ExecRetryCtx5(cancelCtx, func() (int, string, bool, float32, error) {
			return 2, "b", false, 3.5, errors.New("fail")
		}, 1, time.Second)
		assert.Equal(t, context.Canceled, err)
		assert.Equal(t, 2, v1)
		assert.Equal(t, "b", v2)
		assert.Equal(t, false, v3)
		assert.Equal(t, float32(3.5), v4)

		count5 := 0
		v1, v2, v3, v4, err = ExecRetryCtx5(ctx, func() (int, string, bool, float32, error) {
			count5++
			if count5 < 2 {
				return 2, "b", false, 3.5, errors.New("fail")
			}
			return 2, "b", false, 3.5, nil
		}, 2, time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 2, v1)
		assert.Equal(t, "b", v2)
		assert.Equal(t, false, v3)
		assert.Equal(t, float32(3.5), v4)
	})
}
