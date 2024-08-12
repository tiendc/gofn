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

func Test_Or(t *testing.T) {
	// Primitive types
	assert.Equal(t, false, Or(false, false))
	assert.Equal(t, true, Or(false, true, false))
	assert.Equal(t, 1, Or(1))
	assert.Equal(t, 1, Or(1, 0, 2))
	assert.Equal(t, -1.5, Or[float64](0, -1.5))
	assert.Equal(t, float32(0), Or[float32](0, 0.0, 0))
	assert.Equal(t, int64(-2), Or[int64](0, 0, -2, 0))
	assert.Equal(t, byte(1), Or[byte](1, 0, 2, 0))
	assert.Equal(t, "", Or("", ""))
	assert.Equal(t, "f", Or("", "f", ""))

	// Pointer to primitive types
	assert.Equal(t, (*int)(nil), Or[*int](nil, nil))
	f1, f2 := float32(0), float32(1)
	assert.Equal(t, &f1, Or(nil, &f1, &f2, nil))
	s1, s2 := "", "1"
	assert.Equal(t, &s1, Or(nil, &s1, &s2, nil))

	// Derived type
	type X string
	assert.Equal(t, X("f"), Or[X]("", "f", "g"))
}

func Test_FirstTrue(t *testing.T) {
	assert.Equal(t, -1, FirstTrue(0, 0, -1, 2, 3))
	assert.Equal(t, "a", FirstTrue("", "", "a", "b"))
	assert.Equal(t, " ", FirstTrue("", "", " ", "b"))
	assert.Equal(t, []int{1}, FirstTrue([]int{}, []int{}, nil, []int{1}, []int{2, 3}))
	assert.Equal(t, map[int]int{1: 1}, FirstTrue(map[int]int{}, nil, map[int]int{1: 1}, map[int]int{2: 2}))
	assert.Nil(t, FirstTrue[*int](nil, nil, nil))
	int1, int2 := 1, 2
	assert.Equal(t, &int2, FirstTrue[*int](nil, nil, &int2, &int1))

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
	assert.Equal(t, []int{1}, FirstTrue[any](false, "", 0, 0.0, cp, Str(""), A{}, B{}, struct{}{},
		nil, ch, dt, iface, &[]int{}, []string{}, map[int]int{}, []int{1}, "x"))
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

func Test_New(t *testing.T) {
	assert.Equal(t, 3, *New(3))
	assert.Equal(t, "abc", *New("abc"))
}

func Test_Head(t *testing.T) {
	assert.Equal(t, 1, Head(1))
	assert.Equal(t, 1, Head(1, 2.0, "3", 1))
}

func Test_HeadOf(t *testing.T) {
	v1, f := HeadOf([]string{})
	assert.True(t, v1 == "" && !f)
	v2, f := HeadOf([]int(nil))
	assert.True(t, v2 == 0 && !f)
	v3, f := HeadOf([]int{1, 2, 3})
	assert.True(t, v3 == 1 && f)
}

func Test_Tail(t *testing.T) {
	t1, _ := Tail[int](-1)
	assert.Equal(t, -1, t1)
	t2, _ := Tail[string](1, 2.0, "3", "-1")
	assert.Equal(t, "-1", t2)
}

func Test_TailOf(t *testing.T) {
	v1, f := TailOf([]string{})
	assert.True(t, v1 == "" && !f)
	v2, f := TailOf([]int(nil))
	assert.True(t, v2 == 0 && !f)
	v3, f := TailOf([]int{1, 2, 3})
	assert.True(t, v3 == 3 && f)
}
