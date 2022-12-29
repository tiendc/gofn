package gofn

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If(t *testing.T) {
	x, y := 1, 2
	assert.Equal(t, 1, If(x < y, 1, 2))
	assert.Equal(t, "b", If(x > y, "a", "b"))
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
