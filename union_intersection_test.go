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

func Test_UnionPred(t *testing.T) {
	assert.Equal(t, []interface{}{}, UnionPred[interface{}](nil, nil, func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{}, UnionPred(nil, []interface{}{}, func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{1}, UnionPred([]interface{}{1}, nil, func(t interface{}) int { return t.(int) }))

	assert.Equal(t, []interface{}{1, 2, 3, 4}, UnionPred([]interface{}{1, 2}, []interface{}{3, 4},
		func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{1, 2, 3, 4}, UnionPred([]interface{}{1, 2, 3, 2}, []interface{}{1, 2, 4, 3},
		func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{"1", "2", "3", "4"}, UnionPred([]interface{}{"1", "2", "3", "2"}, []interface{}{"1", "2", "4", "3"},
		func(t interface{}) string { return t.(string) }))
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

func Test_IntersectionPred(t *testing.T) {
	assert.Equal(t, []interface{}{}, IntersectionPred[interface{}](nil, nil, func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{}, IntersectionPred(nil, []interface{}{}, func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{}, IntersectionPred([]interface{}{1}, nil, func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{}, IntersectionPred(nil, []interface{}{1}, func(t interface{}) int { return t.(int) }))

	assert.Equal(t, []interface{}{}, IntersectionPred([]interface{}{1, 2}, []interface{}{3, 4},
		func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{1, 2, 3}, IntersectionPred([]interface{}{1, 2, 3, 2}, []interface{}{1, 2, 4, 3},
		func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{"1", "2", "3"}, IntersectionPred([]interface{}{"1", "2", "3", "2"}, []interface{}{"1", "2", "4", "3"},
		func(t interface{}) string { return t.(string) }))
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

func Test_DifferencePred(t *testing.T) {
	l, r := DifferencePred[interface{}](nil, nil, func(t interface{}) int { return t.(int) })
	assert.Equal(t, []interface{}{}, l)
	assert.Equal(t, []interface{}{}, r)
	l, r = DifferencePred([]interface{}{}, nil, func(t interface{}) int { return t.(int) })
	assert.Equal(t, []interface{}{}, l)
	assert.Equal(t, []interface{}{}, r)

	l, r = DifferencePred([]interface{}{1, 2}, []interface{}{3, 4}, func(t interface{}) int { return t.(int) })
	assert.Equal(t, []interface{}{1, 2}, l)
	assert.Equal(t, []interface{}{3, 4}, r)

	l, r = DifferencePred([]interface{}{1, 2, 3, 2}, []interface{}{1, 4, 4, 3}, func(t interface{}) int { return t.(int) })
	assert.Equal(t, []interface{}{2, 2}, l)
	assert.Equal(t, []interface{}{4, 4}, r)

	l, r = DifferencePred([]interface{}{"1", "2", "3", "2"}, []interface{}{"1", "4", "2", "", "3"},
		func(t interface{}) string { return t.(string) })
	assert.Equal(t, []interface{}{}, l)
	assert.Equal(t, []interface{}{"4", ""}, r)
}
