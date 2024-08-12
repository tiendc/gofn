package gofn

// Zip combines values from 2 slices by each position
func Zip[T any, U any, S1 ~[]T, S2 ~[]U](slice1 S1, slice2 S2) []*Tuple2[T, U] {
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

// Zip3 combines values from 3 slices by each position
func Zip3[T any, U any, V any, S1 ~[]T, S2 ~[]U, S3 ~[]V](slice1 S1, slice2 S2, slice3 S3) []*Tuple3[T, U, V] {
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
