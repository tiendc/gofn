package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
