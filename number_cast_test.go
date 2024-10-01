package gofn

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testToSafeInt8[T IntExt | UIntExt](t *testing.T, v T, err error) {
	i8, e := ToSafeInt8(v)
	assert.Equal(t, e, err)
	assert.Equal(t, i8, int8(v))
}

func Test_ToSafeInt8(t *testing.T) {
	// int64
	testToSafeInt8(t, int64(math.MaxInt64), ErrOverflow)
	testToSafeInt8(t, int64(math.MinInt64), ErrOverflow)
	testToSafeInt8(t, int64(math.MinInt64)+1, ErrOverflow)
	testToSafeInt8(t, int64(math.MinInt8)-1, ErrOverflow)
	testToSafeInt8(t, int64(math.MaxInt8)+1, ErrOverflow)
	testToSafeInt8(t, int64(math.MaxInt8), nil)
	testToSafeInt8(t, int64(math.MinInt8), nil)
	testToSafeInt8(t, int64(0), nil)

	// uint64
	testToSafeInt8(t, uint64(math.MaxUint64), ErrOverflow)
	testToSafeInt8(t, uint64(math.MaxInt8)+1, ErrOverflow)
	testToSafeInt8(t, uint64(math.MaxInt8), nil)
	testToSafeInt8(t, uint64(0), nil)

	// int32
	testToSafeInt8(t, int32(math.MaxInt32), ErrOverflow)
	testToSafeInt8(t, int32(math.MinInt32), ErrOverflow)
	testToSafeInt8(t, int32(math.MinInt32)+1, ErrOverflow)
	testToSafeInt8(t, int32(math.MinInt8)-1, ErrOverflow)
	testToSafeInt8(t, int32(math.MaxInt8)+1, ErrOverflow)
	testToSafeInt8(t, int32(math.MaxInt8), nil)
	testToSafeInt8(t, int32(math.MinInt8), nil)
	testToSafeInt8(t, int32(0), nil)

	// uint32
	testToSafeInt8(t, uint32(math.MaxUint32), ErrOverflow)
	testToSafeInt8(t, uint32(math.MaxInt8)+1, ErrOverflow)
	testToSafeInt8(t, uint32(math.MaxInt8), nil)
	testToSafeInt8(t, uint32(0), nil)

	// int16
	testToSafeInt8(t, int16(math.MaxInt16), ErrOverflow)
	testToSafeInt8(t, int16(math.MinInt16), ErrOverflow)
	testToSafeInt8(t, int16(math.MinInt16)+1, ErrOverflow)
	testToSafeInt8(t, int16(math.MinInt8)-1, ErrOverflow)
	testToSafeInt8(t, int16(math.MaxInt8)+1, ErrOverflow)
	testToSafeInt8(t, int16(math.MaxInt8), nil)
	testToSafeInt8(t, int16(math.MinInt8), nil)
	testToSafeInt8(t, int16(0), nil)

	// uint16
	testToSafeInt8(t, uint16(math.MaxUint16), ErrOverflow)
	testToSafeInt8(t, uint16(math.MaxInt8)+1, ErrOverflow)
	testToSafeInt8(t, uint16(math.MaxInt8), nil)
	testToSafeInt8(t, uint16(0), nil)

	// int8
	testToSafeInt8(t, int8(math.MinInt8), nil)
	testToSafeInt8(t, int8(math.MinInt8)+1, nil)
	testToSafeInt8(t, int8(math.MaxInt8), nil)
	testToSafeInt8(t, int8(math.MinInt8), nil)
	testToSafeInt8(t, int8(0), nil)

	// uint8
	testToSafeInt8(t, uint8(math.MaxUint8), ErrOverflow)
	testToSafeInt8(t, uint8(math.MaxInt8)+1, ErrOverflow)
	testToSafeInt8(t, uint8(math.MaxInt8), nil)
	testToSafeInt8(t, uint8(0), nil)
}

func testToSafeInt16[T IntExt | UIntExt](t *testing.T, v T, err error) {
	i16, e := ToSafeInt16(v)
	assert.Equal(t, e, err)
	assert.Equal(t, i16, int16(v))
}

