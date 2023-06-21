package gofn

import "reflect"

// If returns the 2nd arg if the condition is true, 3rd arg otherwise
// This is similar to C-language ternary operation (cond ? val1 : val2)
// Deprecated: this function may cause unexpected behavior upon misuses
//
//	For example: gofn.If(len(slice) > 0, slice[0], dafaultVal) will crash if slice is empty
func If[C bool, T any](cond C, v1 T, v2 T) T {
	if cond {
		return v1
	}
	return v2
}

// FirstTrue returns the first "true" value in the given arguments if found
// True value is not:
//   - zero value (0, "", nil, false)
//   - empty slice, array, map, channel
func FirstTrue[T any](a0 T, args ...T) T {
	a := a0
	for i := -1; i < len(args); i++ {
		if i >= 0 {
			a = args[i]
		}
		v := reflect.ValueOf(a)
		if !v.IsValid() || v.IsZero() {
			continue
		}
		k := v.Kind()
		if k == reflect.Slice || k == reflect.Array || k == reflect.Map || k == reflect.Chan {
			if v.Len() > 0 {
				return a
			}
			continue
		}
		return a
	}
	return a0
}

func Must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}

// New returns pointer to the address of the input
// Helpful for unit testing when create a struct that has field requires pointer type
func New[T any](t T) *T {
	return &t
}

// Head returns the first argument
func Head[T any](t T, s ...interface{}) T {
	return t
}

// Tail returns the last argument
func Tail[T any](t interface{}, s ...interface{}) (T, bool) {
	v := t
	if len(s) > 0 {
		v = s[len(s)-1]
	}
	ret, ok := v.(T)
	return ret, ok
}
