package gofn

import (
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
