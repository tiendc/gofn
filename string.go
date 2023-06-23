package gofn

import (
	"math/rand"
	"strings"
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

// LinesTrimLeft trim leading characters for every line in the given string
func LinesTrimLeft(s string, cutset string) string {
	if s == "" || cutset == "" {
		return s
	}
	ret := make([]byte, 0, len(s))
	s = strings.TrimLeft(s, cutset) // trim the first line
	for i := 0; i < len(s); i++ {
		ch := s[i]
		ret = append(ret, ch)
		if ch == '\n' || ch == '\r' {
			s = strings.TrimLeft(s[i+1:], cutset)
			i = -1
		}
	}
	return string(ret)
}

// LinesTrimLeftSpace trim leading spaces for every line in the given string
func LinesTrimLeftSpace(s string) string {
	// See unicode.IsSpace for what are considered spaces
	return LinesTrimLeft(s, string([]rune{' ', '\t', '\v', '\f', 0x85, 0xA0}))
}

// LinesTrimRight trim trailing characters for every line in the given string
func LinesTrimRight(s string, cutset string) string {
	if s == "" || cutset == "" {
		return s
	}

	ret := make([]byte, len(s))
	j := len(ret) - 1

	s = strings.TrimRight(s, cutset) // trim the last line
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		ret[j] = ch
		j--
		if ch == '\n' || ch == '\r' {
			s = strings.TrimRight(s[:i], cutset)
			i = len(s)
		}
	}
	return string(ret[j+1:])
}

// LinesTrimRightSpace trim trailing characters for every line in the given string
func LinesTrimRightSpace(s string) string {
	// See unicode.IsSpace for what are considered spaces
	return LinesTrimRight(s, string([]rune{' ', '\t', '\v', '\f', 0x85, 0xA0}))
}

// LinesTrim trim leading and trailing characters for every line in the given string
func LinesTrim(s string, cutset string) string {
	return LinesTrimLeft(LinesTrimRight(s, cutset), cutset)
}

// LinesTrimSpace trim leading and trailing spaces for every line in the given string
func LinesTrimSpace(s string) string {
	return LinesTrim(s, string([]rune{' ', '\t', '\v', '\f', 0x85, 0xA0}))
}

var MultilineString = LinesTrimLeftSpace
