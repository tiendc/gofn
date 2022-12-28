package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Filter(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, Filter[int](nil, func(v int) bool { return true }))
	assert.Equal(t, []float32{}, Filter([]float32{}, func(v float32) bool { return true }))

	assert.Equal(t, []int{-4, 0}, Filter([]int{2, -4, 6, 0}, func(v int) bool { return v <= 0 }))
	assert.Equal(t, []float32{2, 6, 0.0001}, Filter([]float32{2, -4, 6, 0.0001}, func(v float32) bool { return v > 0 }))
	assert.Equal(t, []string{"one", "two"}, Filter([]string{"one", "", "two"}, func(v string) bool { return len(v) > 0 }))
}

func Test_FilterPtr(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterPtr[int](nil, func(v *int) bool { return true }))
	assert.Equal(t, []float32{}, FilterPtr([]float32{}, func(v *float32) bool { return true }))

	assert.Equal(t, []int{-4, 0}, FilterPtr([]int{2, -4, 6, 0}, func(v *int) bool { return *v <= 0 }))
	assert.Equal(t, []float32{2, 6, 0.0001}, FilterPtr([]float32{2, -4, 6, 0.0001}, func(v *float32) bool { return *v > 0 }))
	assert.Equal(t, []string{"one", "two"}, FilterPtr([]string{"one", "", "two"}, func(v *string) bool { return len(*v) > 0 }))
}

func Test_FilterLT(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterLT[int](nil, 0))
	assert.Equal(t, []int{}, FilterLT([]int{}, 0))

	assert.Equal(t, []int{-10, -3}, FilterLT([]int{0, -10, 1, -3, 0, 7}, 0))
	assert.Equal(t, []string{"Ab"}, FilterLT([]string{"Ab", "cd", "a"}, "a"))
}

func Test_FilterLTE(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterLTE[int](nil, 0))
	assert.Equal(t, []int{}, FilterLTE([]int{}, 0))

	assert.Equal(t, []int{0, -10, -3, 0}, FilterLTE([]int{0, -10, 1, -3, 0, 7}, 0))
	assert.Equal(t, []string{"Ab", "a"}, FilterLTE([]string{"Ab", "cd", "a"}, "a"))
}

func Test_FilterGT(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterGT[int](nil, 0))
	assert.Equal(t, []int{}, FilterGT([]int{}, 0))

	assert.Equal(t, []int{1, 7}, FilterGT([]int{0, -10, 1, -3, 0, 7}, 0))
	assert.Equal(t, []string{"cd"}, FilterGT([]string{"Ab", "cd", "a"}, "a"))
}

func Test_FilterGTE(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterGTE[int](nil, 0))
	assert.Equal(t, []int{}, FilterGTE([]int{}, 0))

	assert.Equal(t, []int{0, 1, 0, 7}, FilterGTE([]int{0, -10, 1, -3, 0, 7}, 0))
	assert.Equal(t, []string{"cd", "a"}, FilterGTE([]string{"Ab", "cd", "a"}, "a"))
}

func Test_FilterNE(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterNE[int](nil, 0))
	assert.Equal(t, []int{}, FilterNE([]int{}, 0))

	assert.Equal(t, []int{-10, 1, -3, 7}, FilterNE([]int{0, -10, 1, -3, 0, 7}, 0))
	assert.Equal(t, []string{"Ab", "cd"}, FilterNE([]string{"Ab", "cd", "a"}, "a"))
}

func Test_FilterIN(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterIN[int](nil, 0))
	assert.Equal(t, []int{}, FilterIN([]int{}, 0))

	assert.Equal(t, []int{1, 7}, FilterIN([]int{0, -10, 1, -3, 0, 7}, 1, 3, 7))
	assert.Equal(t, []string{"a"}, FilterIN([]string{"Ab", "cd", "a"}, "a", "b"))
}

func Test_FilterNIN(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterNIN[int](nil, 0))
	assert.Equal(t, []int{}, FilterNIN([]int{}, 0))

	assert.Equal(t, []int{-10, 1, -3}, FilterNIN([]int{0, -10, 1, -3, 0, 7}, 3, 7, 0))
	assert.Equal(t, []string{"Ab"}, FilterNIN([]string{"Ab", "cd", "a"}, "a", "b", "cd"))
}
