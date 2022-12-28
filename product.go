package gofn

func Product[T IntEx | UIntEx | FloatEx](s ...T) T {
	if len(s) == 0 {
		return 0
	}
	product := T(1)
	for _, v := range s {
		product *= v
	}
	return product
}

// ProductAs calculates product value with conversion to another type.
// Type size of the result should be wider than the input's
// E.g. product := ProductAs[int64](int32Slice...)
func ProductAs[U, T IntEx | UIntEx | FloatEx](s ...T) U {
	if len(s) == 0 {
		return 0
	}
	product := U(1)
	for _, v := range s {
		product *= U(v)
	}
	return product
}
