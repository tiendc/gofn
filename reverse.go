package gofn

// Reverse func Reverse[S ~[]E, E any](s S)  {
func Reverse[T any, S ~[]T](s S) S {
	if len(s) == 0 {
		return s
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseCopy[T any, S ~[]T](s S) S {
	result := make(S, len(s))
	for i, j := 0, len(s)-1; j >= 0; j-- {
		result[i] = s[j]
		i++
	}
	return result
}
