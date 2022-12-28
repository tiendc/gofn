package gofn

func All[T comparable](s ...T) bool {
	var zeroT T
	for i := range s {
		if s[i] == zeroT {
			return false
		}
	}
	return true
}

func Any[T comparable](s ...T) bool {
	var zeroT T
	for i := range s {
		if s[i] != zeroT {
			return true
		}
	}
	return false
}
