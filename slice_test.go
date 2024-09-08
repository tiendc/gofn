package gofn

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Equal(t *testing.T) {
	assert.True(t, Equal([]int{}, []int{}))
	assert.True(t, Equal([]int{}, nil))
	assert.True(t, Equal(nil, []int{}))
	assert.True(t, Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.True(t, Equal([]string{"3", "1", "2"}, []string{"3", "1", "2"}))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, Equal([]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{1, "1"}, {2, "2"}, {3, "3"}}))

	assert.False(t, Equal([]int{}, []int{1}))
	assert.False(t, Equal([]int{1}, nil))
	assert.False(t, Equal([]int{1, 2, 3}, []int{1, 2, 3, 4}))
	assert.False(t, Equal([]int{1, 2, 3}, []int{3, 2, 1}))
	assert.False(t, Equal([]string{"3", "1", "2"}, []string{"1", "2", "3"}))
	assert.False(t, Equal([]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{1, "1"}, {2, "2"}}))
	assert.False(t, Equal([]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{3, "3"}, {1, "1"}, {2, "2"}}))

	// Derived types
	type SI []int
	assert.True(t, Equal([]int{1, 2, 3}, SI{1, 2, 3}))
	assert.False(t, Equal(SI{1, 2}, []int{1, 2, 3}))
}

