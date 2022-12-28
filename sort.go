package gofn

import "sort"

// Sort sorts slice values
func Sort[T NumberEx | StringEx](s []T) []T {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}

// SortDesc sorts slice values in descending order
func SortDesc[T NumberEx | StringEx](s []T) []T {
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
	return s
}

// SortEx sorts slice values
func SortEx[T any](s []T, less func(i, j int) bool) []T {
	sort.Slice(s, less)
	return s
}

// SortStable sorts slice values
func SortStable[T NumberEx | StringEx](s []T) []T {
	sort.SliceStable(s, func(i, j int) bool { return s[i] < s[j] })
	return s
}

// SortStableDesc sorts slice values in descending order
func SortStableDesc[T NumberEx | StringEx](s []T) []T {
	sort.SliceStable(s, func(i, j int) bool { return s[i] > s[j] })
	return s
}

// SortStableEx sorts slice values
func SortStableEx[T any](s []T, less func(i, j int) bool) []T {
	sort.SliceStable(s, less)
	return s
}

// IsSorted checks if a slice is sorted
func IsSorted[T NumberEx | StringEx](s []T) bool {
	return sort.SliceIsSorted(s, func(i, j int) bool { return s[i] < s[j] })
}

// IsSortedDesc checks if a slice is sorted in descending order
func IsSortedDesc[T NumberEx | StringEx](s []T) bool {
	return sort.SliceIsSorted(s, func(i, j int) bool { return s[i] > s[j] })
}
