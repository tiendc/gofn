package gofn

// MapSlice transforms a slice to another with map function
func MapSlice[T any, U any](s []T, mapFunc func(b T) U) []U {
	result := make([]U, len(s))
	for i := range s {
		result[i] = mapFunc(s[i])
	}
	return result
}

// MapSliceEx supports map and filter functionality
func MapSliceEx[T any, U any](s []T, mapFunc func(b T) (U, bool)) []U {
	result := make([]U, 0, len(s))
	for i := range s {
		item, ok := mapFunc(s[i])
		if ok {
			result = append(result, item)
		}
	}
	return result
}

// MapSliceToMap transforms a slice to a map with map function
func MapSliceToMap[T any, K comparable, V any](s []T, mapFunc func(b T) (K, V)) map[K]V {
	result := make(map[K]V, len(s))
	for i := range s {
		k, v := mapFunc(s[i])
		result[k] = v
	}
	return result
}

// MapSliceToMapEx supports map and filter functionality
func MapSliceToMapEx[T any, K comparable, V any](s []T, mapFunc func(b T) (K, V, bool)) map[K]V {
	result := make(map[K]V, len(s))
	for i := range s {
		k, v, ok := mapFunc(s[i])
		if ok {
			result[k] = v
		}
	}
	return result
}
