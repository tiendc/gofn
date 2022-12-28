package gofn

func ForEach[T any](s []T, pred func(i int, t T)) {
	for i := range s {
		pred(i, s[i])
	}
}

func ForEachPtr[T any](s []T, pred func(i int, t *T)) {
	for i := range s {
		pred(i, &s[i])
	}
}

func ForEachReverse[T any](s []T, pred func(i int, t T)) {
	for i := len(s) - 1; i >= 0; i-- {
		pred(i, s[i])
	}
}

func ForEachReversePtr[T any](s []T, pred func(i int, t *T)) {
	for i := len(s) - 1; i >= 0; i-- {
		pred(i, &s[i])
	}
}
