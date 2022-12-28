package gofn

func Zip[T any, U any](slice1 []T, slice2 []U) []*Tuple2[T, U] {
	minLen := len(slice1)
	if minLen > len(slice2) {
		minLen = len(slice2)
	}
	result := make([]*Tuple2[T, U], minLen)
	for i := 0; i < minLen; i++ {
		result[i] = &Tuple2[T, U]{slice1[i], slice2[i]}
	}
	return result
}

func Zip3[T any, U any, V any](slice1 []T, slice2 []U, slice3 []V) []*Tuple3[T, U, V] {
	minLen := len(slice1)
	if minLen > len(slice2) {
		minLen = len(slice2)
	}
	if minLen > len(slice3) {
		minLen = len(slice3)
	}
	result := make([]*Tuple3[T, U, V], minLen)
	for i := 0; i < minLen; i++ {
		result[i] = &Tuple3[T, U, V]{slice1[i], slice2[i], slice3[i]}
	}
	return result
}
