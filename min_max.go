package gofn

// Min find the minimum value in the list
func Min[T NumberEx | ~string](v1 T, s ...T) T {
	min := v1
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

// MinIn find the minimum value in the list
// Use min := Must(MinIn(slice)) to panic on error
func MinIn[T NumberEx | ~string, S ~[]T](s S) (T, error) {
	if len(s) == 0 {
		var zeroT T
		return zeroT, ErrEmpty
	}
	min := s[0]
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}
	return min, nil
}

// MinInPred find the minimum value in the list
func MinInPred[T any, S ~[]T](s S, lessFunc func(a, b T) bool) (T, error) {
	if len(s) == 0 {
		var zeroT T
		return zeroT, ErrEmpty
	}
	min := s[0]
	for i := range s {
		if lessFunc(s[i], min) {
			min = s[i]
		}
	}
	return min, nil
}

// Max find the maximum value in the list
func Max[T NumberEx | ~string](v1 T, s ...T) T {
	max := v1
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

// MaxIn finds the maximum value in the list
// Use max := Must(MaxIn(slice)) to panic on error
func MaxIn[T NumberEx | ~string, S ~[]T](s S) (T, error) {
	if len(s) == 0 {
		var zeroT T
		return zeroT, ErrEmpty
	}
	max := s[0]
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return max, nil
}

// MaxInPred finds the maximum value in the list
func MaxInPred[T any, S ~[]T](s S, lessFunc func(a, b T) bool) (T, error) {
	if len(s) == 0 {
		var zeroT T
		return zeroT, ErrEmpty
	}
	max := s[0]
	for i := range s {
		if lessFunc(max, s[i]) {
			max = s[i]
		}
	}
	return max, nil
}

// MinMax finds the minimum and maximum values in the list
func MinMax[T NumberEx | ~string](v1 T, s ...T) (T, T) {
	min := v1
	max := v1
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
		if s[i] > max {
			max = s[i]
		}
	}
	return min, max
}
