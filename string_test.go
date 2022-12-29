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
		assert.True(t, strings.ContainsRune(string(StrDefaultChars), ch))
	}
}

func Test_RandStringEx(t *testing.T) {
	// Empty string
	assert.Equal(t, "", RandStringEx(0, StrLowerAlpha))

	// Only digits
	s := RandStringEx(10, StrDigits)
	assert.Equal(t, 10, len(s))
	for _, ch := range s {
		assert.True(t, strings.ContainsRune(string(StrDigits), ch))
	}

	// Only alphabet
	s = RandStringEx(12, StrLowerAlpha)
	assert.Equal(t, 12, len(s))
	for _, ch := range s {
		assert.True(t, strings.ContainsRune(string(StrLowerAlpha), ch))
	}
}
