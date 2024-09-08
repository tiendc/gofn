package gofn

// Equal compares 2 slices with preserving order
func Equal[T comparable, S ~[]T](a, b S) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// EqualBy compares 2 slices with preserving order
func EqualBy[T any, S ~[]T](a, b S, equalCmp func(a, b T) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalCmp(a[i], b[i]) {
			return false
		}
	}
	return true
}

// Deprecated: use EqualBy instead
func EqualPred[T any, S ~[]T](a, b S, equalFunc func(a, b T) bool) bool {
	return EqualBy(a, b, equalFunc)
}

// EqualByPtr compares 2 slices with preserving order
func EqualByPtr[T any, S ~[]T](a, b S, equalCmp func(a, b *T) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalCmp(&a[i], &b[i]) {
			return false
		}
	}
	return true
}

// Deprecated: use EqualByPtr instead
func EqualPredPtr[T any, S ~[]T](a, b S, equalCmp func(a, b *T) bool) bool {
	return EqualByPtr(a, b, equalCmp)
}

// ContentEqual compares 2 slices without caring about order.
// NOTE: if you want to compare content of slices of pointers, use ContentEqualPtr.
func ContentEqual[T comparable, S ~[]T](a, b S) bool {
	if len(a) != len(b) {
		return false
	}

	mapA := make(map[T]*int, len(a))
	for i := range a {
		k := a[i]
		if count, ok := mapA[k]; !ok {
			cnt := 1
			mapA[k] = &cnt
		} else {
			*count++
		}
	}

	for i := range b {
		k := b[i]
		switch count, ok := mapA[k]; {
		case !ok:
			return false
		case *count == 1:
			delete(mapA, k)
		default:
			*count--
		}
	}

	return len(mapA) == 0
}

// ContentEqualPtr compares 2 slices of pointers without caring about order
func ContentEqualPtr[T comparable, S ~[]*T](a, b S) bool {
	if len(a) != len(b) {
		return false
	}

	mapA := make(map[T]*int, len(a))
	for i := range a {
		k := *a[i]
		if count, ok := mapA[k]; !ok {
			cnt := 1
			mapA[k] = &cnt
		} else {
			*count++
		}
	}

	for i := range b {
		k := *b[i]
		switch count, ok := mapA[k]; {
		case !ok:
			return false
		case *count == 1:
			delete(mapA, k)
		default:
			*count--
		}
	}

	return len(mapA) == 0
}

// ContentEqualBy compares 2 slices without preserving order
func ContentEqualBy[T any, K comparable, S ~[]T](a, b S, keyFunc func(t T) K) bool {
	if len(a) != len(b) {
		return false
	}

	mapA := make(map[K]*int, len(a))
	for i := range a {
		k := keyFunc(a[i])
		if count, ok := mapA[k]; !ok {
			cnt := 1
			mapA[k] = &cnt
		} else {
			*count++
		}
	}

	for i := range b {
		k := keyFunc(b[i])
		switch count, ok := mapA[k]; {
		case !ok:
			return false
		case *count == 1:
			delete(mapA, k)
		default:
			*count--
		}
	}

	return len(mapA) == 0
}

// Deprecated: use ContentEqualBy instead
func ContentEqualPred[T any, K comparable, S ~[]T](a, b S, keyFunc func(t T) K) bool {
	return ContentEqualBy(a, b, keyFunc)
}

