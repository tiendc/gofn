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

func Test_LinesTrimLeft(t *testing.T) {
	assert.Equal(t, "", LinesTrimLeftSpace(""))
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
	assert.Equal(t, "\nê b â  line-1\nê line-2 â ê ", LinesTrimLeft(`
		 a ê b â  line-1
		  b aê line-2 â ê `, "a b\t")) // ascii cutset
	assert.Equal(t, "\nâ  line-1\nline-2 â ê", LinesTrimLeft(`
		 ê â  line-1
		 ê line-2 â ê`, "a b\tê")) // unicode cutset
}

func Test_LinesTrimRight(t *testing.T) {
	assert.Equal(t, "", LinesTrimRightSpace(""))
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
	assert.Equal(t, "\nâ  line-1  â  ê\nline-2  ê", LinesTrimRight(`
â  line-1  â  ê  
line-2  ê  `, "a b\t")) // ascii cutset
	assert.Equal(t, "\n â  line-1  â\nline-2", LinesTrimRight(`
 â  line-1  â  ê  
line-2  ê  `, "a b\tê")) // unicode cutset
}

func Test_LinesTrim(t *testing.T) {
	assert.Equal(t, "", LinesTrimSpace(""))
	assert.Equal(t, "line-1\nline-2", LinesTrimSpace(`line-1
line-2`))
	assert.Equal(t, "\nline-1\nline-2\n", LinesTrimSpace(`
line-1   
   line-2
	`))
	assert.Equal(t, "\nline-1\nline-2\n", LinesTrimSpace(`
   line-1
   line-2
	`))
	assert.Equal(t, "\nx  line-1  x\nline-2", LinesTrimSpace(`
   x  line-1  x  
   line-2   `))
	// Unicode
	assert.Equal(t, "\nâ  line-1  â\nline-2", LinesTrimSpace(`
  â  line-1  â  
  line-2  `))
	// Extra func test
	assert.Equal(t, "\nâ  line-1  â\nline-2", LinesTrim(`
    â  line-1  â  ê  
    line-2  ê  `, "a b\tê"))
}
