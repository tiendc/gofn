package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Union(t *testing.T) {
	assert.Equal(t, []int{}, Union[int]([]int(nil), nil))
	assert.Equal(t, []int{}, Union(nil, []int{}))
	assert.Equal(t, []int{1}, Union([]int{1}, nil))

	assert.Equal(t, []int{1, 2, 3, 4}, Union([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4}, Union([]int{1, 2, 3, 2}, []int{1, 2, 4, 3}))
	assert.Equal(t, []string{"1", "2", "3", "4"}, Union([]string{"1", "2", "3", "2"}, []string{"1", "2", "4", "3"}))
}

// nolint: forcetypeassert
func Test_UnionBy(t *testing.T) {
	assert.Equal(t, []any{}, UnionBy[any]([]any(nil), nil, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, UnionBy(nil, []any{}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1}, UnionBy([]any{1}, nil, func(t any) int { return t.(int) }))

	assert.Equal(t, []any{1, 2, 3, 4}, UnionBy([]any{1, 2}, []any{3, 4},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1, 2, 3, 4}, UnionBy([]any{1, 2, 3, 2}, []any{1, 2, 4, 3},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{"1", "2", "3", "4"},
		UnionBy([]any{"1", "2", "3", "2"}, []any{"1", "2", "4", "3"},
			func(t any) string { return t.(string) }))
}

// nolint: forcetypeassert
func Test_UnionPred_Deprecated(t *testing.T) {
	assert.Equal(t, []any{1, 2, 3, 4}, UnionPred([]any{1, 2}, []any{3, 4},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1, 2, 3, 4}, UnionPred([]any{1, 2, 3, 2}, []any{1, 2, 4, 3},
		func(t any) int { return t.(int) }))
}

func Test_Intersection(t *testing.T) {
	assert.Equal(t, []int{}, Intersection[int]([]int(nil), nil))
	assert.Equal(t, []int{}, Intersection(nil, []int{}))
	assert.Equal(t, []int{}, Intersection([]int{1}, nil))
	assert.Equal(t, []int{}, Intersection(nil, []int{1}))

	assert.Equal(t, []int{}, Intersection([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{1, 2, 3}, Intersection([]int{1, 2, 3, 2}, []int{1, 2, 4, 3}))
	assert.Equal(t, []string{"1", "2", "3"}, Intersection([]string{"1", "2", "3", "2"}, []string{"1", "2", "4", "3"}))
}

// nolint: forcetypeassert
func Test_IntersectionBy(t *testing.T) {
	assert.Equal(t, []any{}, IntersectionBy[any]([]any(nil), nil, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, IntersectionBy(nil, []any{}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, IntersectionBy([]any{1}, nil, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, IntersectionBy(nil, []any{1}, func(t any) int { return t.(int) }))

	assert.Equal(t, []any{}, IntersectionBy([]any{1, 2}, []any{3, 4},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1, 2, 3}, IntersectionBy([]any{1, 2, 3, 2}, []any{1, 2, 4, 3},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{"1", "2", "3"},
		IntersectionBy([]any{"1", "2", "3", "2"}, []any{"1", "2", "4", "3"},
			func(t any) string { return t.(string) }))
}

// nolint: forcetypeassert
func Test_IntersectionPred_Deprecated(t *testing.T) {
	assert.Equal(t, []any{}, IntersectionPred([]any{1, 2}, []any{3, 4},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1, 2, 3}, IntersectionPred([]any{1, 2, 3, 2}, []any{1, 2, 4, 3},
		func(t any) int { return t.(int) }))
}

func Test_Difference(t *testing.T) {
	l, r := Difference[int]([]int(nil), nil)
	assert.Equal(t, []int{}, l)
	assert.Equal(t, []int{}, r)
	l, r = Difference([]int{}, nil)
	assert.Equal(t, []int{}, l)
	assert.Equal(t, []int{}, r)

	l, r = Difference([]int{1, 2}, []int{3, 4})
	assert.Equal(t, []int{1, 2}, l)
	assert.Equal(t, []int{3, 4}, r)

	l, r = Difference([]int{1, 2, 3, 2}, []int{1, 4, 4, 3})
	assert.Equal(t, []int{2, 2}, l)
	assert.Equal(t, []int{4, 4}, r)

	l2, r2 := Difference([]string{"1", "2", "3", "2"}, []string{"1", "4", "2", "", "3"})
	assert.Equal(t, []string{}, l2)
	assert.Equal(t, []string{"4", ""}, r2)
}

// nolint: forcetypeassert
func Test_DifferenceBy(t *testing.T) {
	l, r := DifferenceBy[any]([]any(nil), nil, func(t any) int { return t.(int) })
	assert.Equal(t, []any{}, l)
	assert.Equal(t, []any{}, r)
	l, r = DifferenceBy([]any{}, nil, func(t any) int { return t.(int) })
	assert.Equal(t, []any{}, l)
	assert.Equal(t, []any{}, r)

	l, r = DifferenceBy([]any{1, 2}, []any{3, 4}, func(t any) int { return t.(int) })
	assert.Equal(t, []any{1, 2}, l)
	assert.Equal(t, []any{3, 4}, r)

	l, r = DifferenceBy([]any{1, 2, 3, 2}, []any{1, 4, 4, 3}, func(t any) int { return t.(int) })
	assert.Equal(t, []any{2, 2}, l)
	assert.Equal(t, []any{4, 4}, r)

	l, r = DifferenceBy([]any{"1", "2", "3", "2"}, []any{"1", "4", "2", "", "3"},
		func(t any) string { return t.(string) })
	assert.Equal(t, []any{}, l)
	assert.Equal(t, []any{"4", ""}, r)
}

// nolint: forcetypeassert
func Test_DifferencePred_Deprecated(t *testing.T) {
	l, r := DifferencePred([]any{1, 2}, []any{3, 4}, func(t any) int { return t.(int) })
	assert.Equal(t, []any{1, 2}, l)
	assert.Equal(t, []any{3, 4}, r)

	l, r = DifferencePred([]any{1, 2, 3, 2}, []any{1, 4, 4, 3}, func(t any) int { return t.(int) })
	assert.Equal(t, []any{2, 2}, l)
	assert.Equal(t, []any{4, 4}, r)
}
