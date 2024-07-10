package gofn

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandChoice(t *testing.T) {
	// Empty input
	v1, b1 := RandChoice[int]()
	assert.Equal(t, 0, v1)
	assert.False(t, b1)
	// One item input
	v2, b2 := RandChoice[float32](1.1)
	assert.Equal(t, float32(1.1), v2)
	assert.True(t, b2)
	// Multiple items input
	v3, b3 := RandChoice[string]("1", "2", "3")
	assert.True(t, Contain([]string{"1", "2", "3"}, v3))
	assert.True(t, b3)
}

func Test_RandChoiceMaker_Next(t *testing.T) {
	// Empty input
	maker1 := NewRandChoiceMaker[int](nil)
	assert.False(t, maker1.HasNext())
	v1, b1 := maker1.Next()
	assert.Equal(t, 0, v1)
	assert.False(t, b1)

	// One item input
	maker2 := NewRandChoiceMaker([]float32{1.1})
	assert.True(t, maker2.HasNext())
	v2, b2 := maker2.Next()
	assert.False(t, maker2.HasNext())
	assert.Equal(t, float32(1.1), v2)
	assert.True(t, b2)

	// Multiple items input (with using custom rand function)
	maker3 := NewRandChoiceMaker([]string{"1", "2", "3"}, rand.Intn)
	assert.True(t, maker3.HasNext())
	v30, b30 := maker3.Next()
	v31, b31 := maker3.Next()
	v32, b32 := maker3.Next()
	assert.False(t, maker3.HasNext())
	v33, b33 := maker3.Next()
	v34, b34 := maker3.Next()
	assert.True(t, b30 && b31 && b32)
	assert.False(t, b33 || b34)
	v3 := []string{v30, v31, v32, v33, v34}
	Sort(v3)
	assert.Equal(t, []string{"", "", "1", "2", "3"}, v3)
}

func Test_Shuffle(t *testing.T) {
	// Empty input slice
	s1 := Shuffle[int](nil)
	assert.Equal(t, []int{}, s1)
	// One item input
	s2 := Shuffle([]float32{1.1})
	assert.Equal(t, []float32{1.1}, s2)
	// Multiple items input (with using custom rand function)
	s3 := Shuffle([]string{"1", "2", "3"}, rand.Intn)
	Sort(s3)
	assert.Equal(t, []string{"1", "2", "3"}, s3)
}

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
