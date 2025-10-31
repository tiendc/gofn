package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Filter(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, Filter[int]([]int(nil), func(v int) bool { return true }))
	assert.Equal(t, []float32{}, Filter([]float32{}, func(v float32) bool { return true }))

	assert.Equal(t, []int{-4, 0},
		Filter([]int{2, -4, 6, 0}, func(v int) bool { return v <= 0 }))
	assert.Equal(t, []float32{2, 6, 0.0001},
		Filter([]float32{2, -4, 6, 0.0001}, func(v float32) bool { return v > 0 }))
	assert.Equal(t, []string{"one", "two"},
		Filter([]string{"one", "", "two"}, func(v string) bool { return len(v) > 0 }))
}

func Test_FilterPtr(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterPtr[int]([]int(nil), func(v *int) bool { return true }))
	assert.Equal(t, []float32{}, FilterPtr([]float32{}, func(v *float32) bool { return true }))

	assert.Equal(t, []int{-4, 0},
		FilterPtr([]int{2, -4, 6, 0}, func(v *int) bool { return *v <= 0 }))
	assert.Equal(t, []float32{2, 6, 0.0001},
		FilterPtr([]float32{2, -4, 6, 0.0001}, func(v *float32) bool { return *v > 0 }))
	assert.Equal(t, []string{"one", "two"},
		FilterPtr([]string{"one", "", "two"}, func(v *string) bool { return len(*v) > 0 }))
}

func Test_FilterLT(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterLT[int]([]int(nil), 0))
	assert.Equal(t, []int{}, FilterLT([]int{}, 0))

	assert.Equal(t, []int{-10, -3}, FilterLT([]int{0, -10, 1, -3, 0, 7}, 0))
	assert.Equal(t, []string{"Ab"}, FilterLT([]string{"Ab", "cd", "a"}, "a"))
}

func Test_FilterLTE(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterLTE[int]([]int(nil), 0))
	assert.Equal(t, []int{}, FilterLTE([]int{}, 0))

	assert.Equal(t, []int{0, -10, -3, 0}, FilterLTE([]int{0, -10, 1, -3, 0, 7}, 0))
	assert.Equal(t, []string{"Ab", "a"}, FilterLTE([]string{"Ab", "cd", "a"}, "a"))
}

func Test_FilterGT(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterGT[int]([]int(nil), 0))
	assert.Equal(t, []int{}, FilterGT([]int{}, 0))

	assert.Equal(t, []int{1, 7}, FilterGT([]int{0, -10, 1, -3, 0, 7}, 0))
	assert.Equal(t, []string{"cd"}, FilterGT([]string{"Ab", "cd", "a"}, "a"))
}

func Test_FilterGTE(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterGTE[int]([]int(nil), 0))
	assert.Equal(t, []int{}, FilterGTE([]int{}, 0))

	assert.Equal(t, []int{0, 1, 0, 7}, FilterGTE([]int{0, -10, 1, -3, 0, 7}, 0))
	assert.Equal(t, []string{"cd", "a"}, FilterGTE([]string{"Ab", "cd", "a"}, "a"))
}

func Test_FilterRange(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterRange[int]([]int(nil), 0, 0))
	assert.Equal(t, []int{}, FilterRange([]int{}, 0, 0))

	assert.Equal(t, []int{1, 2}, FilterRange([]int{0, -10, 1, 2, -3, 0, 7}, 1, 5))
	assert.Equal(t, []string{"a", "b", "c"}, FilterRange([]string{"Ab", "cd", "a", "b", "", "c"}, "a", "c"))
}

func Test_FilterNE(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterNE[int]([]int(nil), 0))
	assert.Equal(t, []int{}, FilterNE([]int{}, 0))

	assert.Equal(t, []int{-10, 1, -3, 7}, FilterNE([]int{0, -10, 1, -3, 0, 7}, 0))
	assert.Equal(t, []string{"Ab", "cd"}, FilterNE([]string{"Ab", "cd", "a"}, "a"))
}

func Test_FilterIN(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterIN[int]([]int(nil), 0))
	assert.Equal(t, []int{}, FilterIN([]int{}, 0))

	assert.Equal(t, []int{1, 7}, FilterIN([]int{0, -10, 1, -3, 0, 7}, 1, 3, 7))
	assert.Equal(t, []string{"a"}, FilterIN([]string{"Ab", "cd", "a"}, "a", "b"))
}

func Test_FilterNIN(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []int{}, FilterNIN[int]([]int(nil), 0))
	assert.Equal(t, []int{}, FilterNIN([]int{}, 0))

	assert.Equal(t, []int{-10, 1, -3}, FilterNIN([]int{0, -10, 1, -3, 0, 7}, 3, 7, 0))
	assert.Equal(t, []string{"Ab"}, FilterNIN([]string{"Ab", "cd", "a"}, "a", "b", "cd"))
}

func Test_FilterLIKE(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []string{}, FilterLIKE[string]([]string(nil), ""))
	assert.Equal(t, []string{}, FilterLIKE([]string{}, ""))

	assert.Equal(t, []string{},
		FilterLIKE([]string{"*Abc*", "*abc*", "abc*", "*abc"}, "ac"))
	assert.Equal(t, []string{"*abc*", "abc*", "*abc"},
		FilterLIKE([]string{"*Abc*", "*abc*", "abc*", "*abc"}, "abc"))
	assert.Equal(t, []string{"*Abc*"},
		FilterLIKE([]string{"*Abc*", "*abc*", "abc*", "*abc"}, "Abc"))

	type X string
	assert.Equal(t, []X{},
		FilterLIKE([]X{"*Abc*", "*abc*", "abc*", "*abc"}, "ac"))
	assert.Equal(t, []X{"*abc*", "abc*", "*abc"},
		FilterLIKE([]X{"*Abc*", "*abc*", "abc*", "*abc"}, "abc"))
	assert.Equal(t, []X{"*Abc*"},
		FilterLIKE([]X{"*Abc*", "*abc*", "abc*", "*abc"}, "Abc"))
}

func Test_FilterILIKE(t *testing.T) {
	// Nil/Empty slices
	assert.Equal(t, []string{}, FilterILIKE[string]([]string(nil), ""))
	assert.Equal(t, []string{}, FilterILIKE([]string{}, ""))

	assert.Equal(t, []string{},
		FilterILIKE([]string{"*Abc*", "*abc*", "abc*", "*abc"}, "aC"))
	assert.Equal(t, []string{"*Abc*", "*abc*", "abc*", "*abc"},
		FilterILIKE([]string{"*Abc*", "*abc*", "abc*", "*abc"}, "abc"))
	assert.Equal(t, []string{"*Abc*", "*abc*", "abc*", "*abc"},
		FilterILIKE([]string{"*Abc*", "*abc*", "abc*", "*abc"}, "AbC"))

	type X string
	assert.Equal(t, []X{},
		FilterILIKE([]X{"*Abc*", "*abc*", "abc*", "*abc"}, "aC"))
	assert.Equal(t, []X{"*Abc*", "*abc*", "abc*", "*abc"},
		FilterILIKE([]X{"*Abc*", "*abc*", "abc*", "*abc"}, "abc"))
	assert.Equal(t, []X{"*Abc*", "*abc*", "abc*", "*abc"},
		FilterILIKE([]X{"*Abc*", "*abc*", "abc*", "*abc"}, "AbC"))
}
