package gofn

// All returns true if all given values are evaluated as `true`
func All[T comparable](s ...T) bool {
	var zeroT T
	for i := range s {
		if s[i] == zeroT {
			return false
		}
	}
	return true
}

// Any returns true if at least one of given values is evaluated as `true`
func Any[T comparable](s ...T) bool {
	var zeroT T
	for i := range s {
		if s[i] != zeroT {
			return true
		}
	}
	return false
}

// Abs calculates absolute value of an integer.
// Ref: http://cavaliercoder.com/blog/optimized-abs-for-int64-in-go.html.
// Note: if inputs MinInt64, the result is negative.
func Abs(n int64) int64 {
	y := n >> 63 //nolint:mnd
	return (n ^ y) - y
}

// Clamp clamps number within the inclusive lower and upper bounds.
func Clamp[T NumberExt | StringExt](value, min, max T) T {
	if min > max {
		min, max = max, min
	}
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Product calculates product value of slice elements
func Product[T NumberExt | ComplexExt](s ...T) T {
	if len(s) == 0 {
		return 0
	}
	product := T(1)
	for _, v := range s {
		product *= v
	}
	return product
}

// ProductAs calculates product value with conversion to another type.
// Type size of the result should be wider than the input's.
// E.g. product := ProductAs[int64](int32Slice...)
func ProductAs[U, T NumberExt](s ...T) U {
	if len(s) == 0 {
		return 0
	}
	product := U(1)
	for _, v := range s {
		product *= U(v)
	}
	return product
}

// Sum calculates sum of slice items
func Sum[T NumberExt | ComplexExt](s ...T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// SumAs calculates sum value with conversion to another type.
// Type size of the result should be wider than the input's.
// E.g. sum := SumAs[int64](int32Slice...)
func SumAs[U, T NumberExt](s ...T) U {
	var sum U
	for _, v := range s {
		sum += U(v)
	}
	return sum
}

// Min find the minimum value in the list
func Min[T NumberExt | StringExt](v1 T, s ...T) T {
	min := v1
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

// MinIn find the minimum value in the list.
// Use min := Must(MinIn(slice)) to panic on error.
func MinIn[T NumberExt | StringExt, S ~[]T](s S) (T, error) {
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

// MinInBy find the minimum value in the list
func MinInBy[T any, S ~[]T](s S, lessFunc func(a, b T) bool) (T, error) {
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
func Max[T NumberExt | StringExt](v1 T, s ...T) T {
	max := v1
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

// MaxIn finds the maximum value in the list.
// Use max := Must(MaxIn(slice)) to panic on error.
func MaxIn[T NumberExt | StringExt, S ~[]T](s S) (T, error) {
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

// MaxInBy finds the maximum value in the list
func MaxInBy[T any, S ~[]T](s S, lessFunc func(a, b T) bool) (T, error) {
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
func MinMax[T NumberExt | StringExt](v1 T, s ...T) (T, T) {
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
