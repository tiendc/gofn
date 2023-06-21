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
	var intf interface{}
	assert.Equal(t, []int{1}, FirstTrue[interface{}](false, "", 0, 0.0, cp, Str(""), A{}, B{}, struct{}{},
		nil, ch, dt, intf, []int{}, []string{}, map[int]int{}, []int{1}, "x"))
}

// nolint: goerr113, forcetypeassert
func Test_Must(t *testing.T) {
	assert.Equal(t, 1, Must(func() (int, error) { return 1, nil }()))
	assert.Equal(t, "a", Must(func() (string, error) { return "a", nil }()))

	// Panic case: error
	defer func() {
		e := recover()
		assert.True(t, e != nil && e.(error).Error() == "error")
	}()
	assert.Equal(t, 0, Must(func() (int, error) { return 0, errors.New("error") }()))
}

func Test_New(t *testing.T) {
	assert.Equal(t, 3, *New(3))
	assert.Equal(t, "abc", *New("abc"))
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
