package gofn

import "math/rand"

// RandChoiceMaker a struct for picking up items randomly from a list of items
type RandChoiceMaker[T any] struct {
	source   []*T // Use pointers to slice items to gain more performance when item type is struct
	intNFunc func(n int) int
}

// Next gets the next item randomly from the source.
// If there is no item in the source, returns false as the 2nd value.
func (m *RandChoiceMaker[T]) Next() (T, bool) {
	count := len(m.source)
	if count == 0 {
		var defaultVal T
		return defaultVal, false
	}
	index := m.intNFunc(count)
	result := m.source[index]
	switch index {
	case 0:
		m.source = m.source[1:]
	case count - 1:
		m.source = m.source[:count-1]
	default:
		m.source = append(m.source[:index], m.source[index+1:]...)
	}
	return *result, true
}

// HasNext checks to see the maker has items to return
func (m *RandChoiceMaker[T]) HasNext() bool {
	return len(m.source) > 0
}

func NewRandChoiceMaker[T any, S ~[]T](s S, randFuncs ...func(n int) int) RandChoiceMaker[T] {
	var source []*T
	for i := range s {
		source = append(source, &s[i])
	}
	var randFunc func(int) int
	if len(randFuncs) > 0 {
		randFunc = randFuncs[0]
	} else {
		randFunc = rand.Intn
	}
	return RandChoiceMaker[T]{source: source, intNFunc: randFunc}
}

// RandChoice picks up an item randomly from a slice
func RandChoice[T any](s ...T) (T, bool) {
	if len(s) == 0 {
		var defaultVal T
		return defaultVal, false
	}
	return s[rand.Intn(len(s))], true //nolint:gosec
}
