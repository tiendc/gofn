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
