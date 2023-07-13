package gofn

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Min(t *testing.T) {
	assert.Equal(t, -10, Min(0, 2, -10, -5, 3, 5))
	assert.Equal(t, float32(-0.2), Min[float32](0.1, -0.2, 0, 0, -0.2, 10))
	assert.Equal(t, "", Min("", "1", "A", "a"))
	assert.Equal(t, "Abc", Min("Abc", "aBC"))
}

func Test_MinIn(t *testing.T) {
	// Nil/Empty slices
	m1, err := MinIn[uint64](nil)
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

func Test_MinInPred(t *testing.T) {
	// Nil/Empty slices
	m1, err := MinInPred[uint64](nil, func(v1, v2 uint64) bool { return v1 < v2 })
	assert.True(t, m1 == 0 && errors.Is(err, ErrEmpty))
	m2, err := MinInPred([]int{}, func(v1, v2 int) bool { return v1 < v2 })
	assert.True(t, m2 == 0 && errors.Is(err, ErrEmpty))

	m2, err = MinInPred([]int{0, 2, -10, -5, 3, 5}, func(v1, v2 int) bool { return v1 < v2 })
	assert.True(t, err == nil && m2 == -10)

	// Float type
	m3, err := MinInPred([]float32{0.1, -0.2, 0, 0, -0.2, 10}, func(v1, v2 float32) bool { return v1 < v2 })
	assert.True(t, err == nil && m3 == -0.2)

	// String type
	m4, err := MinInPred([]string{"", "1", "A", "a"}, func(v1, v2 string) bool { return v1 < v2 })
	assert.True(t, err == nil && m4 == "")
	m4, err = MinInPred([]string{"Abc", "aBC"}, func(v1, v2 string) bool { return v1 < v2 })
	assert.True(t, err == nil && m4 == "Abc")

	// Struct type
	type st struct {
		Int int8
	}
	m5, err := MinInPred([]st{{0}, {1}, {10}, {20}, {10}}, func(v1, v2 st) bool { return v1.Int < v2.Int })
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
	m1, err := MaxIn[uint64](nil)
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

func Test_MaxInPred(t *testing.T) {
	// Nil/Empty slices
	m1, err := MaxInPred[uint64](nil, func(v1, v2 uint64) bool { return v1 < v2 })
	assert.True(t, m1 == 0 && errors.Is(err, ErrEmpty))
	m2, err := MaxInPred([]int{}, func(v1, v2 int) bool { return v1 < v2 })
	assert.True(t, m2 == 0 && errors.Is(err, ErrEmpty))

	m2, err = MaxInPred([]int{0, 2, -10, -5, 30, 5, 30}, func(v1, v2 int) bool { return v1 < v2 })
	assert.True(t, err == nil && m2 == 30)

	// Float type
	m3, err := MaxInPred([]float32{0.1, -0.2, 10.1, 0, -0.2, 10, 10.11}, func(v1, v2 float32) bool { return v1 < v2 })
	assert.True(t, err == nil && m3 == 10.11)

	// String type
	m4, err := MaxInPred([]string{"", "1", "A", "a"}, func(v1, v2 string) bool { return v1 < v2 })
	assert.True(t, err == nil && m4 == "a")
	m4, err = MaxInPred([]string{"Abc", "aBC"}, func(v1, v2 string) bool { return v1 < v2 })
	assert.True(t, err == nil && m4 == "aBC")

	// Struct type
	type st struct {
		Int int8
	}
	m5, err := MaxInPred([]st{{0}, {1}, {10}, {20}, {10}}, func(v1, v2 st) bool { return v1.Int < v2.Int })
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