// Concat concatenates slices
func Concat[T any, S ~[]T](slices ...S) S {
	capacity := 0
	for _, s := range slices {
		capacity += len(s)
	}
	result := make(S, 0, capacity)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// Contain tests if a slice contains an item
func Contain[T comparable, S ~[]T](a S, t T) bool {
	for i := range a {
		if a[i] == t {
			return true
		}
	}
	return false
}

// ContainBy tests if a slice contains an item by predicate
func ContainBy[T any, S ~[]T](a S, pred func(t T) bool) bool {
	for i := range a {
		if pred(a[i]) {
			return true
		}
	}
	return false
}

// Deprecated: use ContainBy instead
func ContainPred[T any, S ~[]T](a S, pred func(t T) bool) bool {
	return ContainBy(a, pred)
}

// ContainByPtr tests if a slice contains an item by predicate
func ContainByPtr[T any, S ~[]T](a S, pred func(t *T) bool) bool {
	for i := range a {
		if pred(&a[i]) {
			return true
		}
	}
	return false
}

// Deprecated: use ContainByPtr instead
func ContainPredPtr[T any, S ~[]T](a S, pred func(t *T) bool) bool {
	return ContainByPtr(a, pred)
}

// ContainAll tests if a slice contains all given values
func ContainAll[T comparable, S ~[]T](a S, b ...T) bool {
	lenA, lenB := len(a), len(b)
	if lenA == 0 || lenB == 0 {
		return false
	}

	if lenA < 10 && lenB < 3 {
		for i := range b {
			if !Contain(a, b[i]) {
				return false
			}
		}
		return true
	}

	seen := make(map[T]struct{}, lenA)
	for i := range a {
		seen[a[i]] = struct{}{}
	}
	for i := range b {
		if _, ok := seen[b[i]]; !ok {
			return false
		}
	}
	return true
}

// ContainAny tests if a slice contains any given value
func ContainAny[T comparable, S ~[]T](a S, b ...T) bool {
	lenA, lenB := len(a), len(b)
	if lenA == 0 || lenB == 0 {
		return false
	}

	if lenA < 10 && lenB < 3 {
		for i := range b {
			if Contain(a, b[i]) {
				return true
			}
		}
		return false
	}

	seen := make(map[T]struct{}, lenA)
	for i := range a {
		seen[a[i]] = struct{}{}
	}
	for i := range b {
		if _, ok := seen[b[i]]; ok {
			return true
		}
	}
	return false
}

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

// Deprecated: use IsUniqueBy instead
func IsUniquePred[T any, U comparable, S ~[]T](s S, keyFunc func(t T) U) bool {
	return IsUniqueBy(s, keyFunc)
}

// Find finds value in slice by predicate
func Find[T any, S ~[]T](a S, pred func(t T) bool) (t T, found bool) {
	for i := range a {
		if pred(a[i]) {
			return a[i], true
		}
	}
	return t, false
}

// Deprecated: use Find instead
func FindPred[T any, S ~[]T](a S, pred func(t T) bool) (t T, found bool) {
	return Find(a, pred)
}

// FindPtr finds value in slice by predicate
func FindPtr[T any, S ~[]T](a S, pred func(t *T) bool) (t T, found bool) {
	for i := range a {
		if pred(&a[i]) {
			return a[i], true
		}
	}
	return t, false
}

// Deprecated: use FindPtr instead
func FindPredPtr[T any, S ~[]T](a S, pred func(t *T) bool) (t T, found bool) {
	return FindPtr(a, pred)
}

// FindLast finds value in slice from the end by predicate
func FindLast[T any, S ~[]T](a S, pred func(t T) bool) (t T, found bool) {
	for i := len(a) - 1; i >= 0; i-- {
		if pred(a[i]) {
			return a[i], true
		}
	}
	return t, false
}

// Deprecated: use FindLast instead
func FindLastPred[T any, S ~[]T](a S, pred func(t T) bool) (t T, found bool) {
	return FindLast(a, pred)
}

// FindLastPtr finds value in slice from the end by predicate
func FindLastPtr[T any, S ~[]T](a S, pred func(t *T) bool) (t T, found bool) {
	for i := len(a) - 1; i >= 0; i-- {
		if pred(&a[i]) {
			return a[i], true
		}
	}
	return t, false
}

// Deprecated: use FindLastPtr instead
func FindLastPredPtr[T any, S ~[]T](a S, pred func(t *T) bool) (t T, found bool) {
	return FindLastPtr(a, pred)
}

// IndexOf gets index of item in slice.
// Returns -1 if not found.
func IndexOf[T comparable, S ~[]T](a S, t T) int {
	for i := range a {
		if a[i] == t {
			return i
		}
	}
	return -1
}

// IndexOfBy gets index of item in slice by predicate.
// Returns -1 if not found.
func IndexOfBy[T any, S ~[]T](a S, pred func(t T) bool) int {
	for i := range a {
		if pred(a[i]) {
			return i
		}
	}
	return -1
}

// Deprecated: use IndexOfBy instead
func IndexOfPred[T any, S ~[]T](a S, pred func(t T) bool) int {
	return IndexOfBy(a, pred)
}

// LastIndexOf gets index of item from the end in slice.
// Returns -1 if not found.
func LastIndexOf[T comparable, S ~[]T](a S, t T) int {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == t {
			return i
		}
	}
	return -1
}

