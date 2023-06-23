package gofn

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseInt(t *testing.T) {
	i8, e := ParseInt[int8]("127")
	assert.True(t, e == nil && i8 == int8(127))
	i8, e = ParseInt[int8]("-128")
	assert.True(t, e == nil && i8 == int8(-128))
	_, e = ParseInt[int8]("128")
	assert.ErrorIs(t, e, strconv.ErrRange)
	_, e = ParseInt[int8]("-129")
	assert.ErrorIs(t, e, strconv.ErrRange)

	i16, e := ParseInt[int16]("32767")
	assert.True(t, e == nil && i16 == int16(32767))
	i16, e = ParseInt[int16]("-32768")
	assert.True(t, e == nil && i16 == int16(-32768))
	_, e = ParseInt[int16]("32768")
	assert.ErrorIs(t, e, strconv.ErrRange)
	_, e = ParseInt[int16]("-32769")
	assert.ErrorIs(t, e, strconv.ErrRange)

	i32, e := ParseInt[int32]("2147483647")
	assert.True(t, e == nil && i32 == int32(2147483647))
	i32, e = ParseInt[int32]("-2147483648")
	assert.True(t, e == nil && i32 == int32(-2147483648))
	_, e = ParseInt[int32]("2147483648")
	assert.ErrorIs(t, e, strconv.ErrRange)
	_, e = ParseInt[int32]("-2147483649")
	assert.ErrorIs(t, e, strconv.ErrRange)

	i64, e := ParseInt[int64]("9223372036854775807")
	assert.True(t, e == nil && i64 == int64(9223372036854775807))
	i64, e = ParseInt[int64]("-9223372036854775808")
	assert.True(t, e == nil && i64 == int64(-9223372036854775808))
	_, e = ParseInt[int64]("9223372036854775808")
	assert.ErrorIs(t, e, strconv.ErrRange)
	_, e = ParseInt[int64]("-9223372036854775809")
	assert.ErrorIs(t, e, strconv.ErrRange)

	// Custom type
	type X8 int8
	x8, e := ParseInt[X8]("127")
	assert.True(t, e == nil && x8 == X8(127))
	x8, e = ParseInt[X8]("-128")
	assert.True(t, e == nil && x8 == X8(-128))
	_, e = ParseInt[X8]("128")
	assert.ErrorIs(t, e, strconv.ErrRange)
	_, e = ParseInt[X8]("-129")
	assert.ErrorIs(t, e, strconv.ErrRange)
}

func Test_ParseIntDef(t *testing.T) {
	i8 := ParseIntDef("127", int8(120))
	assert.True(t, i8 == int8(127))
	i8 = ParseIntDef("-129", int8(100)) // overflow, return default
	assert.True(t, i8 == int8(100))

	i16 := ParseIntDef("-32768", int16(100))
	assert.True(t, i16 == int16(-32768))
	i16 = ParseIntDef("32768", int16(100)) // overflow, return default
	assert.True(t, i16 == int16(100))

	i32 := ParseIntDef("2147483647", int32(100))
	assert.True(t, i32 == int32(2147483647))
	i32 = ParseIntDef("-2147483649", int32(-100)) // overflow, return default
	assert.True(t, i32 == int32(-100))

	i64 := ParseIntDef("-9223372036854775808", int64(100))
	assert.True(t, i64 == int64(-9223372036854775808))
	i64 = ParseIntDef("9223372036854775808", int64(-100)) // overflow, return default
	assert.True(t, i64 == int64(-100))

	// Custom type
	type X32 int32
	x32 := ParseIntDef("2147483647", X32(100))
	assert.True(t, x32 == X32(2147483647))
	x32 = ParseIntDef("-2147483649", X32(-100)) // overflow, return default
	assert.True(t, x32 == X32(-100))
}

