package gofn

// ToSet calculates unique values of a slice
func ToSet[T comparable, S ~[]T](s S) S {
	length := len(s)
	if length <= 1 {
		return append(S{}, s...)
	}

	seen := make(map[T]struct{}, length)
	result := make(S, 0, length)

	for i := 0; i < length; i++ {
		v := s[i]
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		result = append(result, v)
	}

	return result
}

// ToSetBy calculates unique values of a slice with custom key function
func ToSetBy[T any, K comparable, S ~[]T](s S, keyFunc func(t T) K) S {
	length := len(s)
	if length <= 1 {
		return append(S{}, s...)
	}

	seen := make(map[K]struct{}, length)
	result := make(S, 0, length)

	for i := 0; i < length; i++ {
		v := s[i]
		k := keyFunc(v)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		result = append(result, v)
	}

	return result
}

// Deprecated: use ToSetBy instead
func ToSetPred[T any, K comparable, S ~[]T](s S, keyFunc func(t T) K) S {
	return ToSetBy(s, keyFunc)
}

// ToSetByReverse calculates unique values of a slice with custom key function.
// Unlike ToSetBy, this function iterates over the slice from the end.
func ToSetByReverse[T any, K comparable, S ~[]T](s S, keyFunc func(t T) K) S {
	length := len(s)
	if length <= 1 {
		return append(S{}, s...)
	}

	seen := make(map[K]struct{}, length)
	result := make(S, 0, length)

	for i := length - 1; i >= 0; i-- {
		v := s[i]
		k := keyFunc(v)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		result = append(result, v)
	}

	return result
}

// Deprecated: use ToSetByReverse instead
func ToSetPredReverse[T any, K comparable, S ~[]T](s S, keyFunc func(t T) K) S {
	return ToSetByReverse(s, keyFunc)
}

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
