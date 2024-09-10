package gofn

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
