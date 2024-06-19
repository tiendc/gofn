package gofn

type IntEx interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Int interface {
	int | int8 | int16 | int32 | int64
}

type IntPtr interface {
	*int | *int8 | *int16 | *int32 | *int64
}

type UIntEx interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type UIntPtr interface {
	*uint | *uint8 | *uint16 | *uint32 | *uint64
}

type FloatEx interface {
	~float32 | ~float64
}

type Float interface {
	float32 | float64
}

type FloatPtr interface {
	*float32 | *float64
}

type NumberEx interface {
	IntEx | UIntEx | FloatEx
}

type Number interface {
	Int | UInt | Float
}

type NumberPtr interface {
	IntPtr | UIntPtr | FloatPtr
}

type StringEx interface {
	~string
}

type StringPtr interface {
	*string
}

type Tuple2[T any, U any] struct {
	Elem1 T
	Elem2 U
}

type Tuple3[T any, U any, V any] struct {
	Elem1 T
	Elem2 U
	Elem3 V
}
