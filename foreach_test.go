package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ForEach(t *testing.T) {
	copy1 := []int{}
	ForEach[int](nil, func(i int, t int) { copy1 = append(copy1, t) })
	assert.Equal(t, []int{}, copy1)

	copy2 := []int64{}
	ForEach([]int64{1, 2, 3}, func(i int, t int64) { copy2 = append(copy2, t) })
	assert.Equal(t, []int64{1, 2, 3}, copy2)
}

func Test_ForEachPtr(t *testing.T) {
	copy1 := []int{}
	ForEachPtr[int](nil, func(i int, t *int) { copy1 = append(copy1, *t) })
	assert.Equal(t, []int{}, copy1)

	copy2 := []int64{}
	ForEachPtr([]int64{1, 2, 3}, func(i int, t *int64) { copy2 = append(copy2, *t) })
	assert.Equal(t, []int64{1, 2, 3}, copy2)
}

func Test_ForEachReverse(t *testing.T) {
	copy1 := []int{}
	ForEachReverse[int](nil, func(i int, t int) { copy1 = append(copy1, t) })
	assert.Equal(t, []int{}, copy1)

	copy2 := []int64{}
	ForEachReverse([]int64{1, 2, 3}, func(i int, t int64) { copy2 = append(copy2, t) })
	assert.Equal(t, []int64{3, 2, 1}, copy2)
}

func Test_ForEachReversePtr(t *testing.T) {
	copy1 := []int{}
	ForEachReversePtr[int](nil, func(i int, t *int) { copy1 = append(copy1, *t) })
	assert.Equal(t, []int{}, copy1)

	copy2 := []int64{}
	ForEachReversePtr([]int64{1, 2, 3}, func(i int, t *int64) { copy2 = append(copy2, *t) })
	assert.Equal(t, []int64{3, 2, 1}, copy2)
}
