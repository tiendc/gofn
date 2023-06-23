package gofn

import (
	"math/rand"
	"strings"
	"unicode/utf8"
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
	if as, ok := makeASCIISet(cutset); ok {
		ret := make([]byte, 0, len(s))
		newLineFound := true
		for i := range s {
			ch := s[i]
			if newLineFound && as.contains(ch) {
				continue
			}
			newLineFound = ch == '\n' || ch == '\r'
			ret = append(ret, ch)
		}
		return string(ret)
	}

	// Process string as runes
	runes := []rune(s)
	ret := make([]rune, 0, len(runes))
	newLineFound := true
	for _, ch := range runes {
		if newLineFound && strings.ContainsRune(cutset, ch) {
			continue
		}
		newLineFound = ch == '\n' || ch == '\r'
		ret = append(ret, ch)
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
	if as, ok := makeASCIISet(cutset); ok {
		length := len(s)
		ret := make([]byte, length)
		i, j := length-1, length-1
		newLineFound := true
		for ; i >= 0; i-- {
			ch := s[i]
			if newLineFound && as.contains(ch) {
				continue
			}
			newLineFound = ch == '\n' || ch == '\r'
			ret[j] = ch
			j--
		}
		return string(ret[j+1:])
	}

	// Process string as runes
	runes := []rune(s)
	length := len(runes)
	ret := make([]rune, length)
	i, j := length-1, length-1
	newLineFound := true
	for ; i >= 0; i-- {
		ch := runes[i]
		if newLineFound && strings.ContainsRune(cutset, ch) {
			continue
		}
		newLineFound = ch == '\n' || ch == '\r'
		ret[j] = ch
		j--
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

// These code are copied form Go strings source code.
// asciiSet is a 32-byte value, where each bit represents the presence of a
// given ASCII character in the set.
type asciiSet [8]uint32

// makeASCIISet creates a set of ASCII characters and reports whether all
// characters in chars are ASCII.
func makeASCIISet(chars string) (as asciiSet, ok bool) {
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c >= utf8.RuneSelf {
			return as, false
		}
		as[c/32] |= 1 << (c % 32)
	}
	return as, true
}

// contains reports whether c is inside the set.
func (as *asciiSet) contains(c byte) bool {
	return (as[c/32] & (1 << (c % 32))) != 0
}
