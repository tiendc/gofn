package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ForEach(t *testing.T) {
	copy1 := []int{}
	ForEach[int]([]int(nil), func(i int, t int) { copy1 = append(copy1, t) })
	assert.Equal(t, []int{}, copy1)

	copy2 := []int64{}
	ForEach([]int64{1, 2, 3}, func(i int, t int64) { copy2 = append(copy2, t) })
	assert.Equal(t, []int64{1, 2, 3}, copy2)
}

func Test_ForEachPtr(t *testing.T) {
	copy1 := []int{}
	ForEachPtr[int]([]int(nil), func(i int, t *int) { copy1 = append(copy1, *t) })
	assert.Equal(t, []int{}, copy1)

	copy2 := []int64{}
	ForEachPtr([]int64{1, 2, 3}, func(i int, t *int64) { copy2 = append(copy2, *t) })
	assert.Equal(t, []int64{1, 2, 3}, copy2)
}

func Test_ForEachReverse(t *testing.T) {
	copy1 := []int{}
	ForEachReverse[int]([]int(nil), func(i int, t int) { copy1 = append(copy1, t) })
	assert.Equal(t, []int{}, copy1)

	copy2 := []int64{}
	ForEachReverse([]int64{1, 2, 3}, func(i int, t int64) { copy2 = append(copy2, t) })
	assert.Equal(t, []int64{3, 2, 1}, copy2)
}

func Test_ForEachPtrReverse(t *testing.T) {
	copy1 := []int{}
	ForEachPtrReverse[int]([]int(nil), func(i int, t *int) { copy1 = append(copy1, *t) })
	assert.Equal(t, []int{}, copy1)

	copy2 := []int64{}
	ForEachPtrReverse([]int64{1, 2, 3}, func(i int, t *int64) { copy2 = append(copy2, *t) })
	assert.Equal(t, []int64{3, 2, 1}, copy2)
}

func Test_Iter(t *testing.T) {
	// No slice input
	copy1 := []int{}
	Iter[int, []int](func(i int, t int) bool { copy1 = append(copy1, t); return true })
	assert.Equal(t, []int{}, copy1)

	// Empty slice input
	copy2 := []int{}
	Iter(func(i int, t int) bool { copy2 = append(copy2, t); return true }, []int{})
	assert.Equal(t, []int{}, copy2)

	// Single slice
	copy3 := []int{}
	index3 := []int{}
	Iter(func(i int, t int) bool { copy3 = append(copy3, t); index3 = append(index3, i); return true },
		[]int{1, 2, 3})
	assert.Equal(t, []int{1, 2, 3}, copy3)
	assert.Equal(t, []int{0, 1, 2}, index3)

	// Multiple slices
	copy4 := []int32{}
	index4 := []int{}
	Iter(func(i int, t int32) bool { copy4 = append(copy4, t); index4 = append(index4, i); return true },
		[]int32{1, 2, 3}, []int32{}, []int32{4, 5})
	assert.Equal(t, []int32{1, 2, 3, 4, 5}, copy4)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, index4)

	// With stop
	copy5 := []string{}
	index5 := []int{}
	Iter(func(i int, t string) bool { copy5 = append(copy5, t); index5 = append(index5, i); return i < 3 },
		[]string{"1", "2", "3"}, []string{}, []string{"4", "5"})
	assert.Equal(t, []string{"1", "2", "3", "4"}, copy5)
	assert.Equal(t, []int{0, 1, 2, 3}, index5)
}