// nolint: forcetypeassert
func Test_EqualBy(t *testing.T) {
	assert.True(t, EqualBy([]any{}, []any{},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualBy([]any{}, nil,
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualBy(nil, []any{},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualBy([]any{1, 2, 3}, []any{1, 2, 3},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualBy([]any{"3", "1", "2"}, []any{"3", "1", "2"},
		func(a, b any) bool { return a.(string) == b.(string) }))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, EqualBy(
		[]any{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(a, b any) bool { return a.(St) == b.(St) }))

	assert.False(t, EqualBy([]any{}, []any{1},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualBy([]any{1}, nil,
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualBy([]any{1, 2, 3}, []any{1, 2, 3, 4},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualBy([]any{1, 2, 3}, []any{3, 2, 1},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualBy([]any{"3", "1", "2"}, []any{"1", "2", "3"},
		func(a, b any) bool { return a.(string) == b.(string) }))
	assert.False(t, EqualBy(
		[]any{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []any{St{1, "1"}, St{2, "2"}},
		func(a, b any) bool { return a.(St) == b.(St) }))
	assert.False(t, EqualBy(
		[]any{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []any{St{3, "3"}, St{2, "2"}, St{1, "1"}},
		func(a, b any) bool { return a.(St) == b.(St) }))
}

// nolint: forcetypeassert
func Test_EqualPred_Deprecated(t *testing.T) {
	assert.True(t, EqualPred([]any{1, 2, 3}, []any{1, 2, 3},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualPred([]any{1, 2, 3}, []any{3, 2, 1},
		func(a, b any) bool { return a.(int) == b.(int) }))
}

func Test_EqualByPtr(t *testing.T) {
	assert.True(t, EqualByPtr([]int{}, []int{},
		func(a, b *int) bool { return *a == *b }))
	assert.True(t, EqualByPtr([]int{}, nil,
		func(a, b *int) bool { return *a == *b }))
	assert.True(t, EqualByPtr(nil, []int{},
		func(a, b *int) bool { return *a == *b }))
	assert.True(t, EqualByPtr([]int{1, 2, 3}, []int{1, 2, 3},
		func(a, b *int) bool { return *a == *b }))
	assert.True(t, EqualByPtr([]string{"3", "1", "2"}, []string{"3", "1", "2"},
		func(a, b *string) bool { return *a == *b }))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, EqualByPtr(
		[]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(a, b *St) bool { return *a == *b }))

	assert.False(t, EqualByPtr([]int{}, []int{1},
		func(a, b *int) bool { return *a == *b }))
	assert.False(t, EqualByPtr([]int{1}, nil,
		func(a, b *int) bool { return *a == *b }))
	assert.False(t, EqualByPtr([]int{1, 2, 3}, []int{1, 2, 3, 4},
		func(a, b *int) bool { return *a == *b }))
	assert.False(t, EqualByPtr([]int{1, 2, 3}, []int{3, 2, 1},
		func(a, b *int) bool { return *a == *b }))
	assert.False(t, EqualByPtr([]string{"3", "1", "2"}, []string{"1", "2", "3"},
		func(a, b *string) bool { return *a == *b }))
	assert.False(t, EqualByPtr(
		[]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{1, "1"}, {2, "2"}},
		func(a, b *St) bool { return *a == *b }))
	assert.False(t, EqualByPtr(
		[]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{3, "3"}, {2, "2"}, {1, "1"}},
		func(a, b *St) bool { return *a == *b }))
}

func Test_EqualPredPtr_Deprecated(t *testing.T) {
	assert.True(t, EqualPredPtr([]int{1, 2, 3}, []int{1, 2, 3},
		func(a, b *int) bool { return *a == *b }))
	assert.False(t, EqualPredPtr([]int{1, 2, 3}, []int{3, 2, 1},
		func(a, b *int) bool { return *a == *b }))
}

func Test_ContentEqual(t *testing.T) {
	assert.True(t, ContentEqual([]int{}, []int{}))
	assert.True(t, ContentEqual([]int{}, nil))
	assert.True(t, ContentEqual(nil, []int{}))
	assert.True(t, ContentEqual([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.True(t, ContentEqual([]int{3, 1, 2}, []int{1, 2, 3}))
	assert.True(t, ContentEqual([]int{1, 2, 1, 3}, []int{1, 2, 3, 1}))
	assert.True(t, ContentEqual([]string{"3", "1", "2"}, []string{"1", "2", "3"}))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, ContentEqual([]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{3, "3"}, {1, "1"}, {2, "2"}}))

	assert.False(t, ContentEqual([]int{}, []int{1}))
	assert.False(t, ContentEqual([]int{1}, nil))
	assert.False(t, ContentEqual([]int{1, 2, 3}, []int{1, 2, 3, 4}))
	assert.False(t, ContentEqual([]int{1, 2, 3}, []int{1, 2, 3, 3}))
	assert.False(t, ContentEqual([]int{1, 2, 1, 3}, []int{1, 2, 2, 3}))
	assert.False(t, ContentEqual([]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{3, "3"}, {1, "1"}, {3, "3"}}))
}

// nolint: forcetypeassert
func Test_ContentEqualBy(t *testing.T) {
	assert.True(t, ContentEqualBy([]any{}, []any{},
		func(t any) int { return t.(int) }))
	assert.True(t, ContentEqualBy([]any{}, nil,
		func(t any) int { return t.(int) }))
	assert.True(t, ContentEqualBy(nil, []any{},
		func(t any) int { return t.(int) }))
	assert.True(t, ContentEqualBy([]any{1, 2, 3}, []any{1, 2, 3},
		func(t any) int { return t.(int) }))
	assert.True(t, ContentEqualBy([]any{3, 1, 2}, []any{1, 2, 3},
		func(t any) int { return t.(int) }))
	assert.True(t, ContentEqualBy([]any{"3", "1", "2"}, []any{"1", "2", "3"},
		func(t any) string { return t.(string) }))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, ContentEqualBy(
		[]any{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []any{St{3, "3"}, St{1, "1"}, St{2, "2"}},
		func(t any) St { return t.(St) }))

	assert.False(t, ContentEqualBy([]any{}, []any{1},
		func(t any) int { return t.(int) }))
	assert.False(t, ContentEqualBy([]any{1}, nil,
		func(t any) int { return t.(int) }))
	assert.False(t, ContentEqualBy([]any{1, 2, 3}, []any{1, 2, 3, 4},
		func(t any) int { return t.(int) }))
	assert.False(t, ContentEqualBy(
		[]any{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []any{St{3, "3"}, St{1, "1"}, St{3, "3"}},
		func(t any) St { return t.(St) }))
}

// nolint: forcetypeassert
func Test_ContentEqualPred_Deprecated(t *testing.T) {
	assert.True(t, ContentEqualPred([]any{1, 2, 3}, []any{1, 2, 3},
		func(t any) int { return t.(int) }))
	assert.False(t, ContentEqualPred([]any{1, 2, 3}, []any{1, 2, 3, 4},
		func(t any) int { return t.(int) }))
}

func Test_ContentEqualPtr(t *testing.T) {
	assert.True(t, ContentEqualPtr([]*int{}, []*int{}))
	assert.True(t, ContentEqualPtr([]*int{}, nil))
	assert.True(t, ContentEqualPtr(nil, []*int{}))
	i1, i2, i3 := New(1), New(2), New(3)
	assert.True(t, ContentEqualPtr([]*int{i3, i1, i2}, []*int{i1, i2, i3}))
	assert.True(t, ContentEqualPtr([]*int{i1, i2, i3, i2, i1}, []*int{i3, i1, i2, i1, i2}))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, ContentEqualPtr([]*St{{1, "1"}, {2, "2"}, {3, "3"}}, []*St{{3, "3"}, {1, "1"}, {2, "2"}}))

	assert.False(t, ContentEqualPtr([]*int{}, []*int{i1}))
	assert.False(t, ContentEqualPtr([]*int{i1}, nil))
	assert.False(t, ContentEqualPtr([]*int{i1, i2, i3}, []*int{i1, i2}))
	assert.False(t, ContentEqualPtr([]*St{{1, "1"}, {2, "2"}, {3, "3"}}, []*St{{3, "3"}, {1, "1"}, {3, "3"}}))
}

func Test_Concat(t *testing.T) {
	assert.Equal(t, []int{}, Concat[int, []int](nil, nil, nil))
	assert.Equal(t, []bool{}, Concat([]bool{}, []bool{}))
	assert.Equal(t, []float64{1.1}, Concat([]float64{}, []float64{}, []float64{1.1}))
	assert.Equal(t, []string{"", "1", "2", "3"}, Concat([]string{""}, []string{"1", "2"}, []string{}, []string{"3"}))
}

func Test_Contain(t *testing.T) {
	assert.False(t, Contain([]int{}, 1))
	assert.False(t, Contain([]string{"one"}, "One"))
	assert.False(t, Contain([]string{"one", "two"}, ""))
	assert.False(t, Contain([]int64{1, 2, 3}, 4))
	assert.False(t, Contain([]float32{1.1, 2.2, 3.3}, 3.35))

	type St struct {
		Int int
		Str string
	}
	assert.False(t, Contain([]St{{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "4"}))

	assert.True(t, Contain([]int64{1}, 1))
	assert.True(t, Contain([]uint{1, 2, 3, 1, 2, 3}, 2))
	assert.True(t, Contain([]string{"one", "two"}, "two"))
	assert.True(t, Contain([]string{"one", "two", ""}, ""))
	assert.True(t, Contain([]float64{1.1, 2.2, 3.3}, 2.2))
	assert.True(t, Contain([]St{{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "3"}))
}

// nolint: goconst, forcetypeassert
func Test_ContainBy(t *testing.T) {
	assert.False(t, ContainBy([]any{},
		func(i any) bool { return i.(int) == 1 }))
	assert.False(t, ContainBy([]any{"one"},
		func(i any) bool { return i == "One" }))
	assert.False(t, ContainBy([]any{"one", "two"},
		func(i any) bool { return i == "" }))
	assert.False(t, ContainBy([]any{1, 2, 3},
		func(i any) bool { return i == 4 }))
	assert.False(t, ContainBy([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 3.35 }))

	type St struct {
		Int int
		Str string
	}
	assert.False(t, ContainBy([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "4"} }))

	assert.True(t, ContainBy([]any{1},
		func(i any) bool { return i == 1 }))
	assert.True(t, ContainBy([]any{1, 2, 3, 1, 2, 3},
		func(i any) bool { return i == 2 }))
	assert.True(t, ContainBy([]any{"one", "two"},
		func(i any) bool { return i == "two" }))
	assert.True(t, ContainBy([]any{"one", "two", ""},
		func(i any) bool { return i == "" }))
	assert.True(t, ContainBy([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 2.2 }))
	assert.True(t, ContainBy([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "3"} }))
}

// nolint: goconst, forcetypeassert
func Test_ContainPred_Deprecated(t *testing.T) {
	assert.False(t, ContainPred([]any{1, 2, 3},
		func(i any) bool { return i == 4 }))
	assert.True(t, ContainPred([]any{1, 2, 3, 1, 2, 3},
		func(i any) bool { return i == 2 }))
}

func Test_ContainByPtr(t *testing.T) {
	assert.False(t, ContainByPtr([]int{},
		func(i *int) bool { return *i == 1 }))
	assert.False(t, ContainByPtr([]string{"one"},
		func(i *string) bool { return *i == "One" }))
	assert.False(t, ContainByPtr([]string{"one", "two"},
		func(i *string) bool { return *i == "" }))
	assert.False(t, ContainByPtr([]int{1, 2, 3},
		func(i *int) bool { return *i == 4 }))
	assert.False(t, ContainByPtr([]float32{1.1, 2.2, 3.3},
		func(i *float32) bool { return *i == 3.35 }))

	type St struct {
		Int int
		Str string
	}
	assert.False(t, ContainByPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return *i == St{3, "4"} }))

	assert.True(t, ContainByPtr([]int{1},
		func(i *int) bool { return *i == 1 }))
	assert.True(t, ContainByPtr([]int{1, 2, 3, 1, 2, 3},
		func(i *int) bool { return *i == 2 }))
	assert.True(t, ContainByPtr([]string{"one", "two"},
		func(i *string) bool { return *i == "two" }))
	assert.True(t, ContainByPtr([]string{"one", "two", ""},
		func(i *string) bool { return *i == "" }))
	assert.True(t, ContainByPtr([]float32{1.1, 2.2, 3.3},
		func(i *float32) bool { return *i == 2.2 }))
	assert.True(t, ContainByPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return *i == St{3, "3"} }))
}

func Test_ContainPredPtr_Deprecated(t *testing.T) {
	assert.False(t, ContainPredPtr([]int{1, 2, 3},
		func(i *int) bool { return *i == 4 }))
	assert.True(t, ContainPredPtr([]int{1, 2, 3, 1, 2, 3},
		func(i *int) bool { return *i == 2 }))
}

func Test_ContainAll(t *testing.T) {
	assert.False(t, ContainAll([]int{}, 0))
	assert.False(t, ContainAll([]string{"one"}, "one", "two"))
	assert.False(t, ContainAll([]string{"one", "two"}, "two", "one", ""))
	assert.False(t, ContainAll([]int64{1, 2, 3}, 3, 3, 2, 1, 0))
	assert.False(t, ContainAll([]float32{1.1, 2.2, 3.3}, 3.35, 2.2))

	assert.True(t, ContainAll([]int64{1}, 1))
	assert.True(t, ContainAll([]uint{1, 2, 3, 1, 2, 3}, 2, 1, 3, 3, 2))
	assert.True(t, ContainAll([]string{"one", "two"}, "two"))
	assert.True(t, ContainAll([]string{"one", "two", ""}, "", "two"))
	assert.True(t, ContainAll([]float64{1.1, 2.2, 3.3}, 2.2, 3.3, 1.1))

	// More than 10 items in slice
	assert.False(t, ContainAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 1, 2, 3, 0))
	assert.False(t, ContainAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 12))
	assert.True(t, ContainAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 1, 2, 3, 3, 2, 1))
	assert.True(t, ContainAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11))
	assert.True(t, ContainAll([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
		"1", "2", "3", "3", "2", "1"))
}

func Test_ContainAny(t *testing.T) {
	assert.False(t, ContainAny([]int{}, 0))
	assert.False(t, ContainAny([]string{"one"}, "One", "three"))
	assert.False(t, ContainAny([]string{"one", "two"}, ""))
	assert.False(t, ContainAny([]int64{1, 2, 3}, 0, 4, 5))
	assert.False(t, ContainAny([]float32{1.1, 2.2, 3.3}, 3.35, 2.22))

	assert.True(t, ContainAny([]int64{1}, 1, 1))
	assert.True(t, ContainAny([]uint{1, 2, 3, 1, 2, 3}, 10, 11, 1))
	assert.True(t, ContainAny([]string{"one", "two"}, "two", "three"))
	assert.True(t, ContainAny([]string{"one", "two", ""}, "", "two"))
	assert.True(t, ContainAny([]float64{1.1, 2.2, 3.3}, 2.2, 3.33, 1.11))

	// More than 10 items in slice
	assert.False(t, ContainAny([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 0, -1, 100))
	assert.False(t, ContainAny([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 13))
	assert.True(t, ContainAny([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 1, 2, 3, 3, 2, 1))
	assert.True(t, ContainAny([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 6, 100, 101, 102))
	assert.True(t, ContainAny([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}, "100", "0", "1", "2"))
}

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

// nolint: forcetypeassert
func Test_IsUniquePred_Deprecated(t *testing.T) {
	assert.True(t, IsUniquePred([]any{1.1, 2.2, 3.3, 1.11},
		func(v any) float64 { return v.(float64) }))
	assert.False(t, IsUniquePred([]any{1, 2, 3, 1, 2},
		func(v any) int { return v.(int) }))
}

func Test_Find(t *testing.T) {
	_, found := Find([]any{}, func(i any) bool { return i == 1 })
	assert.False(t, found)
	_, found = Find([]any{"one"}, func(i any) bool { return i == "One" })
	assert.False(t, found)
	_, found = Find([]any{"one", "two"}, func(i any) bool { return i == "" })
	assert.False(t, found)
	_, found = Find([]any{1, 2, 3}, func(i any) bool { return i == 4 })
	assert.False(t, found)
	_, found = Find([]any{1.1, 2.2, 3.3}, func(i any) bool { return i == 3.35 })
	assert.False(t, found)

	type St struct {
		Int int
		Str string
	}
	_, found = Find([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i St) bool { return i.Int == 4 })
	assert.False(t, found)

	v1, found := Find([]any{1}, func(i any) bool { return i == 1 })
	assert.True(t, found)
	assert.Equal(t, 1, v1)
	v2, found := Find([]any{1, 2, 3, 1, 2, 3}, func(i any) bool { return i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
	v3, found := Find([]any{"one", "two"}, func(i any) bool { return i == "two" })
	assert.True(t, found)
	assert.Equal(t, "two", v3)
	v4, found := Find([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "3"} })
	assert.True(t, found)
	assert.Equal(t, St{3, "3"}, v4)
	v5, found := Find([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i St) bool { return i.Int == 2 })
	assert.True(t, found)
	assert.Equal(t, St{2, "2"}, v5)
}

func Test_FindPred_Deprecated(t *testing.T) {
	_, found := FindPred([]any{1.1, 2.2, 3.3}, func(i any) bool { return i == 3.35 })
	assert.False(t, found)

	v2, found := FindPred([]any{1, 2, 3, 1, 2, 3}, func(i any) bool { return i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
}

func Test_FindPtr(t *testing.T) {
	_, found := FindPtr([]any{}, func(i *any) bool { return *i == 1 })
	assert.False(t, found)
	_, found = FindPtr([]any{"one"}, func(i *any) bool { return *i == "One" })
	assert.False(t, found)
	_, found = FindPtr([]any{"one", "two"}, func(i *any) bool { return *i == "" })
	assert.False(t, found)
	_, found = FindPtr([]any{1, 2, 3}, func(i *any) bool { return *i == 4 })
	assert.False(t, found)
	_, found = FindPtr([]any{1.1, 2.2, 3.3}, func(i *any) bool { return *i == 3.35 })
	assert.False(t, found)

	type St struct {
		Int int
		Str string
	}
	_, found = FindPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return i.Int == 4 })
	assert.False(t, found)

	v1, found := FindPtr([]any{1}, func(i *any) bool { return *i == 1 })
	assert.True(t, found)
	assert.Equal(t, 1, v1)
	v2, found := FindPtr([]any{1, 2, 3, 1, 2, 3}, func(i *any) bool { return *i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
	v3, found := FindPtr([]any{"one", "two"}, func(i *any) bool { return *i == "two" })
	assert.True(t, found)
	assert.Equal(t, "two", v3)
	v4, found := FindPtr([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i *any) bool { return *i == St{3, "3"} })
	assert.True(t, found)
	assert.Equal(t, St{3, "3"}, v4)
	v5, found := FindPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return i.Int == 2 })
	assert.True(t, found)
	assert.Equal(t, St{2, "2"}, v5)
}

func Test_FindPredPtr_Deprecated(t *testing.T) {
	_, found := FindPredPtr([]any{1.1, 2.2, 3.3}, func(i *any) bool { return *i == 3.35 })
	assert.False(t, found)

	v2, found := FindPredPtr([]any{1, 2, 3, 1, 2, 3}, func(i *any) bool { return *i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
}

func Test_FindLast(t *testing.T) {
	_, found := FindLast([]any{}, func(i any) bool { return i == 1 })
	assert.False(t, found)
	_, found = FindLast([]any{"one"}, func(i any) bool { return i == "One" })
	assert.False(t, found)
	_, found = FindLast([]any{"one", "two"}, func(i any) bool { return i == "" })
	assert.False(t, found)
	_, found = FindLast([]any{1, 2, 3}, func(i any) bool { return i == 4 })
	assert.False(t, found)
	_, found = FindLast([]any{1.1, 2.2, 3.3}, func(i any) bool { return i == 3.35 })
	assert.False(t, found)

	type St struct {
		Int int
		Str string
	}
	_, found = FindLast([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i St) bool { return i.Int == 4 })
	assert.False(t, found)

	v1, found := FindLast([]any{1}, func(i any) bool { return i == 1 })
	assert.True(t, found)
	assert.Equal(t, 1, v1)
	v2, found := FindLast([]any{1, 2, 3, 1, 2, 3}, func(i any) bool { return i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
	v3, found := FindLast([]any{"one", "two"}, func(i any) bool { return i == "two" })
	assert.True(t, found)
	assert.Equal(t, "two", v3)
	v4, found := FindLast([]any{"one", "", "two", ""}, func(i any) bool { return i == "" })
	assert.True(t, found)
	assert.Equal(t, "", v4)
	v5, found := FindLast([]any{1.1, 1.1, 2.2, 3.3}, func(i any) bool { return i == 1.1 })
	assert.True(t, found)
	assert.Equal(t, 1.1, v5)
	v6, found := FindLast([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i St) bool { return i.Int == 3 })
	assert.True(t, found)
	assert.Equal(t, St{3, "3"}, v6)
}

func Test_FindLastPred_Deprecated(t *testing.T) {
	_, found := FindLastPred([]any{1, 2, 3}, func(i any) bool { return i == 4 })
	assert.False(t, found)

	v2, found := FindLastPred([]any{1, 2, 3, 1, 2, 3}, func(i any) bool { return i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
}

func Test_FindLastPtr(t *testing.T) {
	_, found := FindLastPtr([]any{}, func(i *any) bool { return *i == 1 })
	assert.False(t, found)
	_, found = FindLastPtr([]any{"one"}, func(i *any) bool { return *i == "One" })
	assert.False(t, found)
	_, found = FindLastPtr([]any{"one", "two"}, func(i *any) bool { return *i == "" })
	assert.False(t, found)
	_, found = FindLastPtr([]any{1, 2, 3}, func(i *any) bool { return *i == 4 })
	assert.False(t, found)
	_, found = FindLastPtr([]any{1.1, 2.2, 3.3}, func(i *any) bool { return *i == 3.35 })
	assert.False(t, found)

	type St struct {
		Int int
		Str string
	}
	_, found = FindLastPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return i.Int == 4 })
	assert.False(t, found)

	v1, found := FindLastPtr([]any{1}, func(i *any) bool { return *i == 1 })
	assert.True(t, found)
	assert.Equal(t, 1, v1)
	v2, found := FindLastPtr([]any{1, 2, 3, 1, 2, 3}, func(i *any) bool { return *i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
	v3, found := FindLastPtr([]any{"one", "two"}, func(i *any) bool { return *i == "two" })
	assert.True(t, found)
	assert.Equal(t, "two", v3)
	v4, found := FindLastPtr([]any{"one", "", "two", ""}, func(i *any) bool { return *i == "" })
	assert.True(t, found)
	assert.Equal(t, "", v4)
	v5, found := FindLastPtr([]any{1.1, 1.1, 2.2, 3.3}, func(i *any) bool { return *i == 1.1 })
	assert.True(t, found)
	assert.Equal(t, 1.1, v5)
	v6, found := FindLastPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return i.Int == 3 })
	assert.True(t, found)
	assert.Equal(t, St{3, "3"}, v6)
}

func Test_FindLastPredPtr(t *testing.T) {
	_, found := FindLastPredPtr([]any{1.1, 2.2, 3.3}, func(i *any) bool { return *i == 3.35 })
	assert.False(t, found)

	v2, found := FindLastPredPtr([]any{1, 2, 3, 1, 2, 3}, func(i *any) bool { return *i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
}

func Test_IndexOf(t *testing.T) {
	assert.Equal(t, -1, IndexOf([]int{}, 1))
	assert.Equal(t, -1, IndexOf([]string{"one"}, "One"))
	assert.Equal(t, -1, IndexOf([]string{"one", "two"}, ""))
	assert.Equal(t, -1, IndexOf([]int64{1, 2, 3}, 4))
	assert.Equal(t, -1, IndexOf([]float32{1.1, 2.2, 3.3}, 3.35))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, -1, IndexOf([]St{{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "4"}))

	assert.Equal(t, 0, IndexOf([]int64{1}, 1))
	assert.Equal(t, 1, IndexOf([]uint{1, 2, 3, 1, 2, 3}, 2))
	assert.Equal(t, 1, IndexOf([]string{"one", "two"}, "two"))
	assert.Equal(t, 2, IndexOf([]string{"one", "two", ""}, ""))
	assert.Equal(t, 2, IndexOf([]float64{1.1, 2.2, 3.3}, 3.3))
	assert.Equal(t, 2, IndexOf([]St{{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "3"}))
}

func Test_IndexOfBy(t *testing.T) {
	assert.Equal(t, -1, IndexOfBy([]any{},
		func(i any) bool { return i == 1 }))
	assert.Equal(t, -1, IndexOfBy([]any{"one"},
		func(i any) bool { return i == "One" }))
	assert.Equal(t, -1, IndexOfBy([]any{"one", "two"},
		func(i any) bool { return i == "" }))
	assert.Equal(t, -1, IndexOfBy([]any{1, 2, 3},
		func(i any) bool { return i == 4 }))
	assert.Equal(t, -1, IndexOfBy([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 3.35 }))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, -1, IndexOfBy([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "4"} }))

	assert.Equal(t, 0, IndexOfBy([]any{1},
		func(i any) bool { return i == 1 }))
	assert.Equal(t, 1, IndexOfBy([]any{1, 2, 3, 1, 2, 3},
		func(i any) bool { return i == 2 }))
	assert.Equal(t, 1, IndexOfBy([]any{"one", "two"},
		func(i any) bool { return i == "two" }))
	assert.Equal(t, 2, IndexOfBy([]any{"one", "two", ""},
		func(i any) bool { return i == "" }))
	assert.Equal(t, 2, IndexOfBy([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 3.3 }))
	assert.Equal(t, 2, IndexOfBy([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "3"} }))
}

func Test_IndexOfPred_Deprecated(t *testing.T) {
	assert.Equal(t, -1, IndexOfPred([]any{1, 2, 3},
		func(i any) bool { return i == 4 }))
	assert.Equal(t, 2, IndexOfPred([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 3.3 }))
}

func Test_LastIndexOf(t *testing.T) {
	assert.Equal(t, -1, LastIndexOf([]int{}, 1))
	assert.Equal(t, -1, LastIndexOf([]string{"one"}, "One"))
	assert.Equal(t, -1, LastIndexOf([]string{"one", "two"}, ""))
	assert.Equal(t, -1, LastIndexOf([]int64{1, 2, 3}, 4))
	assert.Equal(t, -1, LastIndexOf([]float32{1.1, 2.2, 3.3}, 3.35))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, -1, LastIndexOf([]St{{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "4"}))

	assert.Equal(t, 0, LastIndexOf([]int64{1}, 1))
	assert.Equal(t, 4, LastIndexOf([]uint{1, 2, 3, 1, 2, 3}, 2))
	assert.Equal(t, 1, LastIndexOf([]string{"one", "two"}, "two"))
	assert.Equal(t, 3, LastIndexOf([]string{"one", "", "two", ""}, ""))
	assert.Equal(t, 0, LastIndexOf([]float64{1.1, 2.2, 3.3}, 1.1))
	assert.Equal(t, 2, LastIndexOf([]St{{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "3"}))
}

func Test_LastIndexOfBy(t *testing.T) {
	assert.Equal(t, -1, LastIndexOfBy([]any{},
		func(i any) bool { return i == 1 }))
	assert.Equal(t, -1, LastIndexOfBy([]any{"one"},
		func(i any) bool { return i == "One" }))
	assert.Equal(t, -1, LastIndexOfBy([]any{"one", "two"},
		func(i any) bool { return i == "" }))
	assert.Equal(t, -1, LastIndexOfBy([]any{1, 2, 3},
		func(i any) bool { return i == 4 }))
	assert.Equal(t, -1, LastIndexOfBy([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 3.35 }))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, -1, LastIndexOfBy([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "4"} }))

	assert.Equal(t, 0, LastIndexOfBy([]any{1},
		func(i any) bool { return i == 1 }))
	assert.Equal(t, 4, LastIndexOfBy([]any{1, 2, 3, 1, 2, 3},
		func(i any) bool { return i == 2 }))
	assert.Equal(t, 1, LastIndexOfBy([]any{"one", "two"},
		func(i any) bool { return i == "two" }))
	assert.Equal(t, 3, LastIndexOfBy([]any{"one", "", "two", ""},
		func(i any) bool { return i == "" }))
	assert.Equal(t, 1, LastIndexOfBy([]any{1.1, 1.1, 2.2, 3.3},
		func(i any) bool { return i == 1.1 }))
	assert.Equal(t, 2, LastIndexOfBy([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "3"} }))
}

func Test_LastIndexOfPred_Deprecated(t *testing.T) {
	assert.Equal(t, -1, LastIndexOfPred([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 3.35 }))
	assert.Equal(t, 1, LastIndexOfPred([]any{1.1, 1.1, 2.2, 3.3},
		func(i any) bool { return i == 1.1 }))
}

func Test_RemoveAt(t *testing.T) {
	s1 := []int{1}
	RemoveAt(&s1, 0)
	assert.Equal(t, []int{}, s1)

	s1 = []int{1, 2, 3}
	RemoveAt(&s1, 1)
	assert.Equal(t, []int{1, 3}, s1)

	s1 = []int{1, 2, 3}
	RemoveAt(&s1, 2)
	assert.Equal(t, []int{1, 2}, s1)

	s2 := []string{"", "one", "two"}
	RemoveAt(&s2, 0)
	assert.Equal(t, []string{"one", "two"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}}
	RemoveAt(&s3, 1)
	assert.Equal(t, []St{{1, "1"}, {3, "3"}}, s3)

	// Test IndexOutOfRange error
	defer func() {
		e := recover()
		assert.ErrorIs(t, e.(error), ErrIndexOutOfRange) // nolint: forcetypeassert
	}()
	s1 = []int{1, 2, 3}
	RemoveAt(&s1, 3)
}

func Test_FastRemoveAt(t *testing.T) {
	s1 := []int{1}
	FastRemoveAt(&s1, 0)
	assert.Equal(t, []int{}, s1)

	s1 = []int{1, 2, 3}
	FastRemoveAt(&s1, 0)
	assert.Equal(t, []int{3, 2}, s1)

	s1 = []int{1, 2, 3}
	FastRemoveAt(&s1, 2)
	assert.Equal(t, []int{1, 2}, s1)

	s2 := []string{"", "one", "two"}
	FastRemoveAt(&s2, 0)
	assert.Equal(t, []string{"two", "one"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}}
	FastRemoveAt(&s3, 1)
	assert.Equal(t, []St{{1, "1"}, {3, "3"}}, s3)

	// Test IndexOutOfRange error
	defer func() {
		e := recover()
		assert.ErrorIs(t, e.(error), ErrIndexOutOfRange) // nolint: forcetypeassert
	}()
	s1 = []int{1, 2, 3}
	FastRemoveAt(&s1, 3)
}

func Test_Remove(t *testing.T) {
	s1 := []int{1}
	Remove(&s1, 2)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1}
	Remove(&s1, 1)
	assert.Equal(t, []int{}, s1)

	s1 = []int{1, 2, 3}
	Remove(&s1, 1)
	assert.Equal(t, []int{2, 3}, s1)

	s1 = []int{1, 2, 3}
	Remove(&s1, 2)
	assert.Equal(t, []int{1, 3}, s1)

	s2 := []string{"", "one", "two"}
	Remove(&s2, "two")
	assert.Equal(t, []string{"", "one"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}}
	Remove(&s3, St{2, "2"})
	assert.Equal(t, []St{{1, "1"}, {3, "3"}}, s3)
}

func Test_FastRemove(t *testing.T) {
	s1 := []int{1}
	FastRemove(&s1, 2)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1}
	FastRemove(&s1, 1)
	assert.Equal(t, []int{}, s1)

	s1 = []int{1, 2, 3}
	FastRemove(&s1, 1)
	assert.Equal(t, []int{3, 2}, s1)

	s1 = []int{1, 2, 3}
	FastRemove(&s1, 3)
	assert.Equal(t, []int{1, 2}, s1)

	s2 := []string{"", "one", "two"}
	FastRemove(&s2, "")
	assert.Equal(t, []string{"two", "one"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}}
	FastRemove(&s3, St{2, "2"})
	assert.Equal(t, []St{{1, "1"}, {3, "3"}}, s3)
}

func Test_RemoveLastOf(t *testing.T) {
	s1 := []int{1}
	RemoveLastOf(&s1, 2)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1, 1}
	RemoveLastOf(&s1, 1)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1, 2, 1}
	RemoveLastOf(&s1, 1)
	assert.Equal(t, []int{1, 2}, s1)

	s1 = []int{1, 2, 3, 2, 2}
	RemoveLastOf(&s1, 2)
	assert.Equal(t, []int{1, 2, 3, 2}, s1)

	s2 := []string{"", "one", "", "two"}
	RemoveLastOf(&s2, "")
	assert.Equal(t, []string{"", "one", "two"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}, {2, "2"}, {3, "3"}}
	RemoveLastOf(&s3, St{2, "2"})
	assert.Equal(t, []St{{1, "1"}, {2, "2"}, {3, "3"}, {3, "3"}}, s3)
}

func Test_FastRemoveLastOf(t *testing.T) {
	s1 := []int{1}
	FastRemoveLastOf(&s1, 2)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1, 1}
	FastRemoveLastOf(&s1, 1)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1, 2, 1, 2, 3}
	FastRemoveLastOf(&s1, 1)
	assert.Equal(t, []int{1, 2, 3, 2}, s1)

	s1 = []int{1, 2, 3, 2, 2}
	FastRemoveLastOf(&s1, 2)
	assert.Equal(t, []int{1, 2, 3, 2}, s1)

	s2 := []string{"", "one", "", "two", "three"}
	FastRemoveLastOf(&s2, "")
	assert.Equal(t, []string{"", "one", "three", "two"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}, {2, "2"}, {3, "3"}}
	FastRemoveLastOf(&s3, St{2, "2"})
	assert.Equal(t, []St{{1, "1"}, {2, "2"}, {3, "3"}, {3, "3"}}, s3)
}

func Test_RemoveAll(t *testing.T) {
	s1 := []int{1}
	assert.Equal(t, 0, RemoveAll(&s1, 2))
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1, 1}
	assert.Equal(t, 2, RemoveAll(&s1, 1))
	assert.Equal(t, []int{}, s1)

	s1 = []int{1, 2, 1, 2, 3, 2}
	assert.Equal(t, 3, RemoveAll(&s1, 2))
	assert.Equal(t, []int{1, 1, 3}, s1)

	s1 = []int{1, 2, 3, 1, 2, 3, 1}
	assert.Equal(t, 3, RemoveAll(&s1, 1))
	assert.Equal(t, []int{2, 3, 2, 3}, s1)

	s2 := []string{"", "one", "", "two", "three", ""}
	assert.Equal(t, 3, RemoveAll(&s2, ""))
	assert.Equal(t, []string{"one", "two", "three"}, s2)
}

func Test_Reverse(t *testing.T) {
	assert.Equal(t, []int{}, Reverse([]int{}))
	assert.Equal(t, []int64{1}, Reverse([]int64{1}))
	assert.Equal(t, []int{3, 2, 1}, Reverse([]int{1, 2, 3}))

	s := []int{-1, -2, 0, 1, 2}
	Reverse(s)
	assert.Equal(t, []int{2, 1, 0, -2, -1}, s)
}

func Test_ReverseCopy(t *testing.T) {
	assert.Equal(t, []int{}, ReverseCopy([]int{}))
	assert.Equal(t, []int64{1}, ReverseCopy([]int64{1}))
	assert.Equal(t, []int{3, 2, 1}, ReverseCopy([]int{1, 2, 3}))

	s := []int{-1, -2, 0, 1, 2}
	s2 := ReverseCopy(s)
	assert.Equal(t, []int{-1, -2, 0, 1, 2}, s)
	assert.Equal(t, []int{2, 1, 0, -2, -1}, s2)
}

func Test_Compact(t *testing.T) {
	assert.Equal(t, []int{1, -1}, Compact([]int{1, 0, -1}))
	assert.Equal(t, []bool{true, true, true}, Compact([]bool{true, true, false, false, true}))
	assert.Equal(t, []string{"1", "2"}, Compact([]string{"1", "", "2"}))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, []St{{1, "1"}, {2, "2"}}, Compact([]St{{1, "1"}, {}, {2, "2"}}))
}

