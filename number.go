package gofn

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

const (
	base10    = 10
	byte2Bits = 8

	fractionSep   = '.'
	noFractionSep = byte(0)
	groupSep      = ','
)

func ParseIntEx[T IntExt](s string, base int) (T, error) {
	var zeroT T
	v, err := strconv.ParseInt(s, base, int(unsafe.Sizeof(zeroT)*byte2Bits)) //nolint:gosec
	if err == nil {
		return T(v), nil
	}
	return zeroT, err // nolint: wrapcheck
}

func ParseInt[T IntExt](s string) (T, error) {
	return ParseIntEx[T](s, base10)
}

// ParseIntUngroup omit all grouping commas then parse the string value
func ParseIntUngroup[T IntExt](s string) (T, error) {
	return ParseIntEx[T](NumberFmtUngroup(s, groupSep), base10)
}

func ParseIntDef[T IntExt](s string, defaultVal T) T {
	v, err := strconv.ParseInt(s, base10, int(unsafe.Sizeof(defaultVal)*byte2Bits)) //nolint:gosec
	if err == nil {
		return T(v)
	}
	return defaultVal
}

func ParseUintEx[T UIntExt](s string, base int) (T, error) {
	var zeroT T
	v, err := strconv.ParseUint(s, base, int(unsafe.Sizeof(zeroT)*byte2Bits)) //nolint:gosec
	if err == nil {
		return T(v), nil
	}
	return zeroT, err // nolint: wrapcheck
}

func ParseUint[T UIntExt](s string) (T, error) {
	return ParseUintEx[T](s, base10)
}

// ParseUintUngroup omit all grouping commas then parse the string value
func ParseUintUngroup[T UIntExt](s string) (T, error) {
	return ParseUintEx[T](NumberFmtUngroup(s, groupSep), base10)
}

func ParseUintDef[T UIntExt](s string, defaultVal T) T {
	v, err := strconv.ParseUint(s, base10, int(unsafe.Sizeof(defaultVal)*byte2Bits)) //nolint:gosec
	if err == nil {
		return T(v)
	}
	return defaultVal
}

func ParseFloat[T FloatExt](s string) (T, error) {
	var zeroT T
	v, err := strconv.ParseFloat(s, int(unsafe.Sizeof(zeroT)*byte2Bits)) //nolint:gosec
	if err == nil {
		return T(v), nil
	}
	return zeroT, err // nolint: wrapcheck
}

// ParseFloatUngroup omit all grouping commas then parse the string value
func ParseFloatUngroup[T FloatExt](s string) (T, error) {
	return ParseFloat[T](NumberFmtUngroup(s, groupSep))
}

func ParseFloatDef[T FloatExt](s string, defaultVal T) T {
	v, err := strconv.ParseFloat(s, int(unsafe.Sizeof(defaultVal)*byte2Bits)) //nolint:gosec
	if err == nil {
		return T(v)
	}
	return defaultVal
}

func FormatIntEx[T IntExt](v T, format string) string {
	return fmt.Sprintf(format, v)
}

func FormatInt[T IntExt](v T) string {
	return strconv.FormatInt(int64(v), base10)
}

// FormatIntGroup format the value then group the decimal using comma
func FormatIntGroup[T IntExt](v T) string {
	s := strconv.FormatInt(int64(v), base10)
	return NumberFmtGroup(s, noFractionSep, groupSep)
}

func FormatUintEx[T UIntExt](v T, format string) string {
	return fmt.Sprintf(format, v)
}

func FormatUint[T UIntExt](v T) string {
	return strconv.FormatUint(uint64(v), base10)
}

// FormatUintGroup format the value then group the decimal using comma
func FormatUintGroup[T UIntExt](v T) string {
	return NumberFmtGroup(strconv.FormatUint(uint64(v), base10), noFractionSep, groupSep)
}

func FormatFloatEx[T FloatExt](v T, format string) string {
	return fmt.Sprintf(format, v)
}

func FormatFloat[T FloatExt](v T) string {
	return fmt.Sprintf("%f", v)
}

// FormatFloatGroup format the value then group the decimal using comma
func FormatFloatGroup[T FloatExt](v T) string {
	return NumberFmtGroup(fmt.Sprintf("%f", v), fractionSep, groupSep)
}

// FormatFloatGroupEx format the value then group the decimal using comma
func FormatFloatGroupEx[T FloatExt](v T, format string) string {
	return NumberFmtGroup(fmt.Sprintf(format, v), fractionSep, groupSep)
}

// NumberFmtGroup separate decimal groups in the value string
func NumberFmtGroup(num string, fractionSep, groupSep byte) string {
	if len(num) < 4 { //nolint:mnd
		return num
	}
	// Format as integer
	if fractionSep == 0 {
		return numberPartFmtGroup(num, groupSep)
	}
	// Format as real number
	fractionIndex := strings.IndexByte(num, fractionSep)
	if fractionIndex >= 0 {
		return numberPartFmtGroup(num[:fractionIndex], groupSep) + num[fractionIndex:]
	}
	return numberPartFmtGroup(num, groupSep)
}

// NumberFmtUngroup ungroup the value string
func NumberFmtUngroup(num string, groupSep byte) string {
	ret := make([]byte, 0, len(num))
	for i := range num {
		if num[i] == groupSep {
			continue
		}
		ret = append(ret, num[i])
	}
	return string(ret)
}

func numberPartFmtGroup(s string, groupSep byte) string {
	if groupSep == 0 || !stringIsNumeric(s, true) {
		return s
	}
	buf := make([]byte, 0, len(s)+5) //nolint:mnd
	ch := s[0]
	if ch == '-' {
		buf = append(buf, ch)
		s = s[1:]
	}
	start := len(s) % 3 //nolint:mnd
	if start == 0 {
		start = 3
	}
	for i := range s {
		ch = s[i]
		if i != 0 && i == start {
			buf = append(buf, groupSep)
			start += 3
		}
		buf = append(buf, ch)
	}
	return string(buf)
}

func stringIsNumeric(s string, allowSign bool) bool {
	if allowSign && len(s) > 0 && s[0] == '-' {
		s = s[1:]
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}
