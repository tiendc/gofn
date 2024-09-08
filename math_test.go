package gofn

import (
	"errors"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Abs(t *testing.T) {
	assert.Equal(t, int64(0), Abs(-0))
	assert.Equal(t, int64(100), Abs(100))
	assert.Equal(t, int64(100), Abs(-100))
	assert.Equal(t, int64(-math.MinInt32), Abs(math.MinInt32))
	assert.Equal(t, int64(math.MaxInt32), Abs(math.MaxInt32))
	assert.Equal(t, int64(math.MaxInt64), Abs(math.MaxInt64))

	// Special case
	assert.Equal(t, int64(math.MinInt64), Abs(math.MinInt64))
}

func Test_Clamp(t *testing.T) {
	assert.Equal(t, 3, Clamp(1, 3, 5))
	assert.Equal(t, 5, Clamp(7, 3, 5))
	assert.Equal(t, 3, Clamp(3, 3, 5))
	assert.Equal(t, 5, Clamp(5, 3, 5))
	assert.Equal(t, 4, Clamp(4, 3, 5))

	// Incorrect position of min and max
	assert.Equal(t, 4, Clamp(4, 5, 3))
	// Negative numbers
	assert.Equal(t, -5, Clamp(-7, -3, -5))

	// Clamp string
	assert.Equal(t, "g", Clamp("a", "g", "z"))
	assert.Equal(t, "p", Clamp("p", "g", "z"))
	assert.Equal(t, "p", Clamp("p", "z", "g"))
}

func Test_Product(t *testing.T) {
	assert.Equal(t, 0, Product[int]())
	assert.Equal(t, -6, Product(1, 2, 3, -1))
	assert.Equal(t, 0, Product(1, 2, 3, -1, 0))
	assert.Equal(t, int8(-6), Product[int8](1, 2, 3, -1))
}

func Test_ProductAs(t *testing.T) {
	assert.Equal(t, 0, ProductAs[int, int]())
	assert.Equal(t, -6, ProductAs[int](1, 2, 3, -1))
	assert.Equal(t, 0, ProductAs[int](1, 2, 3, -1, 0))
	assert.Equal(t, int64(6000000000), ProductAs[int64](int32(1000), int32(2000), int32(3000)))
	// Overflow
	assert.Equal(t, int32(1705032704), ProductAs[int32](int32(1000), int32(2000), int32(3000)))
}

func Test_Sum(t *testing.T) {
	assert.Equal(t, 0, Sum[int]())
	assert.Equal(t, 5, Sum(1, 2, 3, -1))
	assert.Equal(t, int8(5), Sum[int8](1, 2, 3, -1))
}

func Test_SumAs(t *testing.T) {
	assert.Equal(t, 0, SumAs[int, int]())
	assert.Equal(t, 5, SumAs[int](1, 2, 3, -1))
	assert.Equal(t, int64(5000000000), SumAs[int64](int32(1000000000), int32(2000000000), int32(2000000000)))
	// Overflow
	assert.Equal(t, int32(705032704), SumAs[int32](int32(1000000000), int32(2000000000), int32(2000000000)))
}

func Test_Min(t *testing.T) {
	assert.Equal(t, -10, Min(0, 2, -10, -5, 3, 5))
	assert.Equal(t, float32(-0.2), Min[float32](0.1, -0.2, 0, 0, -0.2, 10))
	assert.Equal(t, "", Min("", "1", "A", "a"))
	assert.Equal(t, "Abc", Min("Abc", "aBC"))
}

func Test_MinIn(t *testing.T) {
	// Nil/Empty slices
	m1, err := MinIn[uint64]([]uint64(nil))
	assert.True(t, m1 == 0 && errors.Is(err, ErrEmpty))
	m2, err := MinIn([]int{})
	assert.True(t, m2 == 0 && errors.Is(err, ErrEmpty))

	m2, err = MinIn([]int{0, 2, -10, -5, 3, 5})
	assert.True(t, err == nil && m2 == -10)

	// Float type
	m3, err := MinIn([]float32{0.1, -0.2, 0, 0, -0.2, 10})
	assert.True(t, err == nil && m3 == -0.2)

	// String type
	m4, err := MinIn([]string{"", "1", "A", "a"})
	assert.True(t, err == nil && m4 == "")
	m4, err = MinIn([]string{"Abc", "aBC"})
	assert.True(t, err == nil && m4 == "Abc")
}

func Test_MinInBy(t *testing.T) {
	// Nil/Empty slices
	m1, err := MinInBy[uint64]([]uint64(nil), func(v1, v2 uint64) bool { return v1 < v2 })
	assert.True(t, m1 == 0 && errors.Is(err, ErrEmpty))
	m2, err := MinInBy([]int{}, func(v1, v2 int) bool { return v1 < v2 })
	assert.True(t, m2 == 0 && errors.Is(err, ErrEmpty))

	m2, err = MinInBy([]int{0, 2, -10, -5, 3, 5}, func(v1, v2 int) bool { return v1 < v2 })
	assert.True(t, err == nil && m2 == -10)

	// Float type
	m3, err := MinInBy([]float32{0.1, -0.2, 0, 0, -0.2, 10}, func(v1, v2 float32) bool { return v1 < v2 })
	assert.True(t, err == nil && m3 == -0.2)

	// String type
	m4, err := MinInBy([]string{"", "1", "A", "a"}, func(v1, v2 string) bool { return v1 < v2 })
	assert.True(t, err == nil && m4 == "")
	m4, err = MinInBy([]string{"Abc", "aBC"}, func(v1, v2 string) bool { return v1 < v2 })
	assert.True(t, err == nil && m4 == "Abc")

	// Struct type
	type st struct {
		Int int8
	}
	m5, err := MinInBy([]st{{0}, {1}, {10}, {20}, {10}}, func(v1, v2 st) bool { return v1.Int < v2.Int })
	assert.True(t, err == nil && m5 == st{0})
}

func Test_MinInPred_Deprecated(t *testing.T) {
	m1, err := MinInPred([]int{0, 2, -10, -5, 3, 5}, func(v1, v2 int) bool { return v1 < v2 })
	assert.True(t, err == nil && m1 == -10)
}

func Test_Max(t *testing.T) {
	assert.Equal(t, 30, Max(0, 2, -10, -5, 30, 5, 30))
	assert.Equal(t, 10.11, Max[float64](0.1, -0.2, 10.1, 0, -0.2, 10, 10.11))
	assert.Equal(t, "a", Max("", "1", "A", "a"))
	assert.Equal(t, "aBC", Max("Abc", "aBC"))
}

func Test_MaxIn(t *testing.T) {
	// Nil/Empty slices
	m1, err := MaxIn[uint64]([]uint64(nil))
	assert.True(t, m1 == 0 && errors.Is(err, ErrEmpty))
	m2, err := MaxIn([]int{})
	assert.True(t, m2 == 0 && errors.Is(err, ErrEmpty))

	m2, err = MaxIn([]int{0, 2, -10, -5, 30, 5, 30})
	assert.True(t, err == nil && m2 == 30)

	// Float type
	m3, err := MaxIn([]float32{0.1, -0.2, 10.1, 0, -0.2, 10, 10.11})
	assert.True(t, err == nil && m3 == 10.11)

	// String type
	m4, err := MaxIn([]string{"", "1", "A", "a"})
	assert.True(t, err == nil && m4 == "a")
	m4, err = MaxIn([]string{"Abc", "aBC"})
	assert.True(t, err == nil && m4 == "aBC")
}

func Test_MaxInBy(t *testing.T) {
	// Nil/Empty slices
	m1, err := MaxInBy[uint64]([]uint64(nil), func(v1, v2 uint64) bool { return v1 < v2 })
	assert.True(t, m1 == 0 && errors.Is(err, ErrEmpty))
	m2, err := MaxInBy([]int{}, func(v1, v2 int) bool { return v1 < v2 })
	assert.True(t, m2 == 0 && errors.Is(err, ErrEmpty))

	m2, err = MaxInBy([]int{0, 2, -10, -5, 30, 5, 30}, func(v1, v2 int) bool { return v1 < v2 })
	assert.True(t, err == nil && m2 == 30)

	// Float type
	m3, err := MaxInBy([]float32{0.1, -0.2, 10.1, 0, -0.2, 10, 10.11}, func(v1, v2 float32) bool { return v1 < v2 })
	assert.True(t, err == nil && m3 == 10.11)

	// String type
	m4, err := MaxInBy([]string{"", "1", "A", "a"}, func(v1, v2 string) bool { return v1 < v2 })
	assert.True(t, err == nil && m4 == "a")
	m4, err = MaxInBy([]string{"Abc", "aBC"}, func(v1, v2 string) bool { return v1 < v2 })
	assert.True(t, err == nil && m4 == "aBC")

	// Struct type
	type st struct {
		Int int8
	}
	m5, err := MaxInBy([]st{{0}, {1}, {10}, {20}, {10}}, func(v1, v2 st) bool { return v1.Int < v2.Int })
	assert.True(t, err == nil && m5 == st{20})
}

func Test_MaxInPred_Deprecated(t *testing.T) {
	m1, err := MaxInPred([]int{0, 2, -10, -5, 30, 5, 30}, func(v1, v2 int) bool { return v1 < v2 })
	assert.True(t, err == nil && m1 == 30)
}

func Test_MinMax(t *testing.T) {
	m1, m2 := MinMax(0, 2, -10, -5, 3, 5)
	assert.True(t, m1 == -10 && m2 == 5)
	f1, f2 := MinMax[float32](0.1, -0.2, 0, 0, -0.2, 10)
	assert.True(t, f1 == -0.2 && f2 == 10)
	s1, s2 := MinMax("", "1", "A", "a")
	assert.True(t, s1 == "" && s2 == "a")
}

func Test_Reduce(t *testing.T) {
	assert.Equal(t, 6, Reduce[int]([]int{1, 2, 3}, func(acc, v int) int { return acc + v }))
	assert.Equal(t, 8, Reduce[int]([]int{1, 2, 4}, func(acc, v int) int { return acc * v }))
	assert.Equal(t, 0, Reduce[int]([]int{1, 2, 0}, func(acc, v int) int { return acc * v }))
}

func Test_ReduceEx(t *testing.T) {
	assert.Equal(t, 7, ReduceEx[int]([]int{1, 2, 3}, func(acc, v, i int) int { return acc + v }, 1))
	assert.Equal(t, 8, ReduceEx[int]([]int{1, 2, 4}, func(acc, v, i int) int { return acc * v }, 1))
	assert.Equal(t, 0, ReduceEx[int]([]int{1, 2, 0}, func(acc, v, i int) int { return acc * v }, 1))
}

func Test_Partition(t *testing.T) {
	// Nil/empty input
	s1, s2 := Partition([]int(nil), func(i int, _ int) bool { return i%2 == 0 })
	assert.Equal(t, []int{}, s1)
	assert.Equal(t, []int{}, s2)
	s1, s2 = Partition([]int{}, func(i int, _ int) bool { return i%2 == 0 })
	assert.Equal(t, []int{}, s1)
	assert.Equal(t, []int{}, s2)

	s1, s2 = Partition([]int{3, 5, 7, 3}, func(i int, _ int) bool { return i%2 == 1 })
	assert.Equal(t, []int{3, 5, 7, 3}, s1)
	assert.Equal(t, []int{}, s2)

	s1, s2 = Partition([]int{3, 5, 7, 3}, func(i int, _ int) bool { return i%2 == 0 })
	assert.Equal(t, []int{}, s1)
	assert.Equal(t, []int{3, 5, 7, 3}, s2)

	s1, s2 = Partition([]int{3, 2, 5, 0, 0, 7, 3, 4}, func(i int, _ int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2, 0, 0, 4}, s1)
	assert.Equal(t, []int{3, 5, 7, 3}, s2)
}

func Test_PartitionN(t *testing.T) {
	// Zero partition
	p := PartitionN([]int{1, 2, 3}, 0, func(i int, _ int) int { return i % 2 })
	assert.Equal(t, 0, len(p))

	// Nil/empty input
	p = PartitionN([]int(nil), 3, func(i int, _ int) int { return i % 2 })
	assert.Equal(t, []int{}, p[0])
	assert.Equal(t, []int{}, p[1])
	assert.Equal(t, []int{}, p[2])
	p = PartitionN([]int{}, 3, func(i int, _ int) int { return i % 2 })
	assert.Equal(t, []int{}, p[0])
	assert.Equal(t, []int{}, p[1])
	assert.Equal(t, []int{}, p[2])

	p = PartitionN([]int{20, 30, 40, 50, 5, -1, 15, 10}, 3, func(i int, _ int) int { return i / 10 })
	assert.Equal(t, []int{5, -1}, p[0])
	assert.Equal(t, []int{15, 10}, p[1])
	assert.Equal(t, []int{20, 30, 40, 50}, p[2])
}

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
