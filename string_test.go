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

func Test_LinesTrimLeftSpace(t *testing.T) {
	assert.Equal(t, "line-1\nline-2", LinesTrimLeftSpace(`line-1
		line-2`))
	assert.Equal(t, "\nline-1\nline-2", LinesTrimLeftSpace(`
		line-1
		line-2`))
	assert.Equal(t, "\nline-1\nline-2\n", LinesTrimLeftSpace(`
		line-1
		line-2
	`))
	assert.Equal(t, "\nx  line-1\nline-2", LinesTrimLeftSpace(`
		  x  line-1
		line-2`))
	// Unicode
	assert.Equal(t, "\nâ  line-1\nline-2 â", LinesTrimLeftSpace(`
		  â  line-1
		line-2 â`))
	// Extra func test
	assert.Equal(t, "\nâ  line-1\nline-2 â ê", LinesTrimLeft(`
		 ê â  line-1
		 ê line-2 â ê`, []rune{' ', '\t', 'ê'}))
}

func Test_LinesTrimRightSpace(t *testing.T) {
	assert.Equal(t, "line-1\nline-2", LinesTrimRightSpace(`line-1
line-2`))
	assert.Equal(t, "\nline-1\nline-2\n", LinesTrimRightSpace(`
line-1   
line-2
	`))
	assert.Equal(t, "\nline-1\nline-2\n", LinesTrimRightSpace(`
line-1
line-2
	`))
	assert.Equal(t, "\nx  line-1  x\nline-2", LinesTrimRightSpace(`
x  line-1  x  
line-2`))
	// Unicode
	assert.Equal(t, "\nâ  line-1  â\nline-2", LinesTrimRightSpace(`
â  line-1  â  
line-2`))
	// Extra func test
	assert.Equal(t, "\nâ  line-1  â\nline-2", LinesTrimRight(`
â  line-1  â  ê  
line-2  ê  `, []rune{' ', '\t', 'ê'}))
}
