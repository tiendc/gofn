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
func RandStringEx(n int, allowedChars ...[]rune) string {
	var letters []rune
	if len(allowedChars) == 0 {
		letters = STR_DEFAULT_CHARS
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