func Test_ParseUint(t *testing.T) {
	i8, e := ParseUint[uint8]("255")
	assert.True(t, e == nil && i8 == uint8(255))
	i8, e = ParseUint[uint8]("0")
	assert.True(t, e == nil && i8 == uint8(0))
	_, e = ParseUint[uint8]("256")
	assert.ErrorIs(t, e, strconv.ErrRange)
	_, e = ParseUint[uint8]("-1")
	assert.ErrorIs(t, e, strconv.ErrSyntax)

	i16, e := ParseUint[uint16]("65535")
	assert.True(t, e == nil && i16 == uint16(65535))
	i16, e = ParseUint[uint16]("0")
	assert.True(t, e == nil && i16 == uint16(0))
	_, e = ParseUint[uint16]("65536")
	assert.ErrorIs(t, e, strconv.ErrRange)
	_, e = ParseUint[uint16]("-1")
	assert.ErrorIs(t, e, strconv.ErrSyntax)

	i32, e := ParseUint[uint32]("4294967295")
	assert.True(t, e == nil && i32 == uint32(4294967295))
	i32, e = ParseUint[uint32]("0")
	assert.True(t, e == nil && i32 == uint32(0))
	_, e = ParseUint[uint32]("4294967296")
	assert.ErrorIs(t, e, strconv.ErrRange)
	_, e = ParseUint[uint32]("-1")
	assert.ErrorIs(t, e, strconv.ErrSyntax)

	i64, e := ParseUint[uint64]("18446744073709551615")
	assert.True(t, e == nil && i64 == uint64(18446744073709551615))
	i64, e = ParseUint[uint64]("0")
	assert.True(t, e == nil && i64 == uint64(0))
	_, e = ParseUint[uint64]("18446744073709551616")
	assert.ErrorIs(t, e, strconv.ErrRange)
	_, e = ParseUint[uint64]("-1")
	assert.ErrorIs(t, e, strconv.ErrSyntax)

	// Custom type
	type X32 uint32
	x32, e := ParseUint[X32]("4294967295")
	assert.True(t, e == nil && x32 == X32(4294967295))
	x32, e = ParseUint[X32]("0")
	assert.True(t, e == nil && x32 == X32(0))
	_, e = ParseUint[X32]("4294967296")
	assert.ErrorIs(t, e, strconv.ErrRange)
	_, e = ParseUint[X32]("-1")
	assert.ErrorIs(t, e, strconv.ErrSyntax)
}

func Test_ParseUintDef(t *testing.T) {
	i8 := ParseUintDef("255", uint8(120))
	assert.True(t, i8 == uint8(255))
	i8 = ParseUintDef("-1", uint8(100)) // overflow, return default
	assert.True(t, i8 == uint8(100))

	i16 := ParseUintDef("65535", uint16(100))
	assert.True(t, i16 == uint16(65535))
	i16 = ParseUintDef("65536", uint16(100)) // overflow, return default
	assert.True(t, i16 == uint16(100))

	i32 := ParseUintDef("4294967295", uint32(100))
	assert.True(t, i32 == uint32(4294967295))
	i32 = ParseUintDef("4294967296", uint32(100)) // overflow, return default
	assert.True(t, i32 == uint32(100))

	i64 := ParseUintDef("18446744073709551615", uint64(100))
	assert.True(t, i64 == uint64(18446744073709551615))
	i64 = ParseUintDef("18446744073709551616", uint64(100)) // overflow, return default
	assert.True(t, i64 == uint64(100))

	// Custom type
	type X64 uint64
	x64 := ParseUintDef("18446744073709551615", X64(100))
	assert.True(t, x64 == X64(18446744073709551615))
	x64 = ParseUintDef("18446744073709551616", X64(100)) // overflow, return default
	assert.True(t, x64 == X64(100))
}

func Test_ParseFloat(t *testing.T) {
	f32, e := ParseFloat[float32]("123456.123456")
	assert.True(t, e == nil && f32 == float32(123456.123456))
	f32, e = ParseFloat[float32]("3.40282346638528859811704183484516925440e+38")
	assert.True(t, e == nil && f32 == math.MaxFloat32)
	f32, e = ParseFloat[float32]("123456.1234567") // out of precision
	assert.True(t, e == nil && f32 == float32(123456.123456))
	_, e = ParseFloat[float32]("4.40282346638528859811704183484516925440e+38") // overflow
	assert.ErrorIs(t, e, strconv.ErrRange)

	f64, e := ParseFloat[float64]("123456789.123456789")
	assert.True(t, e == nil && f64 == float64(123456789.123456789))
	f64, e = ParseFloat[float64]("1.79769313486231570814527423731704356798070e+308")
	assert.True(t, e == nil && f64 == math.MaxFloat64)
	_, e = ParseFloat[float64]("2.79769313486231570814527423731704356798070e+308") // overflow
	assert.ErrorIs(t, e, strconv.ErrRange)
}

func Test_ParseFloatDef(t *testing.T) {
	f32 := ParseFloatDef("123456.123456", float32(1))
	assert.True(t, f32 == float32(123456.123456))
	f32 = ParseFloatDef("3.40282346638528859811704183484516925440e+38", float32(0))
	assert.True(t, f32 == math.MaxFloat32)
	f32 = ParseFloatDef("123456.1234567", float32(100)) // out of precision
	assert.True(t, f32 == float32(123456.123456))
	f32 = ParseFloatDef("4.40282346638528859811704183484516925440e+38", float32(1)) // overflow
	assert.True(t, f32 == float32(1))

	f64 := ParseFloatDef("123456789.123456789", float64(1))
	assert.True(t, f64 == float64(123456789.123456789))
	f64 = ParseFloatDef("1.79769313486231570814527423731704356798070e+308", float64(100))
	assert.True(t, f64 == math.MaxFloat64)
	f64 = ParseFloatDef("2.79769313486231570814527423731704356798070e+308", float64(1.2)) // overflow
	assert.True(t, f64 == float64(1.2))
}
