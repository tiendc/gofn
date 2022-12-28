package gofn

// ToIntfSlice convert a slice to a slice of interface
func ToIntfSlice[T any](s []T) []interface{} {
	result := make([]interface{}, len(s))
	for i := range s {
		result[i] = s[i]
	}
	return result
}

// ToStringSlice converts str-approximate slice to string slice
func ToStringSlice[U, T ~string](slice []T) []U {
	result := make([]U, len(slice))
	for i := range slice {
		result[i] = U(slice[i])
	}
	return result
}

// ToNumberSlice converts int-approximate slice to int slice
func ToNumberSlice[U, T NumberEx](slice []T) []U {
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
