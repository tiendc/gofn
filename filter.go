package gofn

import (
	"strings"
	"unsafe"
)

// Filter filters slice elements with condition.
func Filter[T any](s []T, filterFunc func(t T) bool) []T {
	result := make([]T, 0, len(s))
	for i := range s {
		if filterFunc(s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

// FilterPtr filters slice elements using pointer in callback.
// This function is faster than Filter() when used on slices of structs.
func FilterPtr[T any](s []T, filterFunc func(t *T) bool) []T {
	result := make([]T, 0, len(s))
	for i := range s {
		if filterFunc(&s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

func FilterLT[T NumberEx | StringEx](s []T, v T) []T {
	return Filter(s, func(t T) bool { return t < v })
}

func FilterLTE[T NumberEx | StringEx](s []T, v T) []T {
	return Filter(s, func(t T) bool { return t <= v })
}

func FilterGT[T NumberEx | StringEx](s []T, v T) []T {
	return Filter(s, func(t T) bool { return t > v })
}

func FilterGTE[T NumberEx | StringEx](s []T, v T) []T {
	return Filter(s, func(t T) bool { return t >= v })
}

func FilterNE[T comparable](s []T, v T) []T {
	return Filter(s, func(t T) bool { return t != v })
}

func FilterIN[T comparable](s []T, v ...T) []T {
	m := make(map[T]struct{}, len(v))
	for _, t := range v {
		m[t] = struct{}{}
	}
	return Filter(s, func(t T) bool {
		_, ok := m[t]
		return ok
	})
}

func FilterNIN[T comparable](s []T, v ...T) []T {
	m := make(map[T]struct{}, len(v))
	for _, t := range v {
		m[t] = struct{}{}
	}
	return Filter(s, func(t T) bool {
		_, ok := m[t]
		return !ok
	})
}

func FilterLIKE[T StringEx](s []T, v string) []T {
	if len(v) == 0 {
		return []T{}
	}
	return Filter(s, func(t T) bool {
		return strings.Contains(*(*string)(unsafe.Pointer(&t)), v)
	})
}

func FilterILIKE[T StringEx](s []T, v string) []T {
	if len(v) == 0 {
		return []T{}
	}
	v = strings.ToLower(v)
	return Filter(s, func(t T) bool {
		return strings.Contains(strings.ToLower(*(*string)(unsafe.Pointer(&t))), v)
	})
}
