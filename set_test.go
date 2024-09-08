package gofn

import (
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
