package gofn

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// RuneLength alias of utf8.RuneCountInString
var RuneLength = utf8.RuneCountInString

// StringJoin join elements from a slice of any type.
// This function calls fmt.Sprintf("%v", elem) to format every element.
func StringJoin[T any, S ~[]T](s S, sep string) string {
	return StringJoinEx(s, sep, "%v")
}

// StringJoinEx join elements from a slice of any type with custom format string
func StringJoinEx[T any, S ~[]T](s S, sep, format string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return stringFormat(format, s[0])
	}

	ss := make([]string, 0, len(s))
	for i := range s {
		ss = append(ss, stringFormat(format, s[i]))
	}
	return strings.Join(ss, sep)
}

// StringJoinBy join elements from a slice of any type with custom format function
func StringJoinBy[T any, S ~[]T](s S, sep string, fmtFunc func(v T) string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return fmtFunc(s[0])
	}

	ss := make([]string, 0, len(s))
	for i := range s {
		ss = append(ss, fmtFunc(s[i]))
	}
	return strings.Join(ss, sep)
}

// Deprecated: use StringJoinBy instead
func StringJoinPred[T any, S ~[]T](s S, sep string, fmtFunc func(v T) string) string {
	return StringJoinBy(s, sep, fmtFunc)
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