func Test_Replace(t *testing.T) {
	// No replacement done
	s := []int{}
	assert.True(t, Replace(s, 1, 11) == false && reflect.DeepEqual(s, []int{}))
	s = []int{0, 1, 2, 3, 4, 5}
	assert.True(t, Replace(s, 100, 1000) == false && reflect.DeepEqual(s, []int{0, 1, 2, 3, 4, 5}))
	s2 := []string{"one", "two"}
	assert.True(t, Replace(s2, "One", "ONE") == false && reflect.DeepEqual(s2, []string{"one", "two"}))

	// Has replacement
	s = []int{0, 1, 2, 3, 0, 1, 2, 3}
	assert.True(t, Replace(s, 1, 11) == true && reflect.DeepEqual(s, []int{0, 11, 2, 3, 0, 1, 2, 3}))
	s2 = []string{"one", "two", "one"}
	assert.True(t, Replace(s2, "one", "1") == true && reflect.DeepEqual(s2, []string{"1", "two", "one"}))
}

func Test_ReplaceN(t *testing.T) {
	// No replacement done
	s := []int{}
	assert.True(t, ReplaceN(s, 1, 11, 10) == 0 && reflect.DeepEqual(s, []int{}))
	s = []int{0, 1, 2, 3, 4, 5}
	assert.True(t, ReplaceN(s, 100, 1000, 10) == 0 && reflect.DeepEqual(s, []int{0, 1, 2, 3, 4, 5}))
	assert.True(t, ReplaceN(s, 100, 1000, 0) == 0 && reflect.DeepEqual(s, []int{0, 1, 2, 3, 4, 5}))
	assert.True(t, ReplaceN(s, 100, 1000, -2) == 0 && reflect.DeepEqual(s, []int{0, 1, 2, 3, 4, 5}))
	s2 := []string{"one", "two"}
	assert.True(t, ReplaceN(s2, "One", "ONE", 10) == 0 && reflect.DeepEqual(s2, []string{"one", "two"}))

	// Has replacement
	s = []int{0, 1, 2, 3, 0, 1, 2, 3}
	assert.True(t, ReplaceN(s, 1, 11, 10) == 2 && reflect.DeepEqual(s, []int{0, 11, 2, 3, 0, 11, 2, 3}))
	s = []int{0, 1, 2, 3, 0, 1, 2, 3}
	assert.True(t, ReplaceN(s, 1, 11, 1) == 1 && reflect.DeepEqual(s, []int{0, 11, 2, 3, 0, 1, 2, 3}))
	s2 = []string{"one", "two", "one"}
	assert.True(t, ReplaceN(s2, "one", "1", -1) == 2 && reflect.DeepEqual(s2, []string{"1", "two", "1"}))
	s2 = []string{"one", "two", "one"}
	assert.True(t, ReplaceN(s2, "one", "1", 3) == 2 && reflect.DeepEqual(s2, []string{"1", "two", "1"}))
}

