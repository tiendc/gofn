package gofn

// Union returns all unique values from multiple slices
func Union[T comparable, S ~[]T](a, b S) S {
	lenA, lenB := len(a), len(b)
	if lenA == 0 {
		return ToSet(b)
	}
	if lenB == 0 {
		return ToSet(a)
	}

	seen := make(map[T]struct{}, lenA+lenB)
	result := make(S, 0, lenA+lenB)

	for i := 0; i < lenA; i++ {
		v := a[i]
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		result = append(result, v)
	}
	for i := 0; i < lenB; i++ {
		v := b[i]
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		result = append(result, v)
	}

	return result
}

// UnionPred returns all unique values from multiple slices with key function
func UnionPred[T any, K comparable, S ~[]T](a, b S, keyFunc func(t T) K) S {
	lenA, lenB := len(a), len(b)
	if lenA == 0 {
		return ToSetPred(b, keyFunc)
	}
	if lenB == 0 {
		return ToSetPred(a, keyFunc)
	}

	seen := make(map[K]struct{}, lenA+lenB)
	result := make(S, 0, lenA+lenB)

	for i := 0; i < lenA; i++ {
		v := a[i]
		k := keyFunc(v)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		result = append(result, v)
	}
	for i := 0; i < lenB; i++ {
		v := b[i]
		k := keyFunc(v)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		result = append(result, v)
	}

	return result
}

// Intersection returns all unique shared values from multiple slices
func Intersection[T comparable, S ~[]T](a, b S) S {
	lenA, lenB := len(a), len(b)
	if lenA == 0 || lenB == 0 {
		return S{}
	}

	seen := make(map[T]*bool, lenA)
	result := make(S, 0, Min(lenA, lenB))

	for i := 0; i < lenA; i++ {
		active := true
		seen[a[i]] = &active
	}
	for i := 0; i < lenB; i++ {
		v := b[i]
		if active, ok := seen[v]; ok && *active {
			result = append(result, v)
			*active = false
		}
	}

	return result
}

// IntersectionPred returns all unique shared values from multiple slices with key function
func IntersectionPred[T any, K comparable, S ~[]T](a, b S, keyFunc func(t T) K) S {
	lenA, lenB := len(a), len(b)
	if lenA == 0 || lenB == 0 {
		return S{}
	}

	seen := make(map[K]*bool, lenA)
	result := make(S, 0, Min(lenA, lenB))

	for i := 0; i < lenA; i++ {
		active := true
		seen[keyFunc(a[i])] = &active
	}
	for i := 0; i < lenB; i++ {
		v := b[i]
		k := keyFunc(v)
		if active, ok := seen[k]; ok && *active {
			result = append(result, v)
			*active = false
		}
	}

	return result
}

// Difference calculates the differences between two slices.
// The first result list contains all the values exist in the left list, but the right.
// The second result list contains all the values exist in the right list, but the left.
// NOTE: this function does not return unique values.
func Difference[T comparable, S ~[]T](a, b S) (S, S) {
	lenA, lenB := len(a), len(b)
	if lenA == 0 || lenB == 0 {
		return append(S{}, a...), append(S{}, b...)
	}

	leftDiff, rightDiff := S{}, S{}
	leftMap, rightMap := MapSliceToMapKeys(a, struct{}{}), MapSliceToMapKeys(b, struct{}{})

	for _, v := range a {
		if _, ok := rightMap[v]; !ok {
			leftDiff = append(leftDiff, v)
		}
	}
	for _, v := range b {
		if _, ok := leftMap[v]; !ok {
			rightDiff = append(rightDiff, v)
		}
	}

	return leftDiff, rightDiff
}

// DifferencePred calculates the differences between two slices using special key function.
// NOTE: this function does not return unique values.
func DifferencePred[T any, K comparable, S ~[]T](a, b S, keyFunc func(t T) K) (S, S) {
	lenA, lenB := len(a), len(b)
	if lenA == 0 || lenB == 0 {
		return append(S{}, a...), append(S{}, b...)
	}

	leftDiff, rightDiff := S{}, S{}
	leftMap, rightMap := make(map[K]struct{}, lenA), make(map[K]struct{}, lenB)

	for _, v := range a {
		leftMap[keyFunc(v)] = struct{}{}
	}
	for _, v := range b {
		rightMap[keyFunc(v)] = struct{}{}
	}

	for _, v := range a {
		if _, ok := rightMap[keyFunc(v)]; !ok {
			leftDiff = append(leftDiff, v)
		}
	}
	for _, v := range b {
		if _, ok := leftMap[keyFunc(v)]; !ok {
			rightDiff = append(rightDiff, v)
		}
	}

	return leftDiff, rightDiff
}
