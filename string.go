package gofn

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode/utf8"
)

var (
	StrLowerAlpha   = []rune("abcdefghijklmnopqrstuvwxyz")
	StrUpperAlpha   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	StrAlpha        = Concat(StrLowerAlpha, StrUpperAlpha)
	StrDigits       = []rune("0123456789")
	StrAlphaNumeric = Concat(StrAlpha, StrDigits)
	StrSpecialChars = []rune("~!@#$%^&*()-_+`'\";:,.<>/?{[}]\\|")
	StrAllChars     = Concat(StrAlpha, StrDigits, StrSpecialChars)
	StrDefaultChars = StrAlphaNumeric
)

// RandString generates a random string
func RandString(n int) string {
	return RandStringEx(n, StrDefaultChars)
}

// RandStringEx generates a random string
func RandStringEx[S ~[]rune](n int, allowedChars S) string {
	b := make([]rune, n)
	numChars := len(allowedChars)
	for i := range b {
		b[i] = allowedChars[rand.Intn(numChars)] // nolint: gosec
	}
	return string(b)
}

// RuneLength alias of utf8.RuneCountInString
var RuneLength = utf8.RuneCountInString

// StringJoin join elements from a slice of any type.
// This function calls fmt.Sprintf("%v", elem) to format every element.
func StringJoin[T any, S ~[]T](s S, sep string) string {
	return StringJoinEx(s, sep, "%v")
}

// StringJoinEx join elements from a slice of any type with custom format string
func StringJoinEx[T any, S ~[]T](s S, sep, format string) string {
	return StringJoinBy(s, sep, func(v T) string {
		return stringFormat(format, v)
	})
}

// StringJoinBy join elements from a slice of any type with custom format function
func StringJoinBy[T any, S ~[]T](s S, sep string, fmtFunc func(v T) string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return fmtFunc(s[0])
	}

	var sb strings.Builder
	for i := range s {
		if i > 0 {
			sb.WriteString(sep)
		}
		sb.WriteString(fmtFunc(s[i]))
	}
	return sb.String()
}

func stringFormat(format string, v any) string {
	if v == nil {
		return "null"
	}
	if stringer, ok := v.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf(format, v)
}

// StringLexJoinEx lexical joins a list of items of any type. The input format will be
// used with fmt.Sprintf() to render slice items as string.
//
// For example:
//
//	StringLexJoinEx(["grape", "apple", "orange"], ", ", " and ", "%v") -> grape, apple and orange
//	StringLexJoinEx(["grape", "apple", "orange"], ", ", " or ", "%v") -> grape, apple or orange
func StringLexJoinEx[T any, S ~[]T](s S, sep, lastSep, format string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return stringFormat(format, s[0])
	}
	var sb strings.Builder
	for i := 0; i < length-1; i++ {
		if i > 0 {
			sb.WriteString(sep)
		}
		sb.WriteString(stringFormat(format, s[i]))
	}
	sb.WriteString(lastSep)
	sb.WriteString(stringFormat(format, s[length-1]))

	return sb.String()
}

// StringLexJoin lexical joins a list of items of any type.
func StringLexJoin[T any, S ~[]T](s S, sep, lastSep string) string {
	return StringLexJoinEx(s, sep, lastSep, "%v")
}

// LinesTrimLeft trim leading characters for every line in the given string.
// Deprecated
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

// LinesTrimLeftSpace trim leading spaces for every line in the given string.
// Deprecated
func LinesTrimLeftSpace(s string) string {
	// See unicode.IsSpace for what are considered spaces
	return LinesTrimLeft(s, string([]rune{' ', '\t', '\v', '\f', 0x85, 0xA0}))
}

// LinesTrimRight trim trailing characters for every line in the given string.
// Deprecated
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

// LinesTrimRightSpace trim trailing characters for every line in the given string.
// Deprecated
func LinesTrimRightSpace(s string) string {
	// See unicode.IsSpace for what are considered spaces
	return LinesTrimRight(s, string([]rune{' ', '\t', '\v', '\f', 0x85, 0xA0}))
}

// LinesTrim trim leading and trailing characters for every line in the given string.
// Deprecated
func LinesTrim(s string, cutset string) string {
	return LinesTrimLeft(LinesTrimRight(s, cutset), cutset)
}

// LinesTrimSpace trim leading and trailing spaces for every line in the given string.
// Deprecated
func LinesTrimSpace(s string) string {
	return LinesTrim(s, string([]rune{' ', '\t', '\v', '\f', 0x85, 0xA0}))
}

// Deprecated
var MultilineString = LinesTrimLeftSpace

// StringWrap wraps a string with the given token
func StringWrap(s string, token string) string {
	return token + s + token
}

// StringUnwrap unwraps a string wrapped with the given token
func StringUnwrap(s string, token string) string {
	return strings.TrimSuffix(strings.TrimPrefix(s, token), token)
}

// StringWrapLR wraps a string with the given tokens for the left and right side
func StringWrapLR(s string, tokenLeft, tokenRight string) string {
	return tokenLeft + s + tokenRight
}

// StringUnwrapLR unwraps a string wrapped with the given tokens
func StringUnwrapLR(s string, tokenLeft, tokenRight string) string {
	return strings.TrimSuffix(strings.TrimPrefix(s, tokenLeft), tokenRight)
}
