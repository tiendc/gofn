package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
	assert.Equal(t, []int{}, Reverse([]int{}))
	assert.Equal(t, []int64{1}, Reverse([]int64{1}))
	assert.Equal(t, []int{3, 2, 1}, Reverse([]int{1, 2, 3}))

	s := []int{-1, -2, 0, 1, 2}
	Reverse(s)
	assert.Equal(t, []int{2, 1, 0, -2, -1}, s)
}

func Test_ReverseCopy(t *testing.T) {
	assert.Equal(t, []int{}, ReverseCopy([]int{}))
	assert.Equal(t, []int64{1}, ReverseCopy([]int64{1}))
	assert.Equal(t, []int{3, 2, 1}, ReverseCopy([]int{1, 2, 3}))

	s := []int{-1, -2, 0, 1, 2}
	s2 := ReverseCopy(s)
	assert.Equal(t, []int{-1, -2, 0, 1, 2}, s)
	assert.Equal(t, []int{2, 1, 0, -2, -1}, s2)
}