func Test_ToSafeInt16(t *testing.T) {
	// int64
	testToSafeInt16(t, int64(math.MaxInt64), ErrOverflow)
	testToSafeInt16(t, int64(math.MinInt64), ErrOverflow)
	testToSafeInt16(t, int64(math.MinInt64)+1, ErrOverflow)
	testToSafeInt16(t, int64(math.MinInt16)-1, ErrOverflow)
	testToSafeInt16(t, int64(math.MaxInt16)+1, ErrOverflow)
	testToSafeInt16(t, int64(math.MaxInt16), nil)
	testToSafeInt16(t, int64(math.MinInt16), nil)
	testToSafeInt16(t, int64(0), nil)

	// uint64
	testToSafeInt16(t, uint64(math.MaxUint64), ErrOverflow)
	testToSafeInt16(t, uint64(math.MaxInt16)+1, ErrOverflow)
	testToSafeInt16(t, uint64(math.MaxInt16), nil)
	testToSafeInt16(t, uint64(0), nil)

	// int32
	testToSafeInt16(t, int32(math.MaxInt32), ErrOverflow)
	testToSafeInt16(t, int32(math.MinInt32), ErrOverflow)
	testToSafeInt16(t, int32(math.MinInt32)+1, ErrOverflow)
	testToSafeInt16(t, int32(math.MinInt16)-1, ErrOverflow)
	testToSafeInt16(t, int32(math.MaxInt16)+1, ErrOverflow)
	testToSafeInt16(t, int32(math.MaxInt16), nil)
	testToSafeInt16(t, int32(math.MinInt16), nil)
	testToSafeInt16(t, int32(0), nil)

	// uint32
	testToSafeInt16(t, uint32(math.MaxUint32), ErrOverflow)
	testToSafeInt16(t, uint32(math.MaxInt16)+1, ErrOverflow)
	testToSafeInt16(t, uint32(math.MaxInt16), nil)
	testToSafeInt16(t, uint32(0), nil)

	// int16
	testToSafeInt16(t, int16(math.MaxInt16), nil)
	testToSafeInt16(t, int16(math.MinInt16), nil)
	testToSafeInt16(t, int16(0), nil)

	// uint16
	testToSafeInt16(t, uint16(math.MaxUint16), ErrOverflow)
	testToSafeInt16(t, uint16(math.MaxInt16), nil)
	testToSafeInt16(t, uint16(0), nil)

	// int8
	testToSafeInt16(t, int8(math.MaxInt8), nil)
	testToSafeInt16(t, int8(math.MinInt8), nil)
	testToSafeInt16(t, int8(math.MinInt8)+1, nil)
	testToSafeInt16(t, int8(0), nil)

	// uint8
	testToSafeInt16(t, uint8(math.MaxUint8), nil)
	testToSafeInt16(t, uint8(math.MaxInt8)+1, nil)
	testToSafeInt16(t, uint8(math.MaxInt8), nil)
	testToSafeInt16(t, uint8(0), nil)
}

func testToSafeInt32[T IntExt | UIntExt](t *testing.T, v T, err error) {
	i32, e := ToSafeInt32(v)
	assert.Equal(t, e, err)
	assert.Equal(t, i32, int32(v))
}

func Test_ToSafeInt32(t *testing.T) {
	// int64
	testToSafeInt32(t, int64(math.MaxInt64), ErrOverflow)
	testToSafeInt32(t, int64(math.MinInt64), ErrOverflow)
	testToSafeInt32(t, int64(math.MinInt64)+1, ErrOverflow)
	testToSafeInt32(t, int64(math.MinInt32)-1, ErrOverflow)
	testToSafeInt32(t, int64(math.MaxInt32)+1, ErrOverflow)
	testToSafeInt32(t, int64(math.MaxInt32), nil)
	testToSafeInt32(t, int64(math.MinInt32), nil)
	testToSafeInt32(t, int64(0), nil)

	// uint64
	testToSafeInt32(t, uint64(math.MaxUint64), ErrOverflow)
	testToSafeInt32(t, uint64(math.MaxInt32)+1, ErrOverflow)
	testToSafeInt32(t, uint64(math.MaxInt32), nil)
	testToSafeInt32(t, uint64(0), nil)

	// int32
	testToSafeInt32(t, int32(math.MaxInt32), nil)
	testToSafeInt32(t, int32(math.MinInt32)+1, nil)
	testToSafeInt32(t, int32(math.MinInt32), nil)
	testToSafeInt32(t, int32(0), nil)

	// uint32
	testToSafeInt32(t, uint32(math.MaxUint32), ErrOverflow)
	testToSafeInt32(t, uint32(math.MaxInt32)+1, ErrOverflow)
	testToSafeInt32(t, uint32(math.MaxInt32), nil)
	testToSafeInt32(t, uint32(0), nil)

	// int16
	testToSafeInt32(t, int16(math.MaxInt16), nil)
	testToSafeInt32(t, int16(math.MinInt16), nil)
	testToSafeInt32(t, int16(0), nil)

	// uint16
	testToSafeInt32(t, uint16(math.MaxUint16), nil)
	testToSafeInt32(t, uint16(math.MaxInt16), nil)
	testToSafeInt32(t, uint16(0), nil)

	// int8
	testToSafeInt32(t, int8(math.MaxInt8), nil)
	testToSafeInt32(t, int8(math.MinInt8), nil)
	testToSafeInt32(t, int8(math.MinInt8)+1, nil)
	testToSafeInt32(t, int8(0), nil)

	// uint8
	testToSafeInt32(t, uint8(math.MaxUint8), nil)
	testToSafeInt32(t, uint8(math.MaxInt8)+1, nil)
	testToSafeInt32(t, uint8(math.MaxInt8), nil)
	testToSafeInt32(t, uint8(0), nil)
}

