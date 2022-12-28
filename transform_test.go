package gofn

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapSlice(t *testing.T) {
	// Nil/Empty maps
	assert.Equal(t, []float32{}, MapSlice[int, float32](nil, func(v int) float32 { return float32(v) }))
	assert.Equal(t, []float32{}, MapSlice([]int{}, func(v int) float32 { return float32(v) }))

	// []int -> []int
	assert.Equal(t, []int{2, 4, 6}, MapSlice([]int{1, 2, 3}, func(v int) int { return v * 2 }))
	// []int -> []float32
	assert.Equal(t, []float32{2, 4, 6}, MapSlice([]int{1, 2, 3}, func(v int) float32 { return float32(v * 2) }))
}

func Test_MapSliceEx(t *testing.T) {
	// Nil/Empty maps
	assert.Equal(t, []float32{}, MapSliceEx[int, float32](nil, func(v int) (float32, bool) { return float32(v), true }))
	assert.Equal(t, []float32{}, MapSliceEx([]int{}, func(v int) (float32, bool) { return float32(v), true }))

	// []int -> []int, and omit all negatives
	assert.Equal(t, []int{1, 3}, MapSliceEx([]int{1, -2, 3, -4}, func(v int) (int, bool) { return v, v >= 0 }))
	// []string -> []string, and omit all empty elements
	assert.Equal(t, []string{"a", "b", "c"}, MapSliceEx([]string{"a", "", "b", "", "c"},
		func(v string) (string, bool) { return v, len(v) > 0 }))
}

func Test_MapSliceToMap(t *testing.T) {
	// Nil/Empty maps
	assert.Equal(t, map[int]bool{}, MapSliceToMap[int, int, bool](nil, func(v int) (int, bool) { return v, v > 0 }))
	assert.Equal(t, map[int]bool{}, MapSliceToMap([]int{}, func(v int) (int, bool) { return v, v > 0 }))

	// []int -> map[int]string
	assert.Equal(t, map[int]string{1: "2", 2: "4", 3: "6"}, MapSliceToMap([]int{1, 2, 3},
		func(v int) (int, string) { return v, strconv.Itoa(v * 2) }))
}

func Test_MapSliceToMapEx(t *testing.T) {
	// Nil/Empty maps
	assert.Equal(t, map[int]bool{}, MapSliceToMapEx[int, int, bool](nil,
		func(v int) (int, bool, bool) { return v, v > 0, v != 0 }))
	assert.Equal(t, map[int]bool{}, MapSliceToMapEx([]int{},
		func(v int) (int, bool, bool) { return v, v > 0, v != 0 }))

	// []int -> map[int]string, and omit all zeros
	assert.Equal(t, map[int]string{1: "2", -2: "-4", -4: "-8"}, MapSliceToMapEx([]int{0, 1, -2, 0, -4},
		func(v int) (int, string, bool) { return v, strconv.Itoa(v * 2), v != 0 }))
}
