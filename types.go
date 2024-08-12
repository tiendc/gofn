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
	uint | uint8 | uint16 | uint32 | uint64
}

type UIntExt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type UIntPtr interface {
	*uint | *uint8 | *uint16 | *uint32 | *uint64
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

type Tuple2[T any, U any] struct {
	Elem1 T
	Elem2 U
}

type Tuple3[T any, U any, V any] struct {
	Elem1 T
	Elem2 U
	Elem3 V
}
