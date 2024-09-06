package gofn

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
// When the `iterFunc` function returns false, the iteration stops.
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

// IterReversePtr iterates over pointers to items from multiple slices with ability to stop.
// When the `iterFunc` function returns false, the iteration stops.
func IterReversePtr[T any, S ~[]T](iterFunc func(index int, v *T) bool, slices ...S) {
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
