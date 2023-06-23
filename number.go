package gofn

import (
	"strconv"
	"unsafe"
)

const (
	base10    = 10
	byte2Bits = 8
)

func ParseIntEx[T IntEx](s string, base int) (T, error) {
	var zeroT T
	v, err := strconv.ParseInt(s, base, int(unsafe.Sizeof(zeroT)*byte2Bits))
	if err == nil {
		return T(v), nil
	}
	return zeroT, err // nolint: wrapcheck
}

func ParseInt[T IntEx](s string) (T, error) {
	return ParseIntEx[T](s, base10)
}

func ParseIntDef[T IntEx](s string, defaultVal T) T {
	v, err := strconv.ParseInt(s, base10, int(unsafe.Sizeof(defaultVal)*byte2Bits))
	if err == nil {
		return T(v)
	}
	return defaultVal
}

func ParseUintEx[T UIntEx](s string, base int) (T, error) {
	var zeroT T
	v, err := strconv.ParseUint(s, base, int(unsafe.Sizeof(zeroT)*byte2Bits))
	if err == nil {
		return T(v), nil
	}
	return zeroT, err // nolint: wrapcheck
}

func ParseUint[T UIntEx](s string) (T, error) {
	return ParseUintEx[T](s, base10)
}

func ParseUintDef[T UIntEx](s string, defaultVal T) T {
	v, err := strconv.ParseUint(s, base10, int(unsafe.Sizeof(defaultVal)*byte2Bits))
	if err == nil {
		return T(v)
	}
	return defaultVal
}

func ParseFloat[T FloatEx](s string) (T, error) {
	var zeroT T
	v, err := strconv.ParseFloat(s, int(unsafe.Sizeof(zeroT)*byte2Bits))
	if err == nil {
		return T(v), nil
	}
	return zeroT, err // nolint: wrapcheck
}

func ParseFloatDef[T FloatEx](s string, defaultVal T) T {
	v, err := strconv.ParseFloat(s, int(unsafe.Sizeof(defaultVal)*byte2Bits))
	if err == nil {
		return T(v)
	}
	return defaultVal
}
