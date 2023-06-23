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

// LinesTrimLeft trim leading characters for every line in the given string
func LinesTrimLeft(s string, cutset []rune) string {
	if len(cutset) == 0 {
		return s
	}
	runes := []rune(s)
	ret := make([]rune, 0, len(runes))
	newLineFound := true
	for _, ch := range runes {
		if newLineFound && containRune(ch, cutset) {
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
	return LinesTrimLeft(s, []rune{' ', '\t', '\v', '\f', 0x85, 0xA0})
}

// LinesTrimRight trim trailing characters for every line in the given string
func LinesTrimRight(s string, cutset []rune) string {
	if len(cutset) == 0 {
		return s
	}
	runes := []rune(s)
	length := len(runes)
	ret := make([]rune, length)
	i, j := length-1, length-1
	newLineFound := true
	for {
		if i < 0 {
			break
		}
		ch := runes[i]
		i--
		if newLineFound && containRune(ch, cutset) {
			continue
		}
		newLineFound = ch == '\n' || ch == '\r'
		ret[j] = ch
		j--
	}
	if j != -1 { // `j` go to -1 if `ret` is fully filled
		ret = ret[j+1:]
	}
	return string(ret)
}

// LinesTrimRightSpace trim trailing characters for every line in the given string
func LinesTrimRightSpace(s string) string {
	// See unicode.IsSpace for what are considered spaces
	return LinesTrimRight(s, []rune{' ', '\t', '\v', '\f', 0x85, 0xA0})
}

func containRune(r rune, s []rune) bool {
	for _, rr := range s {
		if r == rr {
			return true
		}
	}
	return false
}

var MultilineString = LinesTrimLeftSpace