func testToSafeInt64[T IntExt | UIntExt](t *testing.T, v T, err error) {
	i64, e := ToSafeInt64(v)
	assert.Equal(t, e, err)
	assert.Equal(t, i64, int64(v))
}

func Test_ToSafeInt64(t *testing.T) {
	// int64
	testToSafeInt64(t, int64(math.MaxInt64), nil)
	testToSafeInt64(t, int64(math.MinInt64), nil)
	testToSafeInt64(t, int64(math.MinInt64)+1, nil)
	testToSafeInt64(t, int64(0), nil)

	// uint64
	testToSafeInt64(t, uint64(math.MaxUint64), ErrOverflow)
	testToSafeInt64(t, uint64(math.MaxInt64)+1, ErrOverflow)
	testToSafeInt64(t, uint64(math.MaxInt64), nil)
	testToSafeInt64(t, uint64(0), nil)

	// int32
	testToSafeInt64(t, int32(math.MaxInt32), nil)
	testToSafeInt64(t, int32(math.MinInt32)+1, nil)
	testToSafeInt64(t, int32(math.MinInt32), nil)
	testToSafeInt64(t, int32(0), nil)

	// uint32
	testToSafeInt64(t, uint32(math.MaxUint32), nil)
	testToSafeInt64(t, uint32(math.MaxInt32)+1, nil)
	testToSafeInt64(t, uint32(math.MaxInt32), nil)
	testToSafeInt64(t, uint32(0), nil)

	// int16
	testToSafeInt64(t, int16(math.MaxInt16), nil)
	testToSafeInt64(t, int16(math.MinInt16), nil)
	testToSafeInt64(t, int16(0), nil)

	// uint16
	testToSafeInt64(t, uint16(math.MaxUint16), nil)
	testToSafeInt64(t, uint16(math.MaxInt16), nil)
	testToSafeInt64(t, uint16(0), nil)

	// int8
	testToSafeInt64(t, int8(math.MaxInt8), nil)
	testToSafeInt64(t, int8(math.MinInt8), nil)
	testToSafeInt64(t, int8(math.MinInt8)+1, nil)
	testToSafeInt64(t, int8(0), nil)

	// uint8
	testToSafeInt64(t, uint8(math.MaxUint8), nil)
	testToSafeInt64(t, uint8(math.MaxInt8)+1, nil)
	testToSafeInt64(t, uint8(math.MaxInt8), nil)
	testToSafeInt64(t, uint8(0), nil)
}

func testToSafeUint8[T IntExt | UIntExt](t *testing.T, v T, err error) {
	i8, e := ToSafeUint8(v)
	assert.Equal(t, e, err)
	assert.Equal(t, i8, uint8(v))
}

