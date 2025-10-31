package gofn

// Compact excludes all zero items in a slice
func Compact[T comparable, S ~[]T](s S) S {
	result := make(S, 0, len(s))
	var zeroT T
	for _, v := range s {
		if v == zeroT {
			continue
		}
		result = append(result, v)
	}
	return result
}

// Drop returns a copied slice with dropping items in the list
func Drop[T comparable, S ~[]T](a S, values ...T) S {
	return FilterNIN(a, values...)
}

// Reverse reverses slice content, this modifies the slice
func Reverse[T any, S ~[]T](s S) S {
	if len(s) == 0 {
		return s
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ReverseCopy returns a new slice which has content in reversed order
func ReverseCopy[T any, S ~[]T](s S) S {
	result := make(S, len(s))
	for i, j := 0, len(s)-1; j >= 0; j-- {
		result[i] = s[j]
		i++
	}
	return result
}

// Shuffle items of a slice and returns a new slice
func Shuffle[T any, S ~[]T](s S, randFuncs ...func(n int) int) S {
	if len(s) <= 1 {
		return append(S{}, s...)
	}
	maker := NewRandChoiceMaker(s, randFuncs...)
	result := make(S, 0, len(s))
	for {
		item, valid := maker.Next()
		if !valid {
			break
		}
		result = append(result, item)
	}
	return result
}

// Chunk splits slice content into chunks by chunk size
func Chunk[T any, S ~[]T](s S, chunkSize int) []S {
	if chunkSize <= 0 {
		return []S{}
	}

	chunks := make([]S, 0, len(s)/chunkSize+1)
	for {
		if len(s) == 0 {
			break
		}
		if len(s) < chunkSize {
			chunkSize = len(s)
		}
		chunks = append(chunks, s[0:chunkSize])
		s = s[chunkSize:]
	}
	return chunks
}

// ChunkByPieces splits slice content into chunks by number of pieces
func ChunkByPieces[T any, S ~[]T](s S, chunkCount int) []S {
	if chunkCount <= 0 {
		return []S{}
	}
	chunkSize := len(s) / chunkCount
	if chunkSize*chunkCount < len(s) {
		chunkSize++
	}

	return Chunk(s, chunkSize)
}

// Reduce reduces a slice to a value
func Reduce[T any, S ~[]T](s S, reduceFunc func(accumulator, currentValue T) T) T {
	length := len(s)
	if length == 0 {
		var zeroT T
		return zeroT
	}
	accumulator := s[0]
	for i := 1; i < length; i++ {
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

// ReduceReverse reduces a slice to a value
func ReduceReverse[T any, S ~[]T](s S, reduceFunc func(accumulator, currentValue T) T) T {
	length := len(s)
	if length == 0 {
		var zeroT T
		return zeroT
	}
	accumulator := s[length-1]
	for i := length - 2; i >= 0; i-- { //nolint:mnd
		accumulator = reduceFunc(accumulator, s[i])
	}
	return accumulator
}

// ReduceReverseEx reduces a slice to a value with custom initial value
func ReduceReverseEx[T any, U any, S ~[]T](
	s S,
	reduceFunc func(accumulator U, currentValue T, currentIndex int) U,
	initVal U,
) U {
	accumulator := initVal
	for i := len(s) - 1; i >= 0; i-- {
		accumulator = reduceFunc(accumulator, s[i], i)
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

// Flatten flattens 2-dimensional slices.
// E.g. Flatten([1,2,3], [3,4,5]) -> [1,2,3,3,4,5].
func Flatten[T any, S ~[]T](s ...S) S {
	result := make(S, 0, len(s)*5) //nolint:mnd
	for _, innerSlice := range s {
		result = append(result, innerSlice...)
	}
	return result
}

// Flatten3 flattens 3-dimensional slices
func Flatten3[T any, S ~[]T, SS ~[]S](s ...SS) S {
	result := make(S, 0, len(s)*30) //nolint:mnd
	for _, innerSlice := range s {
		for _, mostInnerSlice := range innerSlice {
			result = append(result, mostInnerSlice...)
		}
	}
	return result
}

// Zip combines values from 2 slices by each position
func Zip[T1, T2 any, S1 ~[]T1, S2 ~[]T2](slice1 S1, slice2 S2) []*Tuple2[T1, T2] {
	minLen := len(slice1)
	if minLen > len(slice2) {
		minLen = len(slice2)
	}
	result := make([]*Tuple2[T1, T2], minLen)
	for i := 0; i < minLen; i++ {
		result[i] = &Tuple2[T1, T2]{slice1[i], slice2[i]}
	}
	return result
}

// Zip3 combines values from 3 slices by each position
func Zip3[T1, T2, T3 any, S1 ~[]T1, S2 ~[]T2, S3 ~[]T3](slice1 S1, slice2 S2, slice3 S3) []*Tuple3[T1, T2, T3] {
	minLen := Min(len(slice1), len(slice2), len(slice3))
	result := make([]*Tuple3[T1, T2, T3], minLen)
	for i := 0; i < minLen; i++ {
		result[i] = &Tuple3[T1, T2, T3]{slice1[i], slice2[i], slice3[i]}
	}
	return result
}

// Zip4 combines values from 4 slices by each position
func Zip4[T1, T2, T3, T4 any, S1 ~[]T1, S2 ~[]T2, S3 ~[]T3, S4 ~[]T4](
	slice1 S1, slice2 S2, slice3 S3, slice4 S4,
) []*Tuple4[T1, T2, T3, T4] {
	minLen := Min(len(slice1), len(slice2), len(slice3), len(slice4))
	result := make([]*Tuple4[T1, T2, T3, T4], minLen)
	for i := 0; i < minLen; i++ {
		result[i] = &Tuple4[T1, T2, T3, T4]{slice1[i], slice2[i], slice3[i], slice4[i]}
	}
	return result
}

// Zip5 combines values from 4 slices by each position
func Zip5[T1, T2, T3, T4, T5 any, S1 ~[]T1, S2 ~[]T2, S3 ~[]T3, S4 ~[]T4, S5 ~[]T5](
	slice1 S1, slice2 S2, slice3 S3, slice4 S4, slice5 S5,
) []*Tuple5[T1, T2, T3, T4, T5] {
	minLen := Min(len(slice1), len(slice2), len(slice3), len(slice4), len(slice5))
	result := make([]*Tuple5[T1, T2, T3, T4, T5], minLen)
	for i := 0; i < minLen; i++ {
		result[i] = &Tuple5[T1, T2, T3, T4, T5]{slice1[i], slice2[i], slice3[i], slice4[i], slice5[i]}
	}
	return result
}
