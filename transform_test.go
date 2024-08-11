package gofn

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapSlice(t *testing.T) {
	// Nil/Empty maps
	assert.Equal(t, []float32{}, MapSlice[int, float32]([]int(nil), func(v int) float32 { return float32(v) }))
	assert.Equal(t, []float32{}, MapSlice([]int{}, func(v int) float32 { return float32(v) }))

	// []int -> []int
	assert.Equal(t, []int{2, 4, 6}, MapSlice([]int{1, 2, 3}, func(v int) int { return v * 2 }))
	// []int -> []float32
	assert.Equal(t, []float32{2, 4, 6}, MapSlice([]int{1, 2, 3}, func(v int) float32 { return float32(v * 2) }))
}

func Test_MapSliceEx(t *testing.T) {
	// Nil/Empty maps
	r1, e := MapSliceEx[int, float32]([]int(nil), func(v int) (float32, error) { return float32(v), nil })
	assert.Nil(t, e)
	assert.Equal(t, []float32{}, r1)
	r2, e := MapSliceEx([]int{}, func(v int) (float32, error) { return float32(v), nil })
	assert.Nil(t, e)
	assert.Equal(t, []float32{}, r2)

	// []string -> []int
	r3, e := MapSliceEx([]string{"1", "2", "3"}, ParseInt[int])
	assert.Nil(t, e)
	assert.Equal(t, []int{1, 2, 3}, r3)

	// []string -> []int with error
	_, e = MapSliceEx([]string{"1", "a", "3"}, ParseInt[int])
	assert.ErrorIs(t, e, strconv.ErrSyntax)
}

func Test_MapSliceToMap(t *testing.T) {
	// Nil/Empty maps
	assert.Equal(t, map[int]bool{}, MapSliceToMap[int, int, bool]([]int(nil), func(v int) (int, bool) { return v, v > 0 }))
	assert.Equal(t, map[int]bool{}, MapSliceToMap([]int{}, func(v int) (int, bool) { return v, v > 0 }))

	// []int -> map[int]string
	assert.Equal(t, map[int]string{1: "2", 2: "4", 3: "6"}, MapSliceToMap([]int{1, 2, 3},
		func(v int) (int, string) { return v, strconv.Itoa(v * 2) }))
}

func Test_MapSliceToMapEx(t *testing.T) {
	// Nil/Empty maps
	r1, e := MapSliceToMapEx[int, int, bool]([]int(nil), func(v int) (int, bool, error) { return v, v > 0, nil })
	assert.Nil(t, e)
	assert.Equal(t, map[int]bool{}, r1)
	r2, e := MapSliceToMapEx([]int{}, func(v int) (int, bool, error) { return v, v > 0, nil })
	assert.Nil(t, e)
	assert.Equal(t, map[int]bool{}, r2)

	// []string -> []int
	r3, e := MapSliceToMapEx([]string{"1", "2", "3"}, func(v string) (string, int, error) {
		u, e := ParseInt[int](v)
		return v, u, e
	})
	assert.Nil(t, e)
	assert.Equal(t, map[string]int{"1": 1, "2": 2, "3": 3}, r3)

	// []string -> []int with error
	_, e = MapSliceToMapEx([]string{"1", "x", "3"}, func(v string) (string, int, error) {
		u, e := ParseInt[int](v)
		return v, u, e
	})
	assert.ErrorIs(t, e, strconv.ErrSyntax)
}