func Test_ToSafeUint8(t *testing.T) {
	// int64
	testToSafeUint8(t, int64(math.MaxInt64), ErrOverflow)
	testToSafeUint8(t, int64(math.MaxUint8)+1, ErrOverflow)
	testToSafeUint8(t, int64(math.MinInt64), ErrOverflow)
	testToSafeUint8(t, int64(-1), ErrOverflow)
	testToSafeUint8(t, int64(math.MaxUint8), nil)
	testToSafeUint8(t, int64(0), nil)

	// uint64
	testToSafeUint8(t, uint64(math.MaxUint64), ErrOverflow)
	testToSafeUint8(t, uint64(math.MaxUint8)+1, ErrOverflow)
	testToSafeUint8(t, uint64(math.MaxUint8), nil)
	testToSafeUint8(t, uint64(0), nil)

	// int32
	testToSafeUint8(t, int32(math.MaxInt32), ErrOverflow)
	testToSafeUint8(t, int32(math.MaxUint8)+1, ErrOverflow)
	testToSafeUint8(t, int32(math.MinInt32), ErrOverflow)
	testToSafeUint8(t, int32(-1), ErrOverflow)
	testToSafeUint8(t, int32(math.MaxUint8), nil)
	testToSafeUint8(t, int32(0), nil)

	// uint32
	testToSafeUint8(t, uint32(math.MaxUint32), ErrOverflow)
	testToSafeUint8(t, uint32(math.MaxUint8)+1, ErrOverflow)
	testToSafeUint8(t, uint32(math.MaxUint8), nil)
	testToSafeUint8(t, uint32(0), nil)

	// int16
	testToSafeUint8(t, int16(math.MaxInt16), ErrOverflow)
	testToSafeUint8(t, int16(math.MaxUint8)+1, ErrOverflow)
	testToSafeUint8(t, int16(math.MinInt16), ErrOverflow)
	testToSafeUint8(t, int16(-1), ErrOverflow)
	testToSafeUint8(t, int16(math.MaxUint8), nil)
	testToSafeUint8(t, int16(0), nil)

	// uint16
	testToSafeUint8(t, uint16(math.MaxUint16), ErrOverflow)
	testToSafeUint8(t, uint16(math.MaxUint8)+1, ErrOverflow)
	testToSafeUint8(t, uint16(math.MaxUint8), nil)
	testToSafeUint8(t, uint16(0), nil)

	// int8
	testToSafeUint8(t, int8(math.MinInt8), ErrOverflow)
	testToSafeUint8(t, int8(-1), ErrOverflow)
	testToSafeUint8(t, int8(math.MaxInt8), nil)
	testToSafeUint8(t, int8(0), nil)

	// uint8
	testToSafeUint8(t, uint8(math.MaxUint8), nil)
	testToSafeUint8(t, uint8(math.MaxInt8)+1, nil)
	testToSafeUint8(t, uint8(0), nil)
}

func testToSafeUint16[T IntExt | UIntExt](t *testing.T, v T, err error) {
	i16, e := ToSafeUint16(v)
	assert.Equal(t, e, err)
	assert.Equal(t, i16, uint16(v))
}

func Test_ToSafeUint16(t *testing.T) {
	// int64
	testToSafeUint16(t, int64(math.MaxInt64), ErrOverflow)
	testToSafeUint16(t, int64(math.MaxUint16)+1, ErrOverflow)
	testToSafeUint16(t, int64(math.MinInt64), ErrOverflow)
	testToSafeUint16(t, int64(-1), ErrOverflow)
	testToSafeUint16(t, int64(math.MaxUint16), nil)
	testToSafeUint16(t, int64(0), nil)

	// uint64
	testToSafeUint16(t, uint64(math.MaxUint64), ErrOverflow)
	testToSafeUint16(t, uint64(math.MaxUint16)+1, ErrOverflow)
	testToSafeUint16(t, uint64(math.MaxUint16), nil)
	testToSafeUint16(t, uint64(0), nil)

	// int32
	testToSafeUint16(t, int32(math.MaxInt32), ErrOverflow)
	testToSafeUint16(t, int32(math.MaxUint16)+1, ErrOverflow)
	testToSafeUint16(t, int32(math.MinInt32), ErrOverflow)
	testToSafeUint16(t, int32(-1), ErrOverflow)
	testToSafeUint16(t, int32(math.MaxUint16), nil)
	testToSafeUint16(t, int32(0), nil)

	// uint32
	testToSafeUint16(t, uint32(math.MaxUint32), ErrOverflow)
	testToSafeUint16(t, uint32(math.MaxUint16)+1, ErrOverflow)
	testToSafeUint16(t, uint32(math.MaxUint16), nil)
	testToSafeUint16(t, uint32(0), nil)

	// int16
	testToSafeUint16(t, int16(math.MinInt16), ErrOverflow)
	testToSafeUint16(t, int16(-1), ErrOverflow)
	testToSafeUint16(t, int16(math.MaxInt16), nil)
	testToSafeUint16(t, int16(0), nil)

	// uint16
	testToSafeUint16(t, uint16(math.MaxUint16), nil)
	testToSafeUint16(t, uint16(0), nil)

	// int8
	testToSafeUint16(t, int8(math.MinInt8), ErrOverflow)
	testToSafeUint16(t, int8(-1), ErrOverflow)
	testToSafeUint16(t, int8(math.MaxInt8), nil)
	testToSafeUint16(t, int8(0), nil)

	// uint8
	testToSafeUint16(t, uint8(math.MaxUint8), nil)
	testToSafeUint16(t, uint8(math.MaxInt8)+1, nil)
	testToSafeUint16(t, uint8(0), nil)
}

