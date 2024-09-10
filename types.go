package gofn

type Int interface {
	int | int8 | int16 | int32 | int64
}

type IntExt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type IntPtr interface {
	*int | *int8 | *int16 | *int32 | *int64
}

type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type UIntExt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type UIntPtr interface {
	*uint | *uint8 | *uint16 | *uint32 | *uint64 | *uintptr
}

type Float interface {
	float32 | float64
}

type FloatExt interface {
	~float32 | ~float64
}

type FloatPtr interface {
	*float32 | *float64
}

type Complex interface {
	complex64 | complex128
}

type ComplexExt interface {
	~complex64 | ~complex128
}

type ComplexPtr interface {
	*complex64 | *complex128
}

type Number interface {
	Int | UInt | Float
}

type NumberExt interface {
	IntExt | UIntExt | FloatExt
}

type NumberPtr interface {
	IntPtr | UIntPtr | FloatPtr
}

type String interface {
	string
}

type StringExt interface {
	~string
}

type StringPtr interface {
	*string
}

// Deprecated: Use IntExt
type IntEx IntExt

// Deprecated: Use UIntExt
type UIntEx UIntExt

// Deprecated: Use FloatExt
type FloatEx FloatExt

// Deprecated: Use NumberExt
type NumberEx NumberExt

// Deprecated: Use StringExt
type StringEx StringExt
