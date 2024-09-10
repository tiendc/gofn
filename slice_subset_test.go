package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ContainSlice(t *testing.T) {
	assert.False(t, ContainSlice([]int{}, nil))
	assert.False(t, ContainSlice([]string{"one"}, []string{}))
	assert.False(t, ContainSlice([]string{"one", "two"}, []string{"Two"}))
	assert.False(t, ContainSlice([]int64{1, 2, 3}, []int64{1, 2, 3, 4}))
	assert.False(t, ContainSlice([]uint{0, 1, 2, 3, 4, 5}, []uint{3, 4, 5, 6}))
	assert.False(t, ContainSlice([]float32{1.1, 2.2, 3.3}, []float32{2.2, 3.31}))

	assert.True(t, ContainSlice([]int{1}, []int{1}))
	assert.True(t, ContainSlice([]int{0, 1, 2}, []int{2}))
	assert.True(t, ContainSlice([]int{0, 1, 2, 0, 1, 2, 3}, []int{0, 1, 2}))
	assert.True(t, ContainSlice([]string{"one", ""}, []string{""}))
	assert.True(t, ContainSlice([]string{"one", "two", "three"}, []string{"one", "two"}))
	assert.True(t, ContainSlice([]int64{1, 2, 3}, []int64{1, 2, 3}))
	assert.True(t, ContainSlice([]uint{0, 1, 1, 1, 1}, []uint{1}))
}

func Test_IndexOfSlice(t *testing.T) {
	assert.Equal(t, -1, IndexOfSlice([]int{}, nil))
	assert.Equal(t, -1, IndexOfSlice([]string{"one"}, []string{}))
	assert.Equal(t, -1, IndexOfSlice([]string{"one", "two"}, []string{"Two"}))
	assert.Equal(t, -1, IndexOfSlice([]int64{1, 2, 3}, []int64{1, 2, 3, 4}))
	assert.Equal(t, -1, IndexOfSlice([]uint{0, 1, 2, 3, 4, 5}, []uint{3, 4, 5, 6}))
	assert.Equal(t, -1, IndexOfSlice([]float32{1.1, 2.2, 3.3}, []float32{2.2, 3.31}))

	assert.Equal(t, 0, IndexOfSlice([]int{1}, []int{1}))
	assert.Equal(t, 2, IndexOfSlice([]int{0, 1, 2}, []int{2}))
	assert.Equal(t, 0, IndexOfSlice([]int{0, 1, 2, 0, 1, 2, 3}, []int{0, 1, 2}))
	assert.Equal(t, 1, IndexOfSlice([]string{"one", ""}, []string{""}))
	assert.Equal(t, 0, IndexOfSlice([]string{"one", "two", "three"}, []string{"one", "two"}))
	assert.Equal(t, 0, IndexOfSlice([]int64{1, 2, 3}, []int64{1, 2, 3}))
	assert.Equal(t, 1, IndexOfSlice([]uint{0, 1, 1, 1, 1}, []uint{1}))
}

func Test_LastIndexOfSlice(t *testing.T) {
	assert.Equal(t, -1, LastIndexOfSlice([]int{}, nil))
	assert.Equal(t, -1, LastIndexOfSlice([]string{"one"}, []string{}))
	assert.Equal(t, -1, LastIndexOfSlice([]string{"one", "two"}, []string{"Two"}))
	assert.Equal(t, -1, LastIndexOfSlice([]int64{1, 2, 3}, []int64{1, 2, 3, 4}))
	assert.Equal(t, -1, LastIndexOfSlice([]uint{0, 1, 2, 3, 4, 5}, []uint{3, 4, 5, 6}))
	assert.Equal(t, -1, LastIndexOfSlice([]float32{1.1, 2.2, 3.3}, []float32{2.2, 3.31}))

	assert.Equal(t, 0, LastIndexOfSlice([]int{1}, []int{1}))
	assert.Equal(t, 2, LastIndexOfSlice([]int{0, 1, 2}, []int{2}))
	assert.Equal(t, 3, LastIndexOfSlice([]int{0, 1, 2, 0, 1, 2, 3}, []int{0, 1, 2}))
	assert.Equal(t, 2, LastIndexOfSlice([]string{"", "one", ""}, []string{""}))
	assert.Equal(t, 0, LastIndexOfSlice([]string{"one", "two", "three"}, []string{"one", "two"}))
	assert.Equal(t, 0, LastIndexOfSlice([]int64{1, 2, 3}, []int64{1, 2, 3}))
	assert.Equal(t, 4, LastIndexOfSlice([]uint{0, 1, 1, 1, 1}, []uint{1}))
}

func Test_SubSlice(t *testing.T) {
	assert.Equal(t, []int{}, SubSlice([]int{}, 0, 100))
	assert.Equal(t, []int{}, SubSlice([]int{1, 2, 3}, 10, 100))
	assert.Equal(t, []int{2, 3}, SubSlice([]int{1, 2, 3}, 1, 100))
	assert.Equal(t, []int{3}, SubSlice([]int{1, 2, 3}, -1, 100))
	assert.Equal(t, []int{2, 3}, SubSlice([]int{1, 2, 3}, -1, -3))
}
