package gofn

func ToSet[T comparable](s []T) []T {
	length := len(s)
	if length <= 1 {
		return ToSlice(s...)
	}

	seen := make(map[T]struct{}, length)
	result := make([]T, 0, length)

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

func ToSetPred[T any, K comparable](s []T, keyFunc func(t T) K) []T {
	length := len(s)
	if length <= 1 {
		return ToSlice(s...)
	}

	seen := make(map[K]struct{}, length)
	result := make([]T, 0, length)

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
