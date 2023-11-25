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
}

// nolint: forcetypeassert
func Test_EqualPred(t *testing.T) {
	assert.True(t, EqualPred([]any{}, []any{},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualPred([]any{}, nil,
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualPred(nil, []any{},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualPred([]any{1, 2, 3}, []any{1, 2, 3},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualPred([]any{"3", "1", "2"}, []any{"3", "1", "2"},
		func(a, b any) bool { return a.(string) == b.(string) }))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, EqualPred(
		[]any{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(a, b any) bool { return a.(St) == b.(St) }))

	assert.False(t, EqualPred([]any{}, []any{1},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualPred([]any{1}, nil,
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualPred([]any{1, 2, 3}, []any{1, 2, 3, 4},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualPred([]any{1, 2, 3}, []any{3, 2, 1},
		func(a, b any) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualPred([]any{"3", "1", "2"}, []any{"1", "2", "3"},
		func(a, b any) bool { return a.(string) == b.(string) }))
	assert.False(t, EqualPred(
		[]any{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []any{St{1, "1"}, St{2, "2"}},
		func(a, b any) bool { return a.(St) == b.(St) }))
	assert.False(t, EqualPred(
		[]any{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []any{St{3, "3"}, St{2, "2"}, St{1, "1"}},
		func(a, b any) bool { return a.(St) == b.(St) }))
}

func Test_EqualPredPtr(t *testing.T) {
	assert.True(t, EqualPredPtr([]int{}, []int{},
		func(a, b *int) bool { return *a == *b }))
	assert.True(t, EqualPredPtr([]int{}, nil,
		func(a, b *int) bool { return *a == *b }))
	assert.True(t, EqualPredPtr(nil, []int{},
		func(a, b *int) bool { return *a == *b }))
	assert.True(t, EqualPredPtr([]int{1, 2, 3}, []int{1, 2, 3},
		func(a, b *int) bool { return *a == *b }))
	assert.True(t, EqualPredPtr([]string{"3", "1", "2"}, []string{"3", "1", "2"},
		func(a, b *string) bool { return *a == *b }))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, EqualPredPtr(
		[]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(a, b *St) bool { return *a == *b }))

	assert.False(t, EqualPredPtr([]int{}, []int{1},
		func(a, b *int) bool { return *a == *b }))
	assert.False(t, EqualPredPtr([]int{1}, nil,
		func(a, b *int) bool { return *a == *b }))
	assert.False(t, EqualPredPtr([]int{1, 2, 3}, []int{1, 2, 3, 4},
		func(a, b *int) bool { return *a == *b }))
	assert.False(t, EqualPredPtr([]int{1, 2, 3}, []int{3, 2, 1},
		func(a, b *int) bool { return *a == *b }))
	assert.False(t, EqualPredPtr([]string{"3", "1", "2"}, []string{"1", "2", "3"},
		func(a, b *string) bool { return *a == *b }))
	assert.False(t, EqualPredPtr(
		[]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{1, "1"}, {2, "2"}},
		func(a, b *St) bool { return *a == *b }))
	assert.False(t, EqualPredPtr(
		[]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{3, "3"}, {2, "2"}, {1, "1"}},
		func(a, b *St) bool { return *a == *b }))
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
func Test_ContentEqualPred(t *testing.T) {
	assert.True(t, ContentEqualPred([]any{}, []any{},
		func(t any) int { return t.(int) }))
	assert.True(t, ContentEqualPred([]any{}, nil,
		func(t any) int { return t.(int) }))
	assert.True(t, ContentEqualPred(nil, []any{},
		func(t any) int { return t.(int) }))
	assert.True(t, ContentEqualPred([]any{1, 2, 3}, []any{1, 2, 3},
		func(t any) int { return t.(int) }))
	assert.True(t, ContentEqualPred([]any{3, 1, 2}, []any{1, 2, 3},
		func(t any) int { return t.(int) }))
	assert.True(t, ContentEqualPred([]any{"3", "1", "2"}, []any{"1", "2", "3"},
		func(t any) string { return t.(string) }))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, ContentEqualPred(
		[]any{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []any{St{3, "3"}, St{1, "1"}, St{2, "2"}},
		func(t any) St { return t.(St) }))

	assert.False(t, ContentEqualPred([]any{}, []any{1},
		func(t any) int { return t.(int) }))
	assert.False(t, ContentEqualPred([]any{1}, nil,
		func(t any) int { return t.(int) }))
	assert.False(t, ContentEqualPred([]any{1, 2, 3}, []any{1, 2, 3, 4},
		func(t any) int { return t.(int) }))
	assert.False(t, ContentEqualPred(
		[]any{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []any{St{3, "3"}, St{1, "1"}, St{3, "3"}},
		func(t any) St { return t.(St) }))
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
	assert.Equal(t, []int{}, Concat[int](nil, nil, nil))
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
func Test_ContainPred(t *testing.T) {
	assert.False(t, ContainPred([]any{},
		func(i any) bool { return i.(int) == 1 }))
	assert.False(t, ContainPred([]any{"one"},
		func(i any) bool { return i == "One" }))
	assert.False(t, ContainPred([]any{"one", "two"},
		func(i any) bool { return i == "" }))
	assert.False(t, ContainPred([]any{1, 2, 3},
		func(i any) bool { return i == 4 }))
	assert.False(t, ContainPred([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 3.35 }))

	type St struct {
		Int int
		Str string
	}
	assert.False(t, ContainPred([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "4"} }))

	assert.True(t, ContainPred([]any{1},
		func(i any) bool { return i == 1 }))
	assert.True(t, ContainPred([]any{1, 2, 3, 1, 2, 3},
		func(i any) bool { return i == 2 }))
	assert.True(t, ContainPred([]any{"one", "two"},
		func(i any) bool { return i == "two" }))
	assert.True(t, ContainPred([]any{"one", "two", ""},
		func(i any) bool { return i == "" }))
	assert.True(t, ContainPred([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 2.2 }))
	assert.True(t, ContainPred([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "3"} }))
}

func Test_ContainPredPtr(t *testing.T) {
	assert.False(t, ContainPredPtr([]int{},
		func(i *int) bool { return *i == 1 }))
	assert.False(t, ContainPredPtr([]string{"one"},
		func(i *string) bool { return *i == "One" }))
	assert.False(t, ContainPredPtr([]string{"one", "two"},
		func(i *string) bool { return *i == "" }))
	assert.False(t, ContainPredPtr([]int{1, 2, 3},
		func(i *int) bool { return *i == 4 }))
	assert.False(t, ContainPredPtr([]float32{1.1, 2.2, 3.3},
		func(i *float32) bool { return *i == 3.35 }))

	type St struct {
		Int int
		Str string
	}
	assert.False(t, ContainPredPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return *i == St{3, "4"} }))

	assert.True(t, ContainPredPtr([]int{1},
		func(i *int) bool { return *i == 1 }))
	assert.True(t, ContainPredPtr([]int{1, 2, 3, 1, 2, 3},
		func(i *int) bool { return *i == 2 }))
	assert.True(t, ContainPredPtr([]string{"one", "two"},
		func(i *string) bool { return *i == "two" }))
	assert.True(t, ContainPredPtr([]string{"one", "two", ""},
		func(i *string) bool { return *i == "" }))
	assert.True(t, ContainPredPtr([]float32{1.1, 2.2, 3.3},
		func(i *float32) bool { return *i == 2.2 }))
	assert.True(t, ContainPredPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return *i == St{3, "3"} }))
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
	assert.True(t, IsUnique[int](nil))
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
func Test_IsUniquePred(t *testing.T) {
	assert.True(t, IsUniquePred[int, int](nil, nil))
	assert.True(t, IsUniquePred([]int{}, func(v int) int { return v }))
	assert.True(t, IsUniquePred([]any{"one"},
		func(v any) string { return v.(string) }))
	assert.True(t, IsUniquePred([]any{"one", "two", "One", "Two"},
		func(v any) string { return v.(string) }))
	assert.True(t, IsUniquePred([]any{1.1, 2.2, 3.3, 1.11},
		func(v any) float64 { return v.(float64) }))

	assert.False(t, IsUniquePred([]any{1, 2, 3, 1, 2},
		func(v any) int { return v.(int) }))
	assert.False(t, IsUniquePred([]any{"one", "two", "one"},
		func(v any) string { return v.(string) }))
	assert.False(t, IsUniquePred([]any{1.1, 2.2, 1.100},
		func(v any) float64 { return v.(float64) }))

	type st struct {
		I int
		S string
	}
	assert.True(t, IsUniquePred([]st{{1, "one"}, {2, "two"}, {3, "three"}}, func(v st) int { return v.I }))
	assert.False(t, IsUniquePred([]st{{1, "one"}, {1, "One"}}, func(v st) int { return v.I }))
	assert.False(t, IsUniquePred([]st{{1, "one"}, {2, "two"}, {1, "one"}}, func(v st) int { return v.I }))
}

func Test_FindPred(t *testing.T) {
	_, found := FindPred([]any{}, func(i any) bool { return i == 1 })
	assert.False(t, found)
	_, found = FindPred([]any{"one"}, func(i any) bool { return i == "One" })
	assert.False(t, found)
	_, found = FindPred([]any{"one", "two"}, func(i any) bool { return i == "" })
	assert.False(t, found)
	_, found = FindPred([]any{1, 2, 3}, func(i any) bool { return i == 4 })
	assert.False(t, found)
	_, found = FindPred([]any{1.1, 2.2, 3.3}, func(i any) bool { return i == 3.35 })
	assert.False(t, found)

	type St struct {
		Int int
		Str string
	}
	_, found = FindPred([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i St) bool { return i.Int == 4 })
	assert.False(t, found)

	v1, found := FindPred([]any{1}, func(i any) bool { return i == 1 })
	assert.True(t, found)
	assert.Equal(t, 1, v1)
	v2, found := FindPred([]any{1, 2, 3, 1, 2, 3}, func(i any) bool { return i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
	v3, found := FindPred([]any{"one", "two"}, func(i any) bool { return i == "two" })
	assert.True(t, found)
	assert.Equal(t, "two", v3)
	v4, found := FindPred([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "3"} })
	assert.True(t, found)
	assert.Equal(t, St{3, "3"}, v4)
	v5, found := FindPred([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i St) bool { return i.Int == 2 })
	assert.True(t, found)
	assert.Equal(t, St{2, "2"}, v5)
}

func Test_FindPredPtr(t *testing.T) {
	_, found := FindPredPtr([]any{}, func(i *any) bool { return *i == 1 })
	assert.False(t, found)
	_, found = FindPredPtr([]any{"one"}, func(i *any) bool { return *i == "One" })
	assert.False(t, found)
	_, found = FindPredPtr([]any{"one", "two"}, func(i *any) bool { return *i == "" })
	assert.False(t, found)
	_, found = FindPredPtr([]any{1, 2, 3}, func(i *any) bool { return *i == 4 })
	assert.False(t, found)
	_, found = FindPredPtr([]any{1.1, 2.2, 3.3}, func(i *any) bool { return *i == 3.35 })
	assert.False(t, found)

	type St struct {
		Int int
		Str string
	}
	_, found = FindPredPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return i.Int == 4 })
	assert.False(t, found)

	v1, found := FindPredPtr([]any{1}, func(i *any) bool { return *i == 1 })
	assert.True(t, found)
	assert.Equal(t, 1, v1)
	v2, found := FindPredPtr([]any{1, 2, 3, 1, 2, 3}, func(i *any) bool { return *i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
	v3, found := FindPredPtr([]any{"one", "two"}, func(i *any) bool { return *i == "two" })
	assert.True(t, found)
	assert.Equal(t, "two", v3)
	v4, found := FindPredPtr([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i *any) bool { return *i == St{3, "3"} })
	assert.True(t, found)
	assert.Equal(t, St{3, "3"}, v4)
	v5, found := FindPredPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return i.Int == 2 })
	assert.True(t, found)
	assert.Equal(t, St{2, "2"}, v5)
}

func Test_FindLastPred(t *testing.T) {
	_, found := FindLastPred([]any{}, func(i any) bool { return i == 1 })
	assert.False(t, found)
	_, found = FindLastPred([]any{"one"}, func(i any) bool { return i == "One" })
	assert.False(t, found)
	_, found = FindLastPred([]any{"one", "two"}, func(i any) bool { return i == "" })
	assert.False(t, found)
	_, found = FindLastPred([]any{1, 2, 3}, func(i any) bool { return i == 4 })
	assert.False(t, found)
	_, found = FindLastPred([]any{1.1, 2.2, 3.3}, func(i any) bool { return i == 3.35 })
	assert.False(t, found)

	type St struct {
		Int int
		Str string
	}
	_, found = FindLastPred([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i St) bool { return i.Int == 4 })
	assert.False(t, found)

	v1, found := FindLastPred([]any{1}, func(i any) bool { return i == 1 })
	assert.True(t, found)
	assert.Equal(t, 1, v1)
	v2, found := FindLastPred([]any{1, 2, 3, 1, 2, 3}, func(i any) bool { return i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
	v3, found := FindLastPred([]any{"one", "two"}, func(i any) bool { return i == "two" })
	assert.True(t, found)
	assert.Equal(t, "two", v3)
	v4, found := FindLastPred([]any{"one", "", "two", ""}, func(i any) bool { return i == "" })
	assert.True(t, found)
	assert.Equal(t, "", v4)
	v5, found := FindLastPred([]any{1.1, 1.1, 2.2, 3.3}, func(i any) bool { return i == 1.1 })
	assert.True(t, found)
	assert.Equal(t, 1.1, v5)
	v6, found := FindLastPred([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i St) bool { return i.Int == 3 })
	assert.True(t, found)
	assert.Equal(t, St{3, "3"}, v6)
}

func Test_FindLastPredPtr(t *testing.T) {
	_, found := FindLastPredPtr([]any{}, func(i *any) bool { return *i == 1 })
	assert.False(t, found)
	_, found = FindLastPredPtr([]any{"one"}, func(i *any) bool { return *i == "One" })
	assert.False(t, found)
	_, found = FindLastPredPtr([]any{"one", "two"}, func(i *any) bool { return *i == "" })
	assert.False(t, found)
	_, found = FindLastPredPtr([]any{1, 2, 3}, func(i *any) bool { return *i == 4 })
	assert.False(t, found)
	_, found = FindLastPredPtr([]any{1.1, 2.2, 3.3}, func(i *any) bool { return *i == 3.35 })
	assert.False(t, found)

	type St struct {
		Int int
		Str string
	}
	_, found = FindLastPredPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return i.Int == 4 })
	assert.False(t, found)

	v1, found := FindLastPredPtr([]any{1}, func(i *any) bool { return *i == 1 })
	assert.True(t, found)
	assert.Equal(t, 1, v1)
	v2, found := FindLastPredPtr([]any{1, 2, 3, 1, 2, 3}, func(i *any) bool { return *i == 2 })
	assert.True(t, found)
	assert.Equal(t, 2, v2)
	v3, found := FindLastPredPtr([]any{"one", "two"}, func(i *any) bool { return *i == "two" })
	assert.True(t, found)
	assert.Equal(t, "two", v3)
	v4, found := FindLastPredPtr([]any{"one", "", "two", ""}, func(i *any) bool { return *i == "" })
	assert.True(t, found)
	assert.Equal(t, "", v4)
	v5, found := FindLastPredPtr([]any{1.1, 1.1, 2.2, 3.3}, func(i *any) bool { return *i == 1.1 })
	assert.True(t, found)
	assert.Equal(t, 1.1, v5)
	v6, found := FindLastPredPtr([]St{{1, "1"}, {2, "2"}, {3, "3"}},
		func(i *St) bool { return i.Int == 3 })
	assert.True(t, found)
	assert.Equal(t, St{3, "3"}, v6)
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

func Test_IndexOfPred(t *testing.T) {
	assert.Equal(t, -1, IndexOfPred([]any{},
		func(i any) bool { return i == 1 }))
	assert.Equal(t, -1, IndexOfPred([]any{"one"},
		func(i any) bool { return i == "One" }))
	assert.Equal(t, -1, IndexOfPred([]any{"one", "two"},
		func(i any) bool { return i == "" }))
	assert.Equal(t, -1, IndexOfPred([]any{1, 2, 3},
		func(i any) bool { return i == 4 }))
	assert.Equal(t, -1, IndexOfPred([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 3.35 }))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, -1, IndexOfPred([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "4"} }))

	assert.Equal(t, 0, IndexOfPred([]any{1},
		func(i any) bool { return i == 1 }))
	assert.Equal(t, 1, IndexOfPred([]any{1, 2, 3, 1, 2, 3},
		func(i any) bool { return i == 2 }))
	assert.Equal(t, 1, IndexOfPred([]any{"one", "two"},
		func(i any) bool { return i == "two" }))
	assert.Equal(t, 2, IndexOfPred([]any{"one", "two", ""},
		func(i any) bool { return i == "" }))
	assert.Equal(t, 2, IndexOfPred([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 3.3 }))
	assert.Equal(t, 2, IndexOfPred([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "3"} }))
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

func Test_LastIndexOfPred(t *testing.T) {
	assert.Equal(t, -1, LastIndexOfPred([]any{},
		func(i any) bool { return i == 1 }))
	assert.Equal(t, -1, LastIndexOfPred([]any{"one"},
		func(i any) bool { return i == "One" }))
	assert.Equal(t, -1, LastIndexOfPred([]any{"one", "two"},
		func(i any) bool { return i == "" }))
	assert.Equal(t, -1, LastIndexOfPred([]any{1, 2, 3},
		func(i any) bool { return i == 4 }))
	assert.Equal(t, -1, LastIndexOfPred([]any{1.1, 2.2, 3.3},
		func(i any) bool { return i == 3.35 }))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, -1, LastIndexOfPred([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "4"} }))

	assert.Equal(t, 0, LastIndexOfPred([]any{1},
		func(i any) bool { return i == 1 }))
	assert.Equal(t, 4, LastIndexOfPred([]any{1, 2, 3, 1, 2, 3},
		func(i any) bool { return i == 2 }))
	assert.Equal(t, 1, LastIndexOfPred([]any{"one", "two"},
		func(i any) bool { return i == "two" }))
	assert.Equal(t, 3, LastIndexOfPred([]any{"one", "", "two", ""},
		func(i any) bool { return i == "" }))
	assert.Equal(t, 1, LastIndexOfPred([]any{1.1, 1.1, 2.2, 3.3},
		func(i any) bool { return i == 1.1 }))
	assert.Equal(t, 2, LastIndexOfPred([]any{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i any) bool { return i == St{3, "3"} }))
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

func Test_CountValuePred(t *testing.T) {
	assert.Equal(t, 0, CountValuePred([]any{1, 2, 3},
		func(t any) bool { return t == 4 }))
	assert.Equal(t, 1, CountValuePred([]any{1, 2, 3},
		func(t any) bool { return t == 2 }))
	assert.Equal(t, 2, CountValuePred([]any{1, 2, 3, 2},
		func(t any) bool { return t == 2 }))
	assert.Equal(t, 2, CountValuePred([]any{1.1, 2.2, 3.3, 1.100001, 1.1},
		func(t any) bool { return t == 1.1 }))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, 2, CountValuePred([]any{St{1, "1"}, St{2, "2"}, St{1, "1"}},
		func(t any) bool { return t == St{1, "1"} }))
	assert.Equal(t, 0, CountValuePred([]any{St{1, "1"}, St{2, "2"}, St{1, "1"}},
		func(t any) bool { return t == St{1, "2"} }))
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
