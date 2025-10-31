package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsUnique(t *testing.T) {
	assert.True(t, IsUnique[int, []int](nil))
	assert.True(t, IsUnique([]int{}))
	assert.True(t, IsUnique([]string{"one"}))
	assert.True(t, IsUnique([]string{"one", "two", "One", "Two"}))
	assert.True(t, IsUnique([]float32{1.1, 2.2, 3.3, 1.11}))

	assert.False(t, IsUnique([]int{1, 2, 3, 1, 2}))
	assert.False(t, IsUnique([]string{"one", "two", "one"}))
	assert.False(t, IsUnique([]float32{1.1, 2.2, 1.100}))

	type st struct {
		I int
		S string
	}
	assert.True(t, IsUnique([]st{{1, "one"}, {2, "two"}, {3, "three"}}))
	assert.True(t, IsUnique([]st{{1, "one"}, {1, "One"}}))
	assert.False(t, IsUnique([]st{{1, "one"}, {2, "two"}, {1, "one"}}))
}

// nolint: forcetypeassert
func Test_IsUniqueBy(t *testing.T) {
	assert.True(t, IsUniqueBy[int, int, []int](nil, nil))
	assert.True(t, IsUniqueBy([]int{}, func(v int) int { return v }))
	assert.True(t, IsUniqueBy([]any{"one"},
		func(v any) string { return v.(string) }))
	assert.True(t, IsUniqueBy([]any{"one", "two", "One", "Two"},
		func(v any) string { return v.(string) }))
	assert.True(t, IsUniqueBy([]any{1.1, 2.2, 3.3, 1.11},
		func(v any) float64 { return v.(float64) }))

	assert.False(t, IsUniqueBy([]any{1, 2, 3, 1, 2},
		func(v any) int { return v.(int) }))
	assert.False(t, IsUniqueBy([]any{"one", "two", "one"},
		func(v any) string { return v.(string) }))
	assert.False(t, IsUniqueBy([]any{1.1, 2.2, 1.100},
		func(v any) float64 { return v.(float64) }))

	type st struct {
		I int
		S string
	}
	assert.True(t, IsUniqueBy([]st{{1, "one"}, {2, "two"}, {3, "three"}}, func(v st) int { return v.I }))
	assert.False(t, IsUniqueBy([]st{{1, "one"}, {1, "One"}}, func(v st) int { return v.I }))
	assert.False(t, IsUniqueBy([]st{{1, "one"}, {2, "two"}, {1, "one"}}, func(v st) int { return v.I }))
}

func Test_FindUniques(t *testing.T) {
	assert.Equal(t, []int{}, FindUniques([]int{}))
	assert.Equal(t, []int{0}, FindUniques([]int{0}))
	assert.Equal(t, []int{0, 1, 2, 3}, FindUniques([]int{0, 1, 2, 3}))
	assert.Equal(t, []int{0, 3}, FindUniques([]int{0, 1, 2, 3, 2, 1, 1}))
	assert.Equal(t, []int{}, FindUniques([]int{0, 1, 2, 3, 2, 1, 0, 0, 3}))
}

//nolint:forcetypeassert
func Test_FindUniquesBy(t *testing.T) {
	assert.Equal(t, []any{}, FindUniquesBy([]any{}, func(any) int { return 0 }))
	assert.Equal(t, []any{0}, FindUniquesBy([]any{0}, func(v any) int { return v.(int) }))
	assert.Equal(t, []any{0, 1, 2}, FindUniquesBy([]any{0, 1, 2}, func(v any) int { return v.(int) }))
	assert.Equal(t, []any{0, 3}, FindUniquesBy([]any{0, 1, 2, 3, 2, 1, 1}, func(v any) int { return v.(int) }))
	assert.Equal(t, []any{}, FindUniquesBy([]any{0, 1, 2, 3, 2, 1, 0, 0, 3}, func(v any) int { return v.(int) }))
}

func Test_FindDuplicates(t *testing.T) {
	assert.Equal(t, []int{}, FindDuplicates([]int{}))
	assert.Equal(t, []int{}, FindDuplicates([]int{0}))
	assert.Equal(t, []int{}, FindDuplicates([]int{0, 1, 2, 3}))
	assert.Equal(t, []int{1, 2}, FindDuplicates([]int{0, 1, 2, 3, 2, 1, 1}))
	assert.Equal(t, []int{0, 1, 2, 3}, FindDuplicates([]int{0, 1, 2, 3, 2, 1, 0, 0, 3}))
}

//nolint:forcetypeassert
func Test_FindDuplicatesBy(t *testing.T) {
	assert.Equal(t, []any{}, FindDuplicatesBy([]any{}, func(any) int { return 0 }))
	assert.Equal(t, []any{}, FindDuplicatesBy([]any{0}, func(v any) int { return v.(int) }))
	assert.Equal(t, []any{}, FindDuplicatesBy([]any{0, 1, 2, 3}, func(v any) int { return v.(int) }))
	assert.Equal(t, []any{1, 2}, FindDuplicatesBy([]any{0, 1, 2, 3, 2, 1, 1}, func(v any) int { return v.(int) }))
	assert.Equal(t, []any{0, 1, 2, 3},
		FindDuplicatesBy([]any{0, 1, 2, 3, 2, 1, 0, 0, 3}, func(v any) int { return v.(int) }))
}

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
