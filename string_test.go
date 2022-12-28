package gofn

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandString(t *testing.T) {
	// Empty string
	assert.Equal(t, "", RandString(0))

	s := RandString(12)
	assert.Equal(t, 12, len(s))
	for _, ch := range s {
		assert.True(t, strings.ContainsRune(string(STR_DEFAULT_CHARS), ch))
	}
}

func Test_RandStringEx(t *testing.T) {
	// Empty string
	assert.Equal(t, "", RandStringEx(0, STR_LOWER_ALPHA))

	// Only digits
	s := RandStringEx(10, STR_DIGITS)
	assert.Equal(t, 10, len(s))
	for _, ch := range s {
		assert.True(t, strings.ContainsRune(string(STR_DIGITS), ch))
	}

	// Only alphabet
	s = RandStringEx(12, STR_LOWER_ALPHA)
	assert.Equal(t, 12, len(s))
	for _, ch := range s {
		assert.True(t, strings.ContainsRune(string(STR_LOWER_ALPHA), ch))
	}
}
