package gofn

import "math/rand"

var (
	StrLowerAlpha   = []rune("abcdefghijklmnopqrstuvwxyz")
	StrUpperAlpha   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	StrDigits       = []rune("0123456789")
	StrDefaultChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

// RandChoiceMaker a struct for picking up items randomly from a list of items
type RandChoiceMaker[T any] struct {
	source   []*T // Use pointers to slice items to gain more performance when item type is struct
	intNFunc func(n int) int
}

// Next gets the next item randomly from the source
// If there is no item in the source, returns false as the 2nd value
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

// Shuffle items of a slice and returns a new slice
func Shuffle[T any, S ~[]T](s S, randFuncs ...func(n int) int) S {
	if len(s) <= 1 {
		return append(S{}, s...)
	}
	maker := NewRandChoiceMaker(s, randFuncs...)
	result := make(S, 0, len(s))
	for {
		item, valid := maker.Next()
		if !valid {
			break
		}
		result = append(result, item)
	}
	return result
}

// RandString generates a random string
func RandString(n int) string {
	return RandStringEx(n, StrDefaultChars)
}

// RandStringEx generates a random string
func RandStringEx[S ~[]rune](n int, allowedChars S) string {
	b := make([]rune, n)
	numChars := len(allowedChars)
	for i := range b {
		b[i] = allowedChars[rand.Intn(numChars)] // nolint: gosec
	}
	return string(b)
}
