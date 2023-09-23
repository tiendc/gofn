package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Union(t *testing.T) {
	assert.Equal(t, []int{}, Union[int](nil, nil))
	assert.Equal(t, []int{}, Union(nil, []int{}))
	assert.Equal(t, []int{1}, Union([]int{1}, nil))

	assert.Equal(t, []int{1, 2, 3, 4}, Union([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4}, Union([]int{1, 2, 3, 2}, []int{1, 2, 4, 3}))
	assert.Equal(t, []string{"1", "2", "3", "4"}, Union([]string{"1", "2", "3", "2"}, []string{"1", "2", "4", "3"}))
}

// nolint: forcetypeassert
func Test_UnionPred(t *testing.T) {
	assert.Equal(t, []any{}, UnionPred[any](nil, nil, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, UnionPred(nil, []any{}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1}, UnionPred([]any{1}, nil, func(t any) int { return t.(int) }))

	assert.Equal(t, []any{1, 2, 3, 4}, UnionPred([]any{1, 2}, []any{3, 4},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1, 2, 3, 4}, UnionPred([]any{1, 2, 3, 2}, []any{1, 2, 4, 3},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{"1", "2", "3", "4"},
		UnionPred([]any{"1", "2", "3", "2"}, []any{"1", "2", "4", "3"},
			func(t any) string { return t.(string) }))
}

func Test_Intersection(t *testing.T) {
	assert.Equal(t, []int{}, Intersection[int](nil, nil))
	assert.Equal(t, []int{}, Intersection(nil, []int{}))
	assert.Equal(t, []int{}, Intersection([]int{1}, nil))
	assert.Equal(t, []int{}, Intersection(nil, []int{1}))

	assert.Equal(t, []int{}, Intersection([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{1, 2, 3}, Intersection([]int{1, 2, 3, 2}, []int{1, 2, 4, 3}))
	assert.Equal(t, []string{"1", "2", "3"}, Intersection([]string{"1", "2", "3", "2"}, []string{"1", "2", "4", "3"}))
}

// nolint: forcetypeassert
func Test_IntersectionPred(t *testing.T) {
	assert.Equal(t, []any{}, IntersectionPred[any](nil, nil, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, IntersectionPred(nil, []any{}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, IntersectionPred([]any{1}, nil, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, IntersectionPred(nil, []any{1}, func(t any) int { return t.(int) }))

	assert.Equal(t, []any{}, IntersectionPred([]any{1, 2}, []any{3, 4},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1, 2, 3}, IntersectionPred([]any{1, 2, 3, 2}, []any{1, 2, 4, 3},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{"1", "2", "3"},
		IntersectionPred([]any{"1", "2", "3", "2"}, []any{"1", "2", "4", "3"},
			func(t any) string { return t.(string) }))
}

func Test_Difference(t *testing.T) {
	l, r := Difference[int](nil, nil)
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
func Test_DifferencePred(t *testing.T) {
	l, r := DifferencePred[any](nil, nil, func(t any) int { return t.(int) })
	assert.Equal(t, []any{}, l)
	assert.Equal(t, []any{}, r)
	l, r = DifferencePred([]any{}, nil, func(t any) int { return t.(int) })
	assert.Equal(t, []any{}, l)
	assert.Equal(t, []any{}, r)

	l, r = DifferencePred([]any{1, 2}, []any{3, 4}, func(t any) int { return t.(int) })
	assert.Equal(t, []any{1, 2}, l)
	assert.Equal(t, []any{3, 4}, r)

	l, r = DifferencePred([]any{1, 2, 3, 2}, []any{1, 4, 4, 3}, func(t any) int { return t.(int) })
	assert.Equal(t, []any{2, 2}, l)
	assert.Equal(t, []any{4, 4}, r)

	l, r = DifferencePred([]any{"1", "2", "3", "2"}, []any{"1", "4", "2", "", "3"},
		func(t any) string { return t.(string) })
	assert.Equal(t, []any{}, l)
	assert.Equal(t, []any{"4", ""}, r)
}
