package gofn

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
