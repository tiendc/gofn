package gofn

// Sum calculates sum of slice items
func Sum[T IntEx | UIntEx | FloatEx](s ...T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// SumAs calculates sum value with conversion to another type.
// Type size of the result should be wider than the input's.
// E.g. sum := SumAs[int64](int32Slice...)
func SumAs[U, T IntEx | UIntEx | FloatEx](s ...T) U {
	var sum U
	for _, v := range s {
		sum += U(v)
	}
	return sum
}