// LastIndexOfBy gets index of item from the end in slice.
// Returns -1 if not found.
func LastIndexOfBy[T any, S ~[]T](a S, pred func(t T) bool) int {
	for i := len(a) - 1; i >= 0; i-- {
		if pred(a[i]) {
			return i
		}
	}
	return -1
}

// Deprecated: use LastIndexOfBy instead
func LastIndexOfPred[T any, S ~[]T](a S, pred func(t T) bool) int {
	return LastIndexOfBy(a, pred)
}

// RemoveAt removes element at the specified index
func RemoveAt[T any, S ~[]T](ps *S, i int) {
	s := *ps
	if i < 0 || i >= len(s) {
		panic(ErrIndexOutOfRange)
	}
	var zeroT T
	s[i] = zeroT
	*ps = s[:i]
	*ps = append(*ps, s[i+1:]...)
}

// FastRemoveAt removes element at the specified index by swapping it with the last item in slice
func FastRemoveAt[T any, S ~[]T](ps *S, i int) {
	s := *ps
	length := len(s)
	if i < 0 || i >= length {
		panic(ErrIndexOutOfRange)
	}
	s[i] = s[length-1]
	var zeroT T
	s[length-1] = zeroT
	*ps = s[:length-1]
}

// Remove removes element value
func Remove[T comparable, S ~[]T](ps *S, v T) bool {
	i := IndexOf(*ps, v)
	if i == -1 {
		return false
	}
	RemoveAt(ps, i)
	return true
}

// FastRemove removes element value
func FastRemove[T comparable, S ~[]T](ps *S, v T) bool {
	i := IndexOf(*ps, v)
	if i == -1 {
		return false
	}
	FastRemoveAt(ps, i)
	return true
}

// RemoveLastOf removes element value
func RemoveLastOf[T comparable, S ~[]T](ps *S, v T) bool {
	i := LastIndexOf(*ps, v)
	if i == -1 {
		return false
	}
	RemoveAt(ps, i)
	return true
}

// FastRemoveLastOf removes element value
func FastRemoveLastOf[T comparable, S ~[]T](ps *S, v T) bool {
	i := LastIndexOf(*ps, v)
	if i == -1 {
		return false
	}
	FastRemoveAt(ps, i)
	return true
}

