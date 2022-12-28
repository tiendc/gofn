package gofn

// Flatten flattens 2-dimensional slices
// E.g. Flatten([1,2,3], [3,4,5]) -> [1,2,3,3,4,5]
func Flatten[T any](s ...[]T) []T {
	result := make([]T, 0, len(s)*5) // just guess the capacity
	for _, innerSlice := range s {
		result = append(result, innerSlice...)
	}
	return result
}

// Flatten3 flattens 3-dimensional slices
func Flatten3[T any](s ...[][]T) []T {
	result := make([]T, 0, len(s)*30) // just guess the capacity
	for _, innerSlice := range s {
		for _, mostInnerSlice := range innerSlice {
			result = append(result, mostInnerSlice...)
		}
	}
	return result
}
