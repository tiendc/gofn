package gofn

// ToIfaceSlice convert a slice to a slice of interface
func ToIfaceSlice[T any, S ~[]T](s S) []any {
	result := make([]any, len(s))
	for i := range s {
		result[i] = s[i]
	}
	return result
}

// Deprecated: use ToIfaceSlice instead
func ToIntfSlice[T any, S ~[]T](s S) []any {
	return ToIfaceSlice(s)
}

// ToStringSlice converts str-approximate slice to string slice
func ToStringSlice[U, T ~string, S ~[]T](slice S) []U {
	result := make([]U, len(slice))
	for i := range slice {
		result[i] = U(slice[i])
	}
	return result
}

// ToNumberSlice converts int-approximate slice to int slice
func ToNumberSlice[U, T NumberExt, S ~[]T](slice S) []U {
	result := make([]U, len(slice))
	for i := range slice {
		result[i] = U(slice[i])
	}
	return result
}

// ToSlice returns a slice for individual input arguments
func ToSlice[T any](s ...T) []T {
	if s == nil {
		return []T{}
	}
	return s
}

// ToPtrSlice returns a slice of pointers point to the input slice's elements
func ToPtrSlice[T any, S ~[]T](s S) []*T {
	result := make([]*T, len(s))
	for i := range s {
		result[i] = &s[i]
	}
	return result
}
