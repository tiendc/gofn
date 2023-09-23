package gofn

// Equal compares 2 slices with preserving order
func Equal[T comparable](a, b []T) bool {
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

// EqualPred compares 2 slices with preserving order
func EqualPred[T any](a, b []T, equalFunc func(a, b T) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalFunc(a[i], b[i]) {
			return false
		}
	}
	return true
}

// EqualPredPtr compares 2 slices with preserving order
func EqualPredPtr[T any](a, b []T, equalFunc func(a, b *T) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !equalFunc(&a[i], &b[i]) {
			return false
		}
	}
	return true
}

// ContentEqual compares 2 slices without caring about order
// NOTE: if you want to compare content of slices of pointers, use ContentEqualPtr
func ContentEqual[T comparable](a, b []T) bool {
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
func ContentEqualPtr[T comparable](a, b []*T) bool {
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

// ContentEqualPred compares 2 slices without preserving order
func ContentEqualPred[T any, K comparable](a, b []T, keyFunc func(t T) K) bool {
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

// Concat concatenates slices
func Concat[T any](slices ...[]T) []T {
	capacity := 0
	for _, s := range slices {
		capacity += len(s)
	}
	result := make([]T, 0, capacity)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// Contain tests if a slice contains an item
func Contain[T comparable](a []T, t T) bool {
	for i := range a {
		if a[i] == t {
			return true
		}
	}
	return false
}

// ContainPred tests if a slice contains an item by predicate
func ContainPred[T any](a []T, pred func(t T) bool) bool {
	for i := range a {
		if pred(a[i]) {
			return true
		}
	}
	return false
}

// ContainPredPtr tests if a slice contains an item by predicate
func ContainPredPtr[T any](a []T, pred func(t *T) bool) bool {
	for i := range a {
		if pred(&a[i]) {
			return true
		}
	}
	return false
}

// ContainAll tests if a slice contains all given values
func ContainAll[T comparable](a []T, b ...T) bool {
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
func ContainAny[T comparable](a []T, b ...T) bool {
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
func IsUnique[T comparable](s []T) bool {
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

// IsUniquePred checks a slice for uniqueness using key function
func IsUniquePred[T any, U comparable](s []T, keyFunc func(t T) U) bool {
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

// FindPred finds value in slice by predicate
func FindPred[T any](a []T, pred func(t T) bool) (t T, found bool) {
	for i := range a {
		if pred(a[i]) {
			return a[i], true
		}
	}
	return t, false
}

// FindPredPtr finds value in slice by predicate
func FindPredPtr[T any](a []T, pred func(t *T) bool) (t T, found bool) {
	for i := range a {
		if pred(&a[i]) {
			return a[i], true
		}
	}
	return t, false
}

// FindLastPred finds value in slice from the end by predicate
func FindLastPred[T any](a []T, pred func(t T) bool) (t T, found bool) {
	for i := len(a) - 1; i >= 0; i-- {
		if pred(a[i]) {
			return a[i], true
		}
	}
	return t, false
}

// FindLastPredPtr finds value in slice from the end by predicate
func FindLastPredPtr[T any](a []T, pred func(t *T) bool) (t T, found bool) {
	for i := len(a) - 1; i >= 0; i-- {
		if pred(&a[i]) {
			return a[i], true
		}
	}
	return t, false
}

// IndexOf gets index of item in slice
// Returns -1 if not found
func IndexOf[T comparable](a []T, t T) int {
	for i := range a {
		if a[i] == t {
			return i
		}
	}
	return -1
}

// IndexOfPred gets index of item in slice by predicate
// Returns -1 if not found
func IndexOfPred[T any](a []T, pred func(t T) bool) int {
	for i := range a {
		if pred(a[i]) {
			return i
		}
	}
	return -1
}

// LastIndexOf gets index of item from the end in slice
// Returns -1 if not found
func LastIndexOf[T comparable](a []T, t T) int {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == t {
			return i
		}
	}
	return -1
}

// LastIndexOfPred gets index of item from the end in slice
// Returns -1 if not found
func LastIndexOfPred[T any](a []T, pred func(t T) bool) int {
	for i := len(a) - 1; i >= 0; i-- {
		if pred(a[i]) {
			return i
		}
	}
	return -1
}

// RemoveAt removes element at the specified index
func RemoveAt[T any](ps *[]T, i int) {
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
func FastRemoveAt[T any](ps *[]T, i int) {
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
func Remove[T comparable](ps *[]T, v T) bool {
	i := IndexOf(*ps, v)
	if i == -1 {
		return false
	}
	RemoveAt(ps, i)
	return true
}

// FastRemove removes element value
func FastRemove[T comparable](ps *[]T, v T) bool {
	i := IndexOf(*ps, v)
	if i == -1 {
		return false
	}
	FastRemoveAt(ps, i)
	return true
}

// RemoveLastOf removes element value
func RemoveLastOf[T comparable](ps *[]T, v T) bool {
	i := LastIndexOf(*ps, v)
	if i == -1 {
		return false
	}
	RemoveAt(ps, i)
	return true
}

// FastRemoveLastOf removes element value
func FastRemoveLastOf[T comparable](ps *[]T, v T) bool {
	i := LastIndexOf(*ps, v)
	if i == -1 {
		return false
	}
	FastRemoveAt(ps, i)
	return true
}

// RemoveAll removes all occurrences of value
func RemoveAll[T comparable](ps *[]T, v T) int {
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

func Compact[T comparable](s []T) []T {
	result := make([]T, 0, len(s))
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
func Replace[T comparable](s []T, value, replacement T) bool {
	for i := range s {
		if s[i] == value {
			s[i] = replacement
			return true
		}
	}
	return false
}

// ReplaceAll replaces a value in slice with another value
func ReplaceAll[T comparable](s []T, value, replacement T) int {
	count := 0
	for i := range s {
		if s[i] == value {
			s[i] = replacement
			count++
		}
	}
	return count
}

// Fill sets slice element values
func Fill[T any](a []T, val T) {
	for i := range a {
		a[i] = val
	}
}

// CountValue counts number of occurrences of an item in the slice
func CountValue[T comparable](a []T, val T) int {
	count := 0
	for i := range a {
		if a[i] == val {
			count++
		}
	}
	return count
}

// CountValuePred counts number of occurrences of an item in the slice
func CountValuePred[T any](a []T, pred func(t T) bool) int {
	count := 0
	for i := range a {
		if pred(a[i]) {
			count++
		}
	}
	return count
}

// ContainSlice tests if a slice contains a slice
func ContainSlice[T comparable](a, b []T) bool {
	return IndexOfSlice(a, b) >= 0
}

// IndexOfSlice gets index of sub-slice in slice
// Returns -1 if not found
func IndexOfSlice[T comparable](a, sub []T) int {
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
func LastIndexOfSlice[T comparable](a []T, sub []T) int {
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

func GetFirst[T any](s []T, defaultVal T) T {
	if len(s) > 0 {
		return s[0]
	}
	return defaultVal
}

func GetLast[T any](s []T, defaultVal T) T {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return defaultVal
}
