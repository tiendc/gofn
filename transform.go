package gofn

// MapSlice transforms a slice to another with map function
func MapSlice[T any, U any, S ~[]T](s S, mapFunc func(T) U) []U {
	result := make([]U, len(s))
	for i := range s {
		result[i] = mapFunc(s[i])
	}
	return result
}

// MapSliceEx transforms a slice to another with map function and error handling
func MapSliceEx[T any, U any, S ~[]T](s S, mapFunc func(T) (U, error)) ([]U, error) {
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
func MapSliceToMap[T any, K comparable, V any, S ~[]T](s S, mapFunc func(T) (K, V)) map[K]V {
	result := make(map[K]V, len(s))
	for i := range s {
		k, v := mapFunc(s[i])
		result[k] = v
	}
	return result
}

// MapSliceToMapEx transforms a slice to a map with map function and error handling
func MapSliceToMapEx[T any, K comparable, V any, S ~[]T](s S, mapFunc func(T) (K, V, error)) (map[K]V, error) {
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

// MapSliceToMapKeys transforms a slice to a map using slice items as map keys.
// For example: MapSliceToMapKeys(s, struct{}{}) -> map[T]struct{}
func MapSliceToMapKeys[T comparable, V any, S ~[]T](s S, defaultVal V) map[T]V {
	if len(s) == 0 {
		return map[T]V{}
	}
	ret := make(map[T]V, len(s))
	for _, v := range s {
		ret[v] = defaultVal
	}
	return ret
}
