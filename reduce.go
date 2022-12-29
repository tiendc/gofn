package gofn

func Reduce[T any](s []T, reduceFunc func(accumulator, currentValue T) T) T {
	accumulator := s[0]
	for i := 1; i < len(s); i++ {
		accumulator = reduceFunc(accumulator, s[i])
	}
	return accumulator
}

func ReduceEx[T any, U any](s []T, reduceFunc func(accumulator U, currentValue T, currentIndex int) U, initVal U) U {
	accumulator := initVal
	for i, v := range s {
		accumulator = reduceFunc(accumulator, v, i)
	}
	return accumulator
}