// RemoveAll removes all occurrences of value
func RemoveAll[T comparable, S ~[]T](ps *S, v T) int {
	newIdx := 0
	count := 0
	s := *ps
	length := len(s)
	for i := 0; i < length; i++ {
		if s[i] != v {
			if newIdx != i {
				s[newIdx] = s[i]
			}
			newIdx++
		} else {
			count++
		}
	}
	*ps = s[:length-count]
	return count
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

// Replace replaces a value in slice with another value
func Replace[T comparable, S ~[]T](s S, value, replacement T) bool {
	for i := range s {
		if s[i] == value {
			s[i] = replacement
			return true
		}
	}
	return false
}

// ReplaceN replaces a value in slice for the first n-occurrences
func ReplaceN[T comparable, S ~[]T](s S, value, replacement T, n int) int {
	if n == 0 || n < -1 {
		return 0
	}
	count := 0
	for i := range s {
		if s[i] == value {
			s[i] = replacement
			count++
			if n != -1 && count == n {
				return count
			}
		}
	}
	return count
}

// ReplaceAll replaces a value in slice with another value
func ReplaceAll[T comparable, S ~[]T](s S, value, replacement T) int {
	return ReplaceN(s, value, replacement, -1)
}

// Fill sets slice element values
func Fill[T any, S ~[]T](a S, val T) {
	for i := range a {
		a[i] = val
	}
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

// CountValue counts number of occurrences of an item in the slice
func CountValue[T comparable, S ~[]T](a S, val T) int {
	count := 0
	for i := range a {
		if a[i] == val {
			count++
		}
	}
	return count
}

// CountValueBy counts number of occurrences of an item in the slice
func CountValueBy[T any, S ~[]T](a S, pred func(t T) bool) int {
	count := 0
	for i := range a {
		if pred(a[i]) {
			count++
		}
	}
	return count
}

// Deprecated: use CountValueBy instead
func CountValuePred[T any, S ~[]T](a S, pred func(t T) bool) int {
	return CountValueBy(a, pred)
}

// Drop returns a copied slice with dropping items in the list
func Drop[T comparable, S ~[]T](a S, values ...T) S {
	return FilterNIN(a, values...)
}

// ContainSlice tests if a slice contains a slice
func ContainSlice[T comparable, S ~[]T](a, b S) bool {
	return IndexOfSlice(a, b) >= 0
}

// IndexOfSlice gets index of sub-slice in slice.
// Returns -1 if not found.
func IndexOfSlice[T comparable, S ~[]T](a, sub S) int {
	lengthA := len(a)
	lengthSub := len(sub)
	if lengthSub == 0 || lengthA < lengthSub {
		return -1
	}
	sub1st := sub[0]
	for i, max := 0, lengthA-lengthSub; i <= max; i++ {
		if a[i] == sub1st {
			found := true
			for j := 1; j < lengthSub; j++ {
				if a[i+j] != sub[j] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}

// LastIndexOfSlice gets last index of sub-slice in slice
// Returns -1 if not found
func LastIndexOfSlice[T comparable, S ~[]T](a, sub S) int {
	lengthA := len(a)
	lengthSub := len(sub)
	if lengthSub == 0 || lengthA < lengthSub {
		return -1
	}
	sub1st := sub[0]
	for i := lengthA - lengthSub; i >= 0; i-- {
		if a[i] == sub1st {
			found := true
			for j := 1; j < lengthSub; j++ {
				if a[i+j] != sub[j] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}

// GetFirst gets the first item in slice.
// Returns the default value if slice is empty.
func GetFirst[T any, S ~[]T](s S, defaultVal T) T {
	if len(s) > 0 {
		return s[0]
	}
	return defaultVal
}

// GetLast gets the last item in slice.
// Returns the default value if slice is empty.
func GetLast[T any, S ~[]T](s S, defaultVal T) T {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return defaultVal
}

// SubSlice gets sub slice from a slice.
// Passing negative numbers to get items from the end of the slice.
// For example, using start=-1, end=-2 to get the last item of the slice
// end param is exclusive.
func SubSlice[T any, S ~[]T](s S, start, end int) S {
	length := len(s)
	if length == 0 {
		return S{}
	}

	for start < 0 {
		start += length
	}
	if start > length {
		start = length
	}
	for end < 0 {
		end += length
	}
	if end > length {
		end = length
	}

	if start > end {
		// NOTE: end is exclusive
		return s[end+1 : start+1]
	}
	return s[start:end]
}

// SliceByRange generates a slice by range.
// start is inclusive, end is exclusive.
func SliceByRange[T NumberExt](start, end, step T) []T {
	if end > start {
		if step <= 0 {
			return []T{}
		}

		count := int((end-start)/step) + 1
		result := make([]T, 0, count)
		for i := start; i < end; i += step {
			result = append(result, i)
		}
		return result
	} else {
		if step >= 0 {
			return []T{}
		}

		count := int((end-start)/step) + 1
		result := make([]T, 0, count)
		for i := start; i > end; i += step {
			result = append(result, i)
		}
		return result
	}
}