func Test_IterPtr(t *testing.T) {
	// No slice input
	copy1 := []int{}
	IterPtr[int, []int](func(i int, t *int) bool { copy1 = append(copy1, *t); return true })
	assert.Equal(t, []int{}, copy1)

	// Empty slice input
	copy2 := []int{}
	IterPtr(func(i int, t *int) bool { copy2 = append(copy2, *t); return true }, []int{})
	assert.Equal(t, []int{}, copy2)

	// Single slice
	copy3 := []int{}
	index3 := []int{}
	IterPtr(func(i int, t *int) bool { copy3 = append(copy3, *t); index3 = append(index3, i); return true },
		[]int{1, 2, 3})
	assert.Equal(t, []int{1, 2, 3}, copy3)
	assert.Equal(t, []int{0, 1, 2}, index3)

	// Multiple slices
	copy4 := []int32{}
	index4 := []int{}
	IterPtr(func(i int, t *int32) bool { copy4 = append(copy4, *t); index4 = append(index4, i); return true },
		[]int32{1, 2, 3}, []int32{}, []int32{4, 5})
	assert.Equal(t, []int32{1, 2, 3, 4, 5}, copy4)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, index4)

	// With stop
	copy5 := []string{}
	index5 := []int{}
	IterPtr(func(i int, t *string) bool { copy5 = append(copy5, *t); index5 = append(index5, i); return i < 3 },
		[]string{"1", "2", "3"}, []string{}, []string{"4", "5"})
	assert.Equal(t, []string{"1", "2", "3", "4"}, copy5)
	assert.Equal(t, []int{0, 1, 2, 3}, index5)
}

func Test_IterReverse(t *testing.T) {
	// No slice input
	copy1 := []int{}
	IterReverse[int, []int](func(i int, t int) bool { copy1 = append(copy1, t); return true })
	assert.Equal(t, []int{}, copy1)

	// Empty slice input
	copy2 := []int{}
	IterReverse(func(i int, t int) bool { copy2 = append(copy2, t); return true }, []int{})
	assert.Equal(t, []int{}, copy2)

	// Single slice
	copy3 := []int{}
	index3 := []int{}
	IterReverse(func(i int, t int) bool { copy3 = append(copy3, t); index3 = append(index3, i); return true },
		[]int{1, 2, 3})
	assert.Equal(t, []int{3, 2, 1}, copy3)
	assert.Equal(t, []int{2, 1, 0}, index3)

	// Multiple slices
	copy4 := []int32{}
	index4 := []int{}
	IterReverse(func(i int, t int32) bool { copy4 = append(copy4, t); index4 = append(index4, i); return true },
		[]int32{1, 2, 3}, []int32{}, []int32{4, 5})
	assert.Equal(t, []int32{5, 4, 3, 2, 1}, copy4)
	assert.Equal(t, []int{4, 3, 2, 1, 0}, index4)

	// With stop
	copy5 := []string{}
	index5 := []int{}
	IterReverse(func(i int, t string) bool { copy5 = append(copy5, t); index5 = append(index5, i); return i > 3 },
		[]string{"1", "2", "3"}, []string{}, []string{"4", "5"})
	assert.Equal(t, []string{"5", "4"}, copy5)
	assert.Equal(t, []int{4, 3}, index5)
}

func Test_IterPtrReverse(t *testing.T) {
	// No slice input
	copy1 := []int{}
	IterPtrReverse[int, []int](func(i int, t *int) bool { copy1 = append(copy1, *t); return true })
	assert.Equal(t, []int{}, copy1)

	// Empty slice input
	copy2 := []int{}
	IterPtrReverse(func(i int, t *int) bool { copy2 = append(copy2, *t); return true }, []int{})
	assert.Equal(t, []int{}, copy2)

	// Single slice
	copy3 := []int{}
	index3 := []int{}
	IterPtrReverse(func(i int, t *int) bool { copy3 = append(copy3, *t); index3 = append(index3, i); return true },
		[]int{1, 2, 3})
	assert.Equal(t, []int{3, 2, 1}, copy3)
	assert.Equal(t, []int{2, 1, 0}, index3)

	// Multiple slices
	copy4 := []int32{}
	index4 := []int{}
	IterPtrReverse(func(i int, t *int32) bool { copy4 = append(copy4, *t); index4 = append(index4, i); return true },
		[]int32{1, 2, 3}, []int32{}, []int32{4, 5})
	assert.Equal(t, []int32{5, 4, 3, 2, 1}, copy4)
	assert.Equal(t, []int{4, 3, 2, 1, 0}, index4)

	// With stop
	copy5 := []string{}
	index5 := []int{}
	IterPtrReverse(func(i int, t *string) bool { copy5 = append(copy5, *t); index5 = append(index5, i); return i > 3 },
		[]string{"1", "2", "3"}, []string{}, []string{"4", "5"})
	assert.Equal(t, []string{"5", "4"}, copy5)
	assert.Equal(t, []int{4, 3}, index5)
}
