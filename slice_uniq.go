package gofn

// IsUnique checks a slice for uniqueness
func IsUnique[T comparable, S ~[]T](s S) bool {
	length := len(s)
	if length <= 1 {
		return true
	}
	seen := make(map[T]struct{}, length)
	for i := 0; i < length; i++ {
		v := s[i]
		if _, ok := seen[v]; ok {
			return false
		}
		seen[v] = struct{}{}
	}
	return true
}

// IsUniqueBy checks a slice for uniqueness using key function
func IsUniqueBy[T any, U comparable, S ~[]T](s S, keyFunc func(t T) U) bool {
	length := len(s)
	if length <= 1 {
		return true
	}
	seen := make(map[U]struct{}, length)
	for i := 0; i < length; i++ {
		v := keyFunc(s[i])
		if _, ok := seen[v]; ok {
			return false
		}
		seen[v] = struct{}{}
	}
	return true
}

// FindUniques finds all elements that are unique in the given slice.
func FindUniques[T comparable, S ~[]T](s S) S {
	return FindUniquesBy(s, func(t T) T { return t })
}

// FindUniquesBy finds all elements that are unique in the given slice.
func FindUniquesBy[T any, U comparable, S ~[]T](s S, keyFunc func(T) U) S {
	length := len(s)
	if length <= 1 {
		return append(S{}, s...)
	}

	seen := make(map[U]int, length)  // Map to store elements and their first indexes
	dupFlags := make([]bool, length) // Array to store flags of duplication of elements
	uniqCount := length

	for i := range s {
		k := keyFunc(s[i])
		if firstIndex, ok := seen[k]; ok {
			dupFlags[firstIndex] = true
			dupFlags[i] = true
			uniqCount--
			continue
		}
		seen[k] = i
	}

	if uniqCount == length {
		return append(S{}, s...)
	}
	result := make(S, 0, uniqCount)
	for i := 0; i < length; i++ {
		if !dupFlags[i] {
			result = append(result, s[i])
		}
	}
	return result
}

// FindDuplicates finds all elements that are duplicated in the given slice.
func FindDuplicates[T comparable, S ~[]T](s S) S {
	return FindDuplicatesBy(s, func(t T) T { return t })
}

// FindDuplicatesBy finds all elements that are duplicated in the given slice.
func FindDuplicatesBy[T any, U comparable, S ~[]T](s S, keyFunc func(T) U) S {
	length := len(s)
	if length <= 1 {
		return S{}
	}

	seen := make(map[U]int, length)  // Map to store elements and their first indexes
	dupFlags := make([]bool, length) // Array to store flags of duplication of elements
	dupCount := 0

	for i := range s {
		k := keyFunc(s[i])
		if firstIndex, ok := seen[k]; ok {
			dupFlags[firstIndex] = true
			dupCount++
			continue
		}
		seen[k] = i
	}

	if dupCount == 0 {
		return S{}
	}
	result := make(S, 0, dupCount)
	for i := 0; i < length; i++ {
		if dupFlags[i] {
			result = append(result, s[i])
		}
	}
	return result
}

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
