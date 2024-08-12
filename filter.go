package gofn

import (
	"strings"
	"unsafe"
)

// Filter filters slice elements with condition.
func Filter[T any, S ~[]T](s S, filterFunc func(t T) bool) S {
	result := make(S, 0, len(s))
	for i := range s {
		if filterFunc(s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

// FilterPtr filters slice elements using pointer in callback.
// This function is faster than Filter() when used on slices of structs.
func FilterPtr[T any, S ~[]T](s S, filterFunc func(t *T) bool) S {
	result := make(S, 0, len(s))
	for i := range s {
		if filterFunc(&s[i]) {
			result = append(result, s[i])
		}
	}
	return result
}

// FilterLT returns all values which are less than the specified value
func FilterLT[T NumberExt | StringExt, S ~[]T](s S, v T) S {
	return Filter(s, func(t T) bool { return t < v })
}

// FilterLTE returns all values which are less than or equal to the specified value
func FilterLTE[T NumberExt | StringExt, S ~[]T](s S, v T) S {
	return Filter(s, func(t T) bool { return t <= v })
}

// FilterGT returns all values which are greater than the specified value
func FilterGT[T NumberExt | StringExt, S ~[]T](s S, v T) S {
	return Filter(s, func(t T) bool { return t > v })
}

// FilterGTE returns all values which are greater than or equal to the specified value
func FilterGTE[T NumberExt | StringExt, S ~[]T](s S, v T) S {
	return Filter(s, func(t T) bool { return t >= v })
}

// FilterNE returns all values which are not equal to the specified value
func FilterNE[T comparable, S ~[]T](s S, v T) S {
	return Filter(s, func(t T) bool { return t != v })
}

// FilterIN returns all values which are present in the specified list
func FilterIN[T comparable, S ~[]T](s S, v ...T) S {
	m := make(map[T]struct{}, len(v))
	for _, t := range v {
		m[t] = struct{}{}
	}
	return Filter(s, func(t T) bool {
		_, ok := m[t]
		return ok
	})
}

// FilterNIN returns all values which are not present in the specified list
func FilterNIN[T comparable, S ~[]T](s S, v ...T) S {
	m := make(map[T]struct{}, len(v))
	for _, t := range v {
		m[t] = struct{}{}
	}
	return Filter(s, func(t T) bool {
		_, ok := m[t]
		return !ok
	})
}

// FilterLIKE returns all strings which contain the specified substring.
// Don't use wildcard in the input string.
// For example: FilterLIKE(names, "tom").
func FilterLIKE[T StringExt, S ~[]T](s S, v string) S {
	if len(v) == 0 {
		return S{}
	}
	return Filter(s, func(t T) bool {
		return strings.Contains(*(*string)(unsafe.Pointer(&t)), v)
	})
}

// FilterILIKE returns all strings which contain the specified substring with case-insensitive
func FilterILIKE[T StringExt, S ~[]T](s S, v string) S {
	if len(v) == 0 {
		return S{}
	}
	v = strings.ToLower(v)
	return Filter(s, func(t T) bool {
		return strings.Contains(strings.ToLower(*(*string)(unsafe.Pointer(&t))), v)
	})
}