func Test_ReplaceAll(t *testing.T) {
	// No replacement done
	s := []int{}
	assert.True(t, ReplaceAll(s, 1, 11) == 0 && reflect.DeepEqual(s, []int{}))
	s = []int{0, 1, 2, 3, 4, 5}
	assert.True(t, ReplaceAll(s, 100, 1000) == 0 && reflect.DeepEqual(s, []int{0, 1, 2, 3, 4, 5}))
	s2 := []string{"one", "two"}
	assert.True(t, ReplaceAll(s2, "One", "ONE") == 0 && reflect.DeepEqual(s2, []string{"one", "two"}))

	// Has replacement
	s = []int{0, 1, 2, 3, 0, 1, 2, 3}
	assert.True(t, ReplaceAll(s, 1, 11) == 2 && reflect.DeepEqual(s, []int{0, 11, 2, 3, 0, 11, 2, 3}))
	s2 = []string{"one", "two", "one"}
	assert.True(t, ReplaceAll(s2, "one", "1") == 2 && reflect.DeepEqual(s2, []string{"1", "two", "1"}))
}

func Test_Fill(t *testing.T) {
	s := make([]int, 5)
	Fill(s, 1)
	assert.Equal(t, []int{1, 1, 1, 1, 1}, s)
	Fill(s, 2)
	assert.Equal(t, []int{2, 2, 2, 2, 2}, s)
	sub := s[2:]
	Fill(sub, 5)
	assert.Equal(t, []int{5, 5, 5}, sub)
	assert.Equal(t, []int{2, 2, 5, 5, 5}, s)

	s2 := []int{}
	Fill(s2, 1)
	assert.Equal(t, []int{}, s2)
}

