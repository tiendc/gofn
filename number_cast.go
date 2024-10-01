package gofn

import (
	"math"
)

// ToSafeInt8 safely cast the value to type int8 (not using reflection)
func ToSafeInt8[T IntExt | UIntExt](v T) (int8, error) {
	if (v >= 0 && uint64(v) > math.MaxInt8) || (v < 0 && int64(v) < math.MinInt8) {
		return int8(v), ErrOverflow
	}
	return int8(v), nil
}

// ToSafeInt16 safely cast the value to type int16 (not using reflection)
func ToSafeInt16[T IntExt | UIntExt](v T) (int16, error) {
	if (v >= 0 && uint64(v) > math.MaxInt16) || (v < 0 && int64(v) < math.MinInt16) {
		return int16(v), ErrOverflow
	}
	return int16(v), nil
}

// ToSafeInt32 safely cast the value to type int32 (not using reflection)
func ToSafeInt32[T IntExt | UIntExt](v T) (int32, error) {
	if (v >= 0 && uint64(v) > math.MaxInt32) || (v < 0 && int64(v) < math.MinInt32) {
		return int32(v), ErrOverflow
	}
	return int32(v), nil
}

// ToSafeInt64 safely cast the value to type int64 (not using reflection)
func ToSafeInt64[T IntExt | UIntExt](v T) (int64, error) {
	if v >= 0 && uint64(v) > math.MaxInt64 {
		return int64(v), ErrOverflow
	}
	// If v < 0, it's always safe to cast to int64
	return int64(v), nil
}

// ToSafeUint8 safely cast the value to type uint8 (not using reflection)
func ToSafeUint8[T IntExt | UIntExt](v T) (uint8, error) {
	if v < 0 || uint64(v) > math.MaxUint8 {
		return uint8(v), ErrOverflow
	}
	return uint8(v), nil
}

// ToSafeUint16 safely cast the value to type uint16 (not using reflection)
func ToSafeUint16[T IntExt | UIntExt](v T) (uint16, error) {
	if v < 0 || uint64(v) > math.MaxUint16 {
		return uint16(v), ErrOverflow
	}
	return uint16(v), nil
}

// ToSafeUint32 safely cast the value to type uint32 (not using reflection)
func ToSafeUint32[T IntExt | UIntExt](v T) (uint32, error) {
	if v < 0 || uint64(v) > math.MaxUint32 {
		return uint32(v), ErrOverflow
	}
	return uint32(v), nil
}

// ToSafeUint64 safely cast the value to type uint64 (not using reflection)
func ToSafeUint64[T IntExt | UIntExt](v T) (uint64, error) {
	if v < 0 {
		return uint64(v), ErrOverflow
	}
	return uint64(v), nil
}
