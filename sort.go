package gofn

import "sort"

// Sort sorts slice values
func Sort[T NumberExt | StringExt, S ~[]T](s S) S {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}

// SortDesc sorts slice values in descending order
func SortDesc[T NumberExt | StringExt, S ~[]T](s S) S {
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
	return s
}

// SortEx sorts slice values
func SortEx[T any, S ~[]T](s S, less func(i, j int) bool) S {
	sort.Slice(s, less)
	return s
}

// SortStable sorts slice values
func SortStable[T NumberExt | StringExt, S ~[]T](s S) S {
	sort.SliceStable(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}

// SortStableDesc sorts slice values in descending order
func SortStableDesc[T NumberExt | StringExt, S ~[]T](s S) S {
	sort.SliceStable(s, func(i, j int) bool { return s[i] > s[j] })
	return s
}

// SortStableEx sorts slice values
func SortStableEx[T any, S ~[]T](s S, less func(i, j int) bool) S {
	sort.SliceStable(s, less)
	return s
}

// IsSorted checks if a slice is sorted
func IsSorted[T NumberExt | StringExt, S ~[]T](s S) bool {
	return sort.SliceIsSorted(s, func(i, j int) bool { return s[i] < s[j] })
}

// IsSortedDesc checks if a slice is sorted in descending order
func IsSortedDesc[T NumberExt | StringExt, S ~[]T](s S) bool {
	return sort.SliceIsSorted(s, func(i, j int) bool { return s[i] > s[j] })
}