func Test_Chunk(t *testing.T) {
	// Empty input
	chunks := Chunk([]int{}, 5)
	assert.True(t, len(chunks) == 0)

	// Nil input
	chunks = Chunk[int]([]int(nil), 5)
	assert.True(t, len(chunks) == 0)

	// Chunk size greater than input size
	chunks = Chunk([]int{1, 2, 3}, 5)
	assert.True(t, len(chunks) == 1 && reflect.DeepEqual(chunks[0], []int{1, 2, 3}))

	// Normal case
	chunks = Chunk([]int{1, 2, 3, 4, 5}, 2)
	assert.True(t, len(chunks) == 3 &&
		len(chunks[0]) == 2 && reflect.DeepEqual(chunks[0], []int{1, 2}) &&
		len(chunks[1]) == 2 && reflect.DeepEqual(chunks[1], []int{3, 4}) &&
		len(chunks[2]) == 1 && reflect.DeepEqual(chunks[2], []int{5}))
}

func Test_ChunkByPieces(t *testing.T) {
	// Empty input
	chunks := ChunkByPieces([]int{}, 5)
	assert.True(t, len(chunks) == 0)

	// Nil input
	chunks = ChunkByPieces[int]([]int(nil), 5)
	assert.True(t, len(chunks) == 0)

	// Chunk count is zero
	chunks = ChunkByPieces([]int{1, 2, 3}, 0)
	assert.Equal(t, [][]int{}, chunks)

	// Chunk count greater than input size
	chunks = ChunkByPieces([]int{1, 2, 3}, 5)
	assert.True(t, len(chunks) == 3 &&
		len(chunks[0]) == 1 && reflect.DeepEqual(chunks[0], []int{1}) &&
		len(chunks[1]) == 1 && reflect.DeepEqual(chunks[1], []int{2}) &&
		len(chunks[2]) == 1 && reflect.DeepEqual(chunks[2], []int{3}))

	// Normal case
	chunks = ChunkByPieces([]int{1, 2, 3, 4, 5}, 2)
	assert.True(t, len(chunks) == 2 &&
		len(chunks[0]) == 3 && reflect.DeepEqual(chunks[0], []int{1, 2, 3}) &&
		len(chunks[1]) == 2 && reflect.DeepEqual(chunks[1], []int{4, 5}))
}

