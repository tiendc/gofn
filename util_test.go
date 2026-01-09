package gofn

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_If(t *testing.T) {
	x, y := 1, 2
	assert.Equal(t, 1, If(x < y, 1, 2))
	assert.Equal(t, "b", If(x > y, "a", "b"))
}

func Test_FirstNonEmpty(t *testing.T) {
	assert.Equal(t, -1, FirstNonEmpty(0, 0, -1, 2, 3))
	assert.Equal(t, "a", FirstNonEmpty("", "", "a", "b"))
	assert.Equal(t, " ", FirstNonEmpty("", "", " ", "b"))
	assert.Equal(t, []int{1}, FirstNonEmpty([]int{}, []int{}, nil, []int{1}, []int{2, 3}))
	assert.Equal(t, map[int]int{1: 1}, FirstNonEmpty(map[int]int{}, nil, map[int]int{1: 1}, map[int]int{2: 2}))
	assert.Nil(t, FirstNonEmpty[*int](nil, nil, nil))
	int1, int2 := 1, 2
	assert.Equal(t, &int2, FirstNonEmpty[*int](nil, nil, &int2, &int1))

	type Str string
	type A struct {
		I  int
		S  string
		SS []int
	}
	type B struct {
		I int
		A A
	}
	ch := make(chan int)
	dt := time.Time{}
	cp := complex(0, 0)
	var iface any
	assert.Equal(t, []int{1}, FirstNonEmpty[any](false, "", 0, 0.0, cp, Str(""), A{}, B{}, struct{}{},
		nil, ch, dt, iface, &[]int{}, []string{}, map[int]int{}, []int{1}, "x"))
}

func Test_FirstTrue_Deprecated(t *testing.T) {
	assert.Equal(t, -1, FirstNonEmpty(0, 0, -1, 2, 3))
	assert.Equal(t, "a", FirstNonEmpty("", "", "a", "b"))
	assert.Equal(t, " ", FirstNonEmpty("", "", " ", "b"))
	assert.Equal(t, []int{1}, FirstNonEmpty([]int{}, []int{}, nil, []int{1}, []int{2, 3}))
	assert.Equal(t, map[int]int{1: 1}, FirstNonEmpty(map[int]int{}, nil, map[int]int{1: 1}, map[int]int{2: 2}))
	assert.Nil(t, FirstNonEmpty[*int](nil, nil, nil))
}

func Test_Coalesce(t *testing.T) {
	// Primitive types
	assert.Equal(t, false, Coalesce(false, false))
	assert.Equal(t, true, Coalesce(false, true, false))
	assert.Equal(t, 1, Coalesce(1))
	assert.Equal(t, 1, Coalesce(1, 0, 2))
	assert.Equal(t, -1.5, Coalesce[float64](0, -1.5))
	assert.Equal(t, float32(0), Coalesce[float32](0, 0.0, 0))
	assert.Equal(t, int64(-2), Coalesce[int64](0, 0, -2, 0))
	assert.Equal(t, byte(1), Coalesce[byte](1, 0, 2, 0))
	assert.Equal(t, "", Coalesce("", ""))
	assert.Equal(t, "f", Coalesce("", "f", ""))

	// Pointer to primitive types
	assert.Equal(t, (*int)(nil), Coalesce[*int](nil, nil))
	f1, f2 := float32(0), float32(1)
	assert.Equal(t, &f1, Coalesce(nil, &f1, &f2, nil))
	s1, s2 := "", "1"
	assert.Equal(t, &s1, Coalesce(nil, &s1, &s2, nil))

	// Derived type
	type X string
	assert.Equal(t, X("f"), Coalesce[X]("", "f", "g"))
}

// nolint: goerr113, forcetypeassert
func Test_Must1(t *testing.T) {
	Must1(func() error { return nil }())

	// Panic case: error
	defer func() {
		e := recover()
		assert.True(t, e != nil && e.(error).Error() == "error 1")
	}()
	Must1(func() error { return errors.New("error 1") }())
}

