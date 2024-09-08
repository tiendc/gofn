package gofn

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

// Deprecated: use MinInBy
func MinInPred[T any, S ~[]T](s S, lessFunc func(a, b T) bool) (T, error) {
	return MinInBy(s, lessFunc)
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

// Deprecated: use MaxInBy
func MaxInPred[T any, S ~[]T](s S, lessFunc func(a, b T) bool) (T, error) {
	return MaxInBy(s, lessFunc)
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

// Reduce reduces a slice to a value
func Reduce[T any, S ~[]T](s S, reduceFunc func(accumulator, currentValue T) T) T {
	accumulator := s[0]
	for i := 1; i < len(s); i++ {
		accumulator = reduceFunc(accumulator, s[i])
	}
	return accumulator
}

// ReduceEx reduces a slice to a value with custom initial value
func ReduceEx[T any, U any, S ~[]T](
	s S,
	reduceFunc func(accumulator U, currentValue T, currentIndex int) U,
	initVal U,
) U {
	accumulator := initVal
	for i, v := range s {
		accumulator = reduceFunc(accumulator, v, i)
	}
	return accumulator
}

// Partition splits slice items into 2 lists.
func Partition[T any, S ~[]T](s S, partitionFunc func(T, int) bool) (S, S) {
	partition0, partitionRemaining := S{}, S{}
	for i, v := range s {
		if partitionFunc(v, i) {
			partition0 = append(partition0, v)
		} else {
			partitionRemaining = append(partitionRemaining, v)
		}
	}
	return partition0, partitionRemaining
}

// PartitionN splits slice items into N lists.
// partitionFunc should return index of the partition to put the corresponding item into.
func PartitionN[T any, S ~[]T](s S, numPartitions uint, partitionFunc func(T, int) int) []S {
	if numPartitions <= 0 {
		return []S{}
	}
	partitions := make([]S, numPartitions)
	for i := range partitions {
		partitions[i] = S{}
	}
	lastIndex := int(numPartitions) - 1 //nolint:gosec
	for i, v := range s {
		pIndex := partitionFunc(v, i)
		if pIndex < 0 || pIndex > lastIndex {
			pIndex = lastIndex
		}
		partitions[pIndex] = append(partitions[pIndex], v)
	}
	return partitions
}

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

// UnionBy returns all unique values from multiple slices with key function
func UnionBy[T any, K comparable, S ~[]T](a, b S, keyFunc func(t T) K) S {
	lenA, lenB := len(a), len(b)
	if lenA == 0 {
		return ToSetBy(b, keyFunc)
	}
	if lenB == 0 {
		return ToSetBy(a, keyFunc)
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

// Deprecated: use UnionBy instead
func UnionPred[T any, K comparable, S ~[]T](a, b S, keyFunc func(t T) K) S {
	return UnionBy(a, b, keyFunc)
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

// IntersectionBy returns all unique shared values from multiple slices with key function
func IntersectionBy[T any, K comparable, S ~[]T](a, b S, keyFunc func(t T) K) S {
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

// Deprecated: use IntersectionBy instead
func IntersectionPred[T any, K comparable, S ~[]T](a, b S, keyFunc func(t T) K) S {
	return IntersectionBy(a, b, keyFunc)
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

// DifferenceBy calculates the differences between two slices using special key function.
// NOTE: this function does not return unique values.
func DifferenceBy[T any, K comparable, S ~[]T](a, b S, keyFunc func(t T) K) (S, S) {
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

// Deprecated: use DifferenceBy instead
func DifferencePred[T any, K comparable, S ~[]T](a, b S, keyFunc func(t T) K) (S, S) {
	return DifferenceBy(a, b, keyFunc)
}
