package gofn

import (
	"math/rand"
)

var (
	StrLowerAlpha   = []rune("abcdefghijklmnopqrstuvwxyz")
	StrUpperAlpha   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	StrDigits       = []rune("0123456789")
	StrDefaultChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

// RandString generates a random string
func RandString(n int) string {
	return RandStringEx(n, StrDefaultChars)
}

// RandStringEx generates a random string
func RandStringEx(n int, allowedChars []rune) string {
	b := make([]rune, n)
	numChars := len(allowedChars)
	for i := range b {
		b[i] = allowedChars[rand.Intn(numChars)] // nolint: gosec
	}
	return string(b)
}