// nolint: goerr113, forcetypeassert
func Test_Must2(t *testing.T) {
	assert.Equal(t, 1, Must2(func() (int, error) { return 1, nil }()))
	assert.Equal(t, "a", Must2(func() (string, error) { return "a", nil }()))

	// Panic case: error
	defer func() {
		e := recover()
		assert.True(t, e != nil && e.(error).Error() == "error 2")
	}()
	assert.Equal(t, 0, Must2(func() (int, error) { return 0, errors.New("error 2") }()))
}

// nolint: goerr113, forcetypeassert
func Test_Must(t *testing.T) {
	assert.Equal(t, 1, Must(func() (int, error) { return 1, nil }()))
	assert.Equal(t, "a", Must(func() (string, error) { return "a", nil }()))
}

// nolint: goerr113, forcetypeassert
func Test_Must3(t *testing.T) {
	v1, v2 := Must3(func() (int, bool, error) { return 1, true, nil }())
	assert.True(t, v1 == 1 && v2 == true)

	// Panic case: error
	defer func() {
		e := recover()
		assert.True(t, e != nil && e.(error).Error() == "error 3")
	}()
	_, _ = Must3(func() (int, bool, error) { return 0, true, errors.New("error 3") }())
}

// nolint: goerr113, forcetypeassert, dogsled
func Test_Must4(t *testing.T) {
	v1, v2, v3 := Must4(func() (int, bool, string, error) { return 1, true, "x", nil }())
	assert.True(t, v1 == 1 && v2 == true && v3 == "x")

	// Panic case: error
	defer func() {
		e := recover()
		assert.True(t, e != nil && e.(error).Error() == "error 4")
	}()
	_, _, _ = Must4(func() (int, bool, string, error) { return 0, true, "", errors.New("error 4") }())
}

// nolint: goerr113, forcetypeassert, dogsled
func Test_Must5(t *testing.T) {
	v1, v2, v3, v4 := Must5(func() (int, bool, string, float32, error) { return 1, true, "x", 2.1, nil }())
	assert.True(t, v1 == 1 && v2 == true && v3 == "x" && v4 == 2.1)

	// Panic case: error
	defer func() {
		e := recover()
		assert.True(t, e != nil && e.(error).Error() == "error 5")
	}()
	_, _, _, _ = Must5(func() (int, bool, string, float32, error) { return 0, true, "", 2.1, errors.New("error 5") }())
}

// nolint: goerr113, forcetypeassert, dogsled
func Test_Must6(t *testing.T) {
	v1, v2, v3, v4, v5 := Must6(func() (int, bool, string, float32, int64, error) { return 1, true, "x", 2.1, 12, nil }())
	assert.True(t, v1 == 1 && v2 == true && v3 == "x" && v4 == 2.1 && v5 == 12)

	// Panic case: error
	defer func() {
		e := recover()
		assert.True(t, e != nil && e.(error).Error() == "error 6")
	}()
	_, _, _, _, _ = Must6(func() (int, bool, string, float32, int64, error) {
		return 0, true, "", 2.1, 12, errors.New("error 6")
	}())
}

func Test_ToPtr(t *testing.T) {
	assert.Equal(t, 3, *ToPtr(3))
	assert.Equal(t, "abc", *ToPtr("abc"))
}

func Test_PtrValueOrEmpty(t *testing.T) {
	assert.Equal(t, 0, PtrValueOrEmpty[int](nil))
	assert.Nil(t, PtrValueOrEmpty[*int](nil))
	assert.Equal(t, 3, PtrValueOrEmpty(ToPtr(3)))
	assert.Equal(t, "abc", PtrValueOrEmpty(ToPtr("abc")))
}

func Test_Head(t *testing.T) {
	assert.Equal(t, 1, Head(1))
	assert.Equal(t, 1, Head(1, 2.0, "3", 1))
}

func Test_Tail(t *testing.T) {
	t1, _ := Tail[int](-1)
	assert.Equal(t, -1, t1)
	t2, _ := Tail[string](1, 2.0, "3", "-1")
	assert.Equal(t, "-1", t2)
}
