package gofn

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

func ToSetPred[T any, K comparable, S ~[]T](s S, keyFunc func(t T) K) S {
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

func ToSetPredReverse[T any, K comparable, S ~[]T](s S, keyFunc func(t T) K) S {
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