func testToSafeUint32[T IntExt | UIntExt](t *testing.T, v T, err error) {
	i32, e := ToSafeUint32(v)
	assert.Equal(t, e, err)
	assert.Equal(t, i32, uint32(v))
}

func Test_ToSafeUint32(t *testing.T) {
	// int64
	testToSafeUint32(t, int64(math.MaxInt64), ErrOverflow)
	testToSafeUint32(t, int64(math.MaxUint32)+1, ErrOverflow)
	testToSafeUint32(t, int64(math.MinInt64), ErrOverflow)
	testToSafeUint32(t, int64(-1), ErrOverflow)
	testToSafeUint32(t, int64(math.MaxUint32), nil)
	testToSafeUint32(t, int64(0), nil)

	// uint64
	testToSafeUint32(t, uint64(math.MaxUint64), ErrOverflow)
	testToSafeUint32(t, uint64(math.MaxUint32)+1, ErrOverflow)
	testToSafeUint32(t, uint64(math.MaxUint32), nil)
	testToSafeUint32(t, uint64(0), nil)

	// int32
	testToSafeUint32(t, int32(math.MinInt32), ErrOverflow)
	testToSafeUint32(t, int32(-1), ErrOverflow)
	testToSafeUint32(t, int32(math.MaxInt32), nil)
	testToSafeUint32(t, int32(0), nil)

	// uint32
	testToSafeUint32(t, uint32(math.MaxUint32), nil)
	testToSafeUint32(t, uint32(0), nil)

	// int16
	testToSafeUint32(t, int16(math.MinInt16), ErrOverflow)
	testToSafeUint32(t, int16(-1), ErrOverflow)
	testToSafeUint32(t, int16(math.MaxInt16), nil)
	testToSafeUint32(t, int16(0), nil)

	// uint16
	testToSafeUint32(t, uint16(math.MaxUint16), nil)
	testToSafeUint32(t, uint16(0), nil)

	// int8
	testToSafeUint32(t, int8(math.MinInt8), ErrOverflow)
	testToSafeUint32(t, int8(-1), ErrOverflow)
	testToSafeUint32(t, int8(math.MaxInt8), nil)
	testToSafeUint32(t, int8(0), nil)

	// uint8
	testToSafeUint32(t, uint8(math.MaxUint8), nil)
	testToSafeUint32(t, uint8(math.MaxInt8)+1, nil)
	testToSafeUint32(t, uint8(0), nil)
}

func testToSafeUint64[T IntExt | UIntExt](t *testing.T, v T, err error) {
	i64, e := ToSafeUint64(v)
	assert.Equal(t, e, err)
	assert.Equal(t, i64, uint64(v))
}

func Test_ToSafeUint64(t *testing.T) {
	// int64
	testToSafeUint64(t, int64(math.MinInt64), ErrOverflow)
	testToSafeUint64(t, int64(-1), ErrOverflow)
	testToSafeUint64(t, int64(math.MaxInt64), nil)
	testToSafeUint64(t, int64(0), nil)

	// uint64
	testToSafeUint64(t, uint64(math.MaxUint64), nil)
	testToSafeUint64(t, uint64(0), nil)

	// int32
	testToSafeUint64(t, int32(math.MinInt32), ErrOverflow)
	testToSafeUint64(t, int32(-1), ErrOverflow)
	testToSafeUint64(t, int32(math.MaxInt32), nil)
	testToSafeUint64(t, int32(0), nil)

	// uint32
	testToSafeUint64(t, uint32(math.MaxUint32), nil)
	testToSafeUint64(t, uint32(0), nil)

	// int16
	testToSafeUint64(t, int16(math.MinInt16), ErrOverflow)
	testToSafeUint64(t, int16(-1), ErrOverflow)
	testToSafeUint64(t, int16(math.MaxInt16), nil)
	testToSafeUint64(t, int16(0), nil)

	// uint16
	testToSafeUint64(t, uint16(math.MaxUint16), nil)
	testToSafeUint64(t, uint16(0), nil)

	// int8
	testToSafeUint64(t, int8(math.MinInt8), ErrOverflow)
	testToSafeUint64(t, int8(-1), ErrOverflow)
	testToSafeUint64(t, int8(math.MaxInt8), nil)
	testToSafeUint64(t, int8(0), nil)

	// uint8
	testToSafeUint64(t, uint8(math.MaxUint8), nil)
	testToSafeUint64(t, uint8(math.MaxInt8)+1, nil)
	testToSafeUint64(t, uint8(0), nil)
}
