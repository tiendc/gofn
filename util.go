package gofn

// If returns the 2nd arg if the condition is true, 3rd arg otherwise
// This is similar to C-language ternary operation (cond ? val1 : val2)
func If[C bool, T any](cond C, v1 T, v2 T) T {
	if cond {
		return v1
	}
	return v2
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
