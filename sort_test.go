package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sort(t *testing.T) {
	assert.Equal(t, []int{}, Sort([]int{}))
	assert.Equal(t, []uint32{1}, Sort([]uint32{1}))
	assert.Equal(t, []int8{-10, 0, 10, 11, 120}, Sort([]int8{0, -10, 120, 11, 10}))
	assert.Equal(t, []float64{-10.3, 0, 10.55, 11.1, 120.120}, Sort([]float64{0, -10.3, 120.120, 11.1, 10.55}))
}

func Test_SortDesc(t *testing.T) {
	assert.Equal(t, []int{}, SortDesc([]int{}))
	assert.Equal(t, []uint32{1}, SortDesc([]uint32{1}))
	assert.Equal(t, []int8{120, 11, 10, 0, -10}, SortDesc([]int8{0, -10, 120, 11, 10}))
	assert.Equal(t, []float64{120.120, 11.1, 10.55, 0, -10.3}, SortDesc([]float64{0, -10.3, 120.120, 11.1, 10.55}))
}

func Test_SortEx(t *testing.T) {
	s := []int{0, -3, 1, 2, -2}
	assert.Equal(t, []int{-3, -2, 0, 1, 2}, SortEx(s, func(i, j int) bool { return s[i] < s[j] }))
}

func Test_SortStable(t *testing.T) {
	assert.Equal(t, []int{}, SortStable([]int{}))
	assert.Equal(t, []uint32{1}, SortStable([]uint32{1}))
	assert.Equal(t, []int8{-10, 0, 10, 11, 120}, SortStable([]int8{0, -10, 120, 11, 10}))
	assert.Equal(t, []float64{-10.3, 0, 10.55, 11.1, 120.120}, SortStable([]float64{0, -10.3, 120.120, 11.1, 10.55}))
}

func Test_SortStableDesc(t *testing.T) {
	assert.Equal(t, []int{}, SortStableDesc([]int{}))
	assert.Equal(t, []uint32{1}, SortStableDesc([]uint32{1}))
	assert.Equal(t, []int8{120, 11, 10, 0, -10}, SortStableDesc([]int8{0, -10, 120, 11, 10}))
	assert.Equal(t, []float64{120.120, 11.1, 10.55, 0, -10.3}, SortStableDesc([]float64{0, -10.3, 120.120, 11.1, 10.55}))
}

func Test_SortStableEx(t *testing.T) {
	s := []int{0, -3, 1, 2, -2}
	assert.Equal(t, []int{-3, -2, 0, 1, 2}, SortStableEx(s, func(i, j int) bool { return s[i] < s[j] }))
}

func Test_IsSorted(t *testing.T) {
	assert.True(t, IsSorted([]int{}))
	assert.True(t, IsSorted([]uint32{1}))
	assert.True(t, IsSorted([]int8{-10, 0, 10, 11, 120}))
	assert.True(t, IsSorted([]float64{-10.3, 0, 10.55, 11.1, 120.120}))
}

func Test_IsSortedDesc(t *testing.T) {
	assert.True(t, IsSortedDesc([]int{}))
	assert.True(t, IsSortedDesc([]uint32{1}))
	assert.True(t, IsSortedDesc([]int8{120, 11, 10, 0, -10}))
	assert.True(t, IsSortedDesc([]float64{120.120, 11.1, 10.55, 0, -10.3}))
}
