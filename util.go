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

// Or logically selects the first value which is not zero value of type T.
// This function is similar to `FirstTrue`, but it uses generic, not reflection.
func Or[T NumberEx | NumberPtr | StringEx | StringPtr | ~bool | *bool](args ...T) T {
	var defaultVal T
	for _, v := range args {
		if v != defaultVal {
			return v
		}
	}
	return defaultVal
}

// FirstTrue returns the first "true" value in the given arguments if found.
// True value must be not:
//   - zero value (0, "", nil, false)
//   - empty slice, array, map, channel
//   - pointer points to zero value
func FirstTrue[T any](a0 T, args ...T) T {
	a := a0
	for i := -1; i < len(args); i++ {
		if i >= 0 {
			a = args[i]
		}
		if isTrueValue(reflect.ValueOf(a)) {
			return a
		}
	}
	return a0
}

func isTrueValue(v reflect.Value) bool {
	if !v.IsValid() || v.IsZero() {
		return false
	}
	k := v.Kind()
	if k == reflect.Pointer || k == reflect.Interface {
		return isTrueValue(v.Elem())
	}
	if k == reflect.Slice || k == reflect.Array || k == reflect.Map || k == reflect.Chan {
		return v.Len() > 0
	}
	return true
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

// New returns pointer to the address of the input.
// Helpful for unit testing when create a struct that has field requires pointer type.
func New[T any](t T) *T {
	return &t
}

// Head returns the first argument
func Head[T any](t T, s ...any) T {
	return t
}

// HeadOf returns the first item in slice.
// Returns zero value and `false` if the slice is empty.
func HeadOf[T any, S ~[]T](s S) (T, bool) {
	if len(s) > 0 {
		return s[0], true
	}
	var zero T
	return zero, false
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

// TailOf returns the last item in slice.
// Returns zero value and `false` if the slice is empty.
func TailOf[T any, S ~[]T](s S) (T, bool) {
	if len(s) > 0 {
		return s[len(s)-1], true
	}
	var zero T
	return zero, false
}
