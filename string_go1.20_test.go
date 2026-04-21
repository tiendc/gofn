//go:build go1.20

package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringOmitAccents(t *testing.T) {
	s, err := StringOmitAccents("")
	assert.NoError(t, err)
	assert.Equal(t, "", s)

	s, err = StringOmitAccents("ab 12")
	assert.NoError(t, err)
	assert.Equal(t, "ab 12", s)

	s, err = StringOmitAccents("trời ơi")
	assert.NoError(t, err)
	assert.Equal(t, "troi oi", s)
}
