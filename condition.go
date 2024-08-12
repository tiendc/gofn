package gofn

// All returns true if all given values are evaluated as `true`
func All[T comparable](s ...T) bool {
	var zeroT T
	for i := range s {
		if s[i] == zeroT {
			return false
		}
	}
	return true
}

// Any returns true if at least one of given values is evaluated as `true`
func Any[T comparable](s ...T) bool {
	var zeroT T
	for i := range s {
		if s[i] != zeroT {
			return true
		}
	}
	return false
}