func Test_CountValue(t *testing.T) {
	assert.Equal(t, 0, CountValue([]int{1, 2, 3}, 4))
	assert.Equal(t, 1, CountValue([]int{1, 2, 3}, 2))
	assert.Equal(t, 2, CountValue([]int{1, 2, 3, 2}, 2))
	assert.Equal(t, 2, CountValue([]float32{1.1, 2.2, 3.3, 1.100001, 1.1}, 1.1))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, 2, CountValue([]St{{1, "1"}, {2, "2"}, {1, "1"}}, St{1, "1"}))
	assert.Equal(t, 0, CountValue([]St{{1, "1"}, {2, "2"}, {1, "1"}}, St{1, "2"}))
}

func Test_CountValueBy(t *testing.T) {
	assert.Equal(t, 0, CountValueBy([]any{1, 2, 3},
		func(t any) bool { return t == 4 }))
	assert.Equal(t, 1, CountValueBy([]any{1, 2, 3},
		func(t any) bool { return t == 2 }))
	assert.Equal(t, 2, CountValueBy([]any{1, 2, 3, 2},
		func(t any) bool { return t == 2 }))
	assert.Equal(t, 2, CountValueBy([]any{1.1, 2.2, 3.3, 1.100001, 1.1},
		func(t any) bool { return t == 1.1 }))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, 2, CountValueBy([]any{St{1, "1"}, St{2, "2"}, St{1, "1"}},
		func(t any) bool { return t == St{1, "1"} }))
	assert.Equal(t, 0, CountValueBy([]any{St{1, "1"}, St{2, "2"}, St{1, "1"}},
		func(t any) bool { return t == St{1, "2"} }))
}

