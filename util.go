package gofn

import "reflect"

// If returns the 2nd arg if the condition is true, 3rd arg otherwise.
// This is similar to C-language ternary operation (cond ? val1 : val2).
// Deprecated: this function may cause unexpected behavior upon misuses.
//
//	For example: gofn.If(len(slice) > 0, slice[0], defaultVal) will crash if slice is empty
func If[C bool, T any](cond C, v1 T, v2 T) T {
	if cond {
		return v1
	}
	return v2
}

// FirstNonEmpty returns the first non-empty value in the given arguments if found, otherwise
// returns the zero value. This function use reflection.
//
// Non-empty value must be not:
//   - zero value (0, "", nil, false)
//   - empty slice, array, map, channel
//   - pointer points to zero value
func FirstNonEmpty[T any](args ...T) (val T) {
	for i := range args {
		if isNonEmptyValue(reflect.ValueOf(args[i])) {
			return args[i]
		}
	}
	return
}

func isNonEmptyValue(v reflect.Value) bool {
	if !v.IsValid() || v.IsZero() {
		return false
	}
	k := v.Kind()
	if k == reflect.Pointer || k == reflect.Interface {
		return isNonEmptyValue(v.Elem())
	}
	if k == reflect.Slice || k == reflect.Array || k == reflect.Map || k == reflect.Chan {
		return v.Len() > 0
	}
	return true
}

// Coalesce returns the first non-zero value if found, otherwise returns zero.
func Coalesce[T comparable](args ...T) (result T) {
	for _, v := range args {
		if v != result {
			return v
		}
	}
	return result
}

func Must1(e error) {
	if e != nil {
		panic(e)
	}
}

func Must2[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}

// Must is the same as Must2
func Must[T any](v T, e error) T {
	if e != nil {
		panic(e)
	}
	return v
}

func Must3[T1, T2 any](v1 T1, v2 T2, e error) (T1, T2) {
	if e != nil {
		panic(e)
	}
	return v1, v2
}

func Must4[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, e error) (T1, T2, T3) {
	if e != nil {
		panic(e)
	}
	return v1, v2, v3
}

func Must5[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, e error) (T1, T2, T3, T4) {
	if e != nil {
		panic(e)
	}
	return v1, v2, v3, v4
}

func Must6[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, e error) (T1, T2, T3, T4, T5) {
	if e != nil {
		panic(e)
	}
	return v1, v2, v3, v4, v5
}

// ToPtr returns pointer to the address of the input
func ToPtr[T any](t T) *T {
	return &t
}

// Head returns the first argument
func Head[T any](t T, s ...any) T {
	return t
}

// Tail returns the last argument
func Tail[T any](t any, s ...any) (T, bool) {
	v := t
	if len(s) > 0 {
		v = s[len(s)-1]
	}
	ret, ok := v.(T)
	return ret, ok
}
