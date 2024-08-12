package gofn

// Reduce reduces a slice to a value
func Reduce[T any, S ~[]T](s S, reduceFunc func(accumulator, currentValue T) T) T {
	accumulator := s[0]
	for i := 1; i < len(s); i++ {
		accumulator = reduceFunc(accumulator, s[i])
	}
	return accumulator
}

// ReduceEx reduces a slice to a value with custom initial value
func ReduceEx[T any, U any, S ~[]T](
	s S,
	reduceFunc func(accumulator U, currentValue T, currentIndex int) U,
	initVal U,
) U {
	accumulator := initVal
	for i, v := range s {
		accumulator = reduceFunc(accumulator, v, i)
	}
	return accumulator
}
