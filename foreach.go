package gofn

func ForEach[T any, S ~[]T](s S, pred func(i int, t T)) {
	for i := range s {
		pred(i, s[i])
	}
}

func ForEachPtr[T any, S ~[]T](s S, pred func(i int, t *T)) {
	for i := range s {
		pred(i, &s[i])
	}
}

func ForEachReverse[T any, S ~[]T](s S, pred func(i int, t T)) {
	for i := len(s) - 1; i >= 0; i-- {
		pred(i, s[i])
	}
}

func ForEachReversePtr[T any, S ~[]T](s S, pred func(i int, t *T)) {
	for i := len(s) - 1; i >= 0; i-- {
		pred(i, &s[i])
	}
}
