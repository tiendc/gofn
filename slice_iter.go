package gofn

// ForEach iterates over slice items.
// For more advanced requirements, see Iter.
func ForEach[T any, S ~[]T](s S, pred func(i int, t T)) {
	for i := range s {
		pred(i, s[i])
	}
}

// ForEachPtr iterates over pointers to slice items.
// You can use this to achieve more performance if you have a slice of big structs.
// For more advanced requirements, see IterPtr.
func ForEachPtr[T any, S ~[]T](s S, pred func(i int, t *T)) {
	for i := range s {
		pred(i, &s[i])
	}
}

// ForEachReverse iterates over slice items from the end.
// For more advanced requirements, see IterReverse.
func ForEachReverse[T any, S ~[]T](s S, pred func(i int, t T)) {
	for i := len(s) - 1; i >= 0; i-- {
		pred(i, s[i])
	}
}

// ForEachPtrReverse iterates over pointers to slice items from the end.
// For more advanced requirements, see IterPtrReverse.
func ForEachPtrReverse[T any, S ~[]T](s S, pred func(i int, t *T)) {
	for i := len(s) - 1; i >= 0; i-- {
		pred(i, &s[i])
	}
}

// Iter iterates over items from multiple slices with ability to stop.
// When the `iterFunc` function returns false, the iteration stops.
func Iter[T any, S ~[]T](iterFunc func(index int, v T) bool, slices ...S) {
	global := 0
	for _, s := range slices {
		for _, v := range s {
			if !iterFunc(global, v) {
				return
			}
			global++
		}
	}
}

// IterPtr iterates over pointers to items from multiple slices with ability to stop.
// When the `iterFunc` function returns false, the iterating stops.
func IterPtr[T any, S ~[]T](iterFunc func(index int, v *T) bool, slices ...S) {
	global := 0
	for _, s := range slices {
		for j := range s {
			if !iterFunc(global, &s[j]) {
				return
			}
			global++
		}
	}
}

// IterReverse iterates over items from multiple slices from the end with ability to stop.
// When the `iterFunc` function returns false, the iteration stops.
func IterReverse[T any, S ~[]T](iterFunc func(index int, v T) bool, slices ...S) {
	global := -1
	for _, s := range slices {
		global += len(s)
	}
	for i := len(slices) - 1; i >= 0; i-- {
		s := slices[i]
		for j := len(s) - 1; j >= 0; j-- {
			if !iterFunc(global, s[j]) {
				return
			}
			global--
		}
	}
}

// IterPtrReverse iterates over pointers to items from multiple slices with ability to stop.
// When the `iterFunc` function returns false, the iteration stops.
func IterPtrReverse[T any, S ~[]T](iterFunc func(index int, v *T) bool, slices ...S) {
	global := -1
	for _, s := range slices {
		global += len(s)
	}
	for i := len(slices) - 1; i >= 0; i-- {
		s := slices[i]
		for j := len(s) - 1; j >= 0; j-- {
			if !iterFunc(global, &s[j]) {
				return
			}
			global--
		}
	}
}
