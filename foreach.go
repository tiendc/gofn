package gofn

// ForEach iterates over slice items
func ForEach[T any, S ~[]T](s S, pred func(i int, t T)) {
	for i := range s {
		pred(i, s[i])
	}
}

// ForEachPtr iterates over pointers to slice items.
// You can use this to achieve more performance if you have a slice of big structs.
func ForEachPtr[T any, S ~[]T](s S, pred func(i int, t *T)) {
	for i := range s {
		pred(i, &s[i])
	}
}

// ForEachReverse iterates over slice items from the end
func ForEachReverse[T any, S ~[]T](s S, pred func(i int, t T)) {
	for i := len(s) - 1; i >= 0; i-- {
		pred(i, s[i])
	}
}

// ForEachReversePtr iterates over pointers to slice items from the end
func ForEachReversePtr[T any, S ~[]T](s S, pred func(i int, t *T)) {
	for i := len(s) - 1; i >= 0; i-- {
		pred(i, &s[i])
	}
}
