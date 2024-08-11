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

func FilterLT[T NumberEx | StringEx, S ~[]T](s S, v T) S {
	return Filter(s, func(t T) bool { return t < v })
}

func FilterLTE[T NumberEx | StringEx, S ~[]T](s S, v T) S {
	return Filter(s, func(t T) bool { return t <= v })
}

func FilterGT[T NumberEx | StringEx, S ~[]T](s S, v T) S {
	return Filter(s, func(t T) bool { return t > v })
}

func FilterGTE[T NumberEx | StringEx, S ~[]T](s S, v T) S {
	return Filter(s, func(t T) bool { return t >= v })
}

func FilterNE[T comparable, S ~[]T](s S, v T) S {
	return Filter(s, func(t T) bool { return t != v })
}

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

func FilterLIKE[T StringEx, S ~[]T](s S, v string) S {
	if len(v) == 0 {
		return S{}
	}
	return Filter(s, func(t T) bool {
		return strings.Contains(*(*string)(unsafe.Pointer(&t)), v)
	})
}

func FilterILIKE[T StringEx, S ~[]T](s S, v string) S {
	if len(v) == 0 {
		return S{}
	}
	v = strings.ToLower(v)
	return Filter(s, func(t T) bool {
		return strings.Contains(strings.ToLower(*(*string)(unsafe.Pointer(&t))), v)
	})
}
