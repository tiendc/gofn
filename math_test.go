package gofn

import (
	"errors"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_All(t *testing.T) {
	assert.True(t, All[int]())
	assert.True(t, All[bool]())
	assert.True(t, All(true, true, true))
	assert.True(t, All(1, -1, 2))

	assert.False(t, All(true, false, true))
	assert.False(t, All(1, -1, 2, 0))
}

func Test_Any(t *testing.T) {
	assert.True(t, Any(true, false, false))
	assert.True(t, Any(0, -1, 2, 0))
	assert.True(t, Any(0, -1, 0, 0))

	assert.False(t, Any[int]())
	assert.False(t, Any[bool]())
}

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

func Test_MinMax(t *testing.T) {
	m1, m2 := MinMax(0, 2, -10, -5, 3, 5)
	assert.True(t, m1 == -10 && m2 == 5)
	f1, f2 := MinMax[float32](0.1, -0.2, 0, 0, -0.2, 10)
	assert.True(t, f1 == -0.2 && f2 == 10)
	s1, s2 := MinMax("", "1", "A", "a")
	assert.True(t, s1 == "" && s2 == "a")
}
