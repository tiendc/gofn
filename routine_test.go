package gofn

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// nolint
func Test_ExecTasks(t *testing.T) {
	type ctxData struct {
		mu     sync.Mutex
		result []int
	}

	errTest := errors.New("test error")

	task1 := func(ctx context.Context) error {
		data := ctx.Value("data").(*ctxData)
		for i := 0; i < 10; i++ {
			if err := ctx.Err(); err != nil {
				return err
			}
			data.mu.Lock()
			data.result = append(data.result, i)
			data.mu.Unlock()
			time.Sleep(time.Duration(20+rand.Intn(100)) * time.Millisecond)
		}
		return nil
	}
	task2 := func(ctx context.Context) error {
		data := ctx.Value("data").(*ctxData)
		for i := 10; i < 20; i++ {
			if err := ctx.Err(); err != nil {
				return err
			}
			data.mu.Lock()
			data.result = append(data.result, i)
			data.mu.Unlock()
			time.Sleep(time.Duration(20+rand.Intn(100)) * time.Millisecond)
		}
		return nil
	}
	task3 := func(ctx context.Context) error {
		data := ctx.Value("data").(*ctxData)
		for i := 20; i < 30; i++ {
			if err := ctx.Err(); err != nil {
				return err
			}
			if i == 25 {
				return errTest
			}
			data.mu.Lock()
			data.result = append(data.result, i)
			data.mu.Unlock()
			time.Sleep(time.Duration(20+rand.Intn(100)) * time.Millisecond)
		}
		return nil
	}
	task4 := func(ctx context.Context) error {
		data := ctx.Value("data").(*ctxData)
		for i := 30; i < 40; i++ {
			if err := ctx.Err(); err != nil {
				return err
			}
			if i == 35 {
				return errTest
			}
			data.mu.Lock()
			data.result = append(data.result, i)
			data.mu.Unlock()
			time.Sleep(time.Duration(20+rand.Intn(100)) * time.Millisecond)
		}
		return nil
	}
	task5 := func(ctx context.Context) error {
		data := ctx.Value("data").(*ctxData)
		for i := 30; i < 40; i++ {
			if err := ctx.Err(); err != nil {
				return err
			}
			if i == 35 {
				panic(errTest)
			}
			data.mu.Lock()
			data.result = append(data.result, i)
			data.mu.Unlock()
			time.Sleep(time.Duration(20+rand.Intn(100)) * time.Millisecond)
		}
		return nil
	}

	checkRes := func(res []int, start, end int) bool {
		values := []int{}
		for i := start; i < end; i++ {
			values = append(values, i)
		}
		return ContainAll(res, values...)
	}

	t.Run("no tasks passed", func(t *testing.T) {
		ctx := context.Background()
		errMap := ExecTasksEx(ctx, 0, false)
		assert.Equal(t, 0, len(errMap))
	})

	t.Run("no pool size, no stop on error", func(t *testing.T) {
		data := &ctxData{}
		ctx := context.WithValue(context.Background(), "data", data)

		errMap := ExecTasksEx(ctx, 0, false, task1, task2, task3, task4)
		assert.Equal(t, 2, len(errMap))

		result := data.result
		assert.Equal(t, 30, len(result))
		assert.True(t, checkRes(result, 0, 10) && checkRes(result, 10, 20) &&
			checkRes(result, 20, 25) && checkRes(result, 30, 35))
	})

	t.Run("big pool size, no stop on error", func(t *testing.T) {
		data := &ctxData{}
		ctx := context.WithValue(context.Background(), "data", data)

		errMap := ExecTasksEx(ctx, 10, false, task1, task2, task3, task4)
		assert.Equal(t, 2, len(errMap))

		result := data.result
		assert.Equal(t, 30, len(result))
		assert.True(t, checkRes(result, 0, 10) && checkRes(result, 10, 20) &&
			checkRes(result, 20, 25) && checkRes(result, 30, 35))
	})

	t.Run("no pool size, stop on error", func(t *testing.T) {
		data := &ctxData{}
		ctx := context.WithValue(context.Background(), "data", data)

		// NOTE: call ExecTasks() as ExecTasks() default to stopOnError is true
		err := ExecTasks(ctx, 0, task1, task2, task3, task4)
		assert.NotNil(t, err)

		data.mu.Lock()
		result := data.result
		assert.True(t, 30 > len(result))
		data.mu.Unlock()
	})

	t.Run("no pool size, stop on error. but no error occurred", func(t *testing.T) {
		data := &ctxData{}
		ctx := context.WithValue(context.Background(), "data", data)

		// NOTE: call ExecTasks() as ExecTasks() default to stopOnError is true
		err := ExecTasks(ctx, 0, task1, task2)
		assert.Nil(t, err)

		result := data.result
		assert.Equal(t, 20, len(result))
		assert.True(t, checkRes(result, 0, 20))
	})

	t.Run("pool size = 1, no stop on error", func(t *testing.T) {
		data := &ctxData{}
		ctx := context.WithValue(context.Background(), "data", data)

		errMap := ExecTasksEx(ctx, 1, false, task1, task2, task3, task4)
		assert.True(t, len(errMap) == 2)

		result := data.result
		// As pool size = 1, result from the tasks are in order of the execution
		assert.True(t, checkRes(result[:10], 0, 10) && checkRes(result[10:20], 10, 20) &&
			checkRes(result[20:25], 20, 25) && checkRes(result[25:], 30, 35))
	})

	t.Run("context timed out", func(t *testing.T) {
		data := &ctxData{}
		ctx := context.WithValue(context.Background(), "data", data)
		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
		defer cancel()

		errMap := ExecTasksEx(ctx, 0, false, task1, task2, task3, task4)
		assert.Equal(t, 4, len(errMap))

		result := data.result
		assert.True(t, 30 > len(result) && len(result) > 0)
	})

	t.Run("context canceled too early", func(t *testing.T) {
		data := &ctxData{}
		ctx := context.WithValue(context.Background(), "data", data)
		ctx, cancel := context.WithCancel(ctx)
		cancel()

		errMap := ExecTasksEx(ctx, 0, false, task1, task2, task3, task4)
		assert.Equal(t, 4, len(errMap))

		result := data.result
		assert.Equal(t, 0, len(result))
	})

	t.Run("panic occurred in the task function", func(t *testing.T) {
		data := &ctxData{}
		ctx := context.WithValue(context.Background(), "data", data)

		errMap := ExecTasksEx(ctx, 0, false, task1, task2, task3, task4, task5)
		assert.Equal(t, 3, len(errMap))

		result := data.result
		assert.True(t, len(result) > 0)
	})
}
