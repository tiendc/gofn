package gofn

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToSet(t *testing.T) {
	assert.Equal(t, []int{}, ToSet([]int{}))
	assert.Equal(t, []string{"one"}, ToSet([]string{"one"}))
	assert.Equal(t, []string{"one", "two", "Two"}, ToSet([]string{"one", "two", "one", "Two"}))
	assert.Equal(t, []int{1, 2, 3}, ToSet([]int{1, 2, 3, 1, 2}))
	assert.Equal(t, []float32{float32(1.1), float32(2.2), float32(3.3), float32(1.11)},
		ToSet([]float32{1.1, 1.1, 2.2, 3.3, 1.11}))

	type st struct {
		I int
		S string
	}
	assert.Equal(t, []st{{1, "one"}, {2, "two"}},
		ToSet([]st{{1, "one"}, {2, "two"}, {1, "one"}}))
}

// nolint: forcetypeassert
func Test_ToSetBy(t *testing.T) {
	// Comparable types
	assert.Equal(t, []int{},
		ToSetBy([]int{}, func(t int) int { return t }))
	assert.Equal(t, []string{"one"},
		ToSetBy([]string{"one"}, func(t string) string { return t }))
	assert.Equal(t, []string{"one", "two", "Two"},
		ToSetBy([]string{"one", "two", "one", "Two"}, func(t string) string { return t }))
	assert.Equal(t, []int{1, 2, 3},
		ToSetBy([]int{1, 2, 3, 1, 2}, func(t int) int { return t }))
	assert.Equal(t, []float32{float32(1.1), float32(2.2), float32(3.3), float32(1.11)},
		ToSetBy([]float32{1.1, 1.1, 2.2, 3.3, 1.11}, func(t float32) float32 { return t }))

	// Incomparable types
	assert.Equal(t, []any{},
		ToSetBy([]any{}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{"one"},
		ToSetBy([]any{"one"}, func(t any) string { return t.(string) }))
	assert.Equal(t, []any{"one", "two", "Two"},
		ToSetBy([]any{"one", "two", "one", "Two"}, func(t any) string { return t.(string) }))
	assert.Equal(t, []any{1, 2, 3},
		ToSetBy([]any{1, 2, 3, 1, 2}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1.1, 2.2, 3.3, 1.11},
		ToSetBy([]any{1.1, 1.1, 2.2, 3.3, 1.11}, func(t any) float64 { return t.(float64) }))
}

// nolint: forcetypeassert
func Test_ToSetPred_Deprecated(t *testing.T) {
	// Comparable types
	assert.Equal(t, []int{1, 2, 3},
		ToSetPred([]int{1, 2, 3, 1, 2}, func(t int) int { return t }))
	assert.Equal(t, []float32{float32(1.1), float32(2.2), float32(3.3), float32(1.11)},
		ToSetPred([]float32{1.1, 1.1, 2.2, 3.3, 1.11}, func(t float32) float32 { return t }))

	// Incomparable types
	assert.Equal(t, []any{1, 2, 3},
		ToSetPred([]any{1, 2, 3, 1, 2}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1.1, 2.2, 3.3, 1.11},
		ToSetPred([]any{1.1, 1.1, 2.2, 3.3, 1.11}, func(t any) float64 { return t.(float64) }))
}

// nolint: forcetypeassert
func Test_ToSetByReverse(t *testing.T) {
	// Comparable types
	assert.Equal(t, []int{},
		ToSetByReverse([]int{}, func(t int) int { return t }))
	assert.Equal(t, []string{"one"},
		ToSetByReverse([]string{"one"}, func(t string) string { return t }))
	assert.Equal(t, []string{"Two", "one", "two"},
		ToSetByReverse([]string{"one", "two", "one", "Two"}, func(t string) string { return t }))
	assert.Equal(t, []int{2, 1, 3},
		ToSetByReverse([]int{1, 2, 3, 1, 2}, func(t int) int { return t }))
	assert.Equal(t, []float32{float32(1.1), float32(2.2), float32(3.3), float32(1.11)},
		ToSetByReverse([]float32{1.11, 3.3, 2.2, 1.1}, func(t float32) float32 { return t }))

	// Incomparable types
	assert.Equal(t, []any{},
		ToSetByReverse([]any{}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{"one"},
		ToSetByReverse([]any{"one"}, func(t any) string { return t.(string) }))
	assert.Equal(t, []any{"Two", "one", "two"},
		ToSetByReverse([]any{"one", "two", "one", "Two"}, func(t any) string { return t.(string) }))
	assert.Equal(t, []any{2, 1, 3},
		ToSetByReverse([]any{1, 2, 3, 1, 2}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1.11, 3.3, 2.2, 1.1},
		ToSetByReverse([]any{1.1, 1.1, 2.2, 3.3, 1.11}, func(t any) float64 { return t.(float64) }))
}

// nolint: forcetypeassert
func Test_ToSetPredReverse_Deprecated(t *testing.T) {
	// Comparable types
	assert.Equal(t, []int{2, 1, 3},
		ToSetPredReverse([]int{1, 2, 3, 1, 2}, func(t int) int { return t }))
	assert.Equal(t, []float32{float32(1.1), float32(2.2), float32(3.3), float32(1.11)},
		ToSetPredReverse([]float32{1.11, 3.3, 2.2, 1.1}, func(t float32) float32 { return t }))

	// Incomparable types
	assert.Equal(t, []any{2, 1, 3},
		ToSetPredReverse([]any{1, 2, 3, 1, 2}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1.11, 3.3, 2.2, 1.1},
		ToSetPredReverse([]any{1.1, 1.1, 2.2, 3.3, 1.11}, func(t any) float64 { return t.(float64) }))
}

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

func Test_MapSliceToMapKeys(t *testing.T) {
	// Nil/Empty maps
	assert.Equal(t, map[int]struct{}{}, MapSliceToMapKeys[int]([]int(nil), struct{}{}))
	assert.Equal(t, map[int]int{}, MapSliceToMapKeys([]int{}, 0))

	// []int -> map[int]string
	assert.Equal(t, map[int]string{1: "x", 2: "x", 3: "x"}, MapSliceToMapKeys([]int{1, 2, 3, 2}, "x"))
}
