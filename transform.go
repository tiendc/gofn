package gofn

// MapSlice transforms a slice to another with map function
func MapSlice[T any, U any](s []T, mapFunc func(b T) U) []U {
	result := make([]U, len(s))
	for i := range s {
		result[i] = mapFunc(s[i])
	}
	return result
}

// MapSliceEx transforms a slice to another with map function and error handling
func MapSliceEx[T any, U any](s []T, mapFunc func(b T) (U, error)) ([]U, error) {
	result := make([]U, 0, len(s))
	for i := range s {
		v, err := mapFunc(s[i])
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
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

// MapSliceToMapEx transforms a slice to a map with map function and error handling
func MapSliceToMapEx[T any, K comparable, V any](s []T, mapFunc func(b T) (K, V, error)) (map[K]V, error) {
	result := make(map[K]V, len(s))
	for i := range s {
		k, v, err := mapFunc(s[i])
		if err != nil {
			return nil, err
		}
		result[k] = v
	}
	return result, nil
}
