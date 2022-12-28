package gofn

import (
	"math/rand"
)

var (
	STR_LOWER_ALPHA   = []rune("abcdefghijklmnopqrstuvwxyz")
	STR_UPPER_ALPHA   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	STR_DIGITS        = []rune("0123456789")
	STR_DEFAULT_CHARS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

// RandString generates a random string
func RandString(n int) string {
	return RandStringEx(n, STR_DEFAULT_CHARS)
}

// RandStringEx generates a random string
func RandStringEx(n int, allowedChars []rune) string {
	b := make([]rune, n)
	numChars := len(allowedChars)
	for i := range b {
		b[i] = allowedChars[rand.Intn(numChars)]
	}
	return string(b)
}
