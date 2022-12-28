package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reduce(t *testing.T) {
	assert.Equal(t, 6, Reduce[int]([]int{1, 2, 3}, func(acc, v int) int { return acc + v }))
	assert.Equal(t, 8, Reduce[int]([]int{1, 2, 4}, func(acc, v int) int { return acc * v }))
	assert.Equal(t, 0, Reduce[int]([]int{1, 2, 0}, func(acc, v int) int { return acc * v }))
}

func Test_ReduceEx(t *testing.T) {
	assert.Equal(t, 7, ReduceEx[int]([]int{1, 2, 3}, func(acc, v, i int) int { return acc + v }, 1))
	assert.Equal(t, 8, ReduceEx[int]([]int{1, 2, 4}, func(acc, v, i int) int { return acc * v }, 1))
	assert.Equal(t, 0, ReduceEx[int]([]int{1, 2, 0}, func(acc, v, i int) int { return acc * v }, 1))
}