func Test_CountValuePred_Deprecated(t *testing.T) {
	assert.Equal(t, 2, CountValuePred([]any{1, 2, 3, 2},
		func(t any) bool { return t == 2 }))
	assert.Equal(t, 2, CountValuePred([]any{1.1, 2.2, 3.3, 1.100001, 1.1},
		func(t any) bool { return t == 1.1 }))
}

func Test_Drop(t *testing.T) {
	// Nil/empty source slice
	assert.Equal(t, []int{}, Drop([]int(nil)))
	assert.Equal(t, []int{}, Drop([]int{}))

	assert.Equal(t, []int{1, 4}, Drop([]int{1, 2, 3, 4, 5}, 5, 3, 2, 7))
}

func Test_GetFirst(t *testing.T) {
	assert.Equal(t, 1, GetFirst([]int{1, 2, 3}, 4))
	assert.Equal(t, 11, GetFirst([]int{}, 11))
}

func Test_GetLast(t *testing.T) {
	assert.Equal(t, 3, GetLast([]int{1, 2, 3}, 4))
	assert.Equal(t, 11, GetLast([]int{}, 11))
}

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

func Test_SliceByRange(t *testing.T) {
	// start < end
	assert.Equal(t, []int{}, SliceByRange(0, 5, 0))
	assert.Equal(t, []int{0, 1, 2, 3, 4}, SliceByRange(0, 5, 1))
	assert.Equal(t, []int{0, 2, 4}, SliceByRange(0, 5, 2))
	assert.Equal(t, []float64{-5, -4, -3, -2, -1}, SliceByRange(float64(-5.0), 0, 1))
	assert.Equal(t, []int32{-5, -3, -1}, SliceByRange(int32(-5), 0, 2))

	// start > end
	assert.Equal(t, []int{}, SliceByRange(5, 0, 0))
	assert.Equal(t, []int{5, 4, 3, 2, 1}, SliceByRange(5, 0, -1))
	assert.Equal(t, []int{5, 3, 1}, SliceByRange(5, 0, -2))
	assert.Equal(t, []float64{0, -1, -2, -3, -4}, SliceByRange(float64(0.0), -5, -1))
	assert.Equal(t, []int32{0, -2, -4}, SliceByRange(int32(0), -5, -2))
}
