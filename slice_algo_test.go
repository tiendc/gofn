package gofn

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Compact(t *testing.T) {
	assert.Equal(t, []int{1, -1}, Compact([]int{1, 0, -1}))
	assert.Equal(t, []bool{true, true, true}, Compact([]bool{true, true, false, false, true}))
	assert.Equal(t, []string{"1", "2"}, Compact([]string{"1", "", "2"}))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, []St{{1, "1"}, {2, "2"}}, Compact([]St{{1, "1"}, {}, {2, "2"}}))
}

func Test_Drop(t *testing.T) {
	// Nil/empty source slice
	assert.Equal(t, []int{}, Drop([]int(nil)))
	assert.Equal(t, []int{}, Drop([]int{}))

	assert.Equal(t, []int{1, 4}, Drop([]int{1, 2, 3, 4, 5}, 5, 3, 2, 7))
}

func Test_Reverse(t *testing.T) {
	assert.Equal(t, []int{}, Reverse([]int{}))
	assert.Equal(t, []int64{1}, Reverse([]int64{1}))
	assert.Equal(t, []int{3, 2, 1}, Reverse([]int{1, 2, 3}))

	s := []int{-1, -2, 0, 1, 2}
	Reverse(s)
	assert.Equal(t, []int{2, 1, 0, -2, -1}, s)
}

func Test_ReverseCopy(t *testing.T) {
	assert.Equal(t, []int{}, ReverseCopy([]int{}))
	assert.Equal(t, []int64{1}, ReverseCopy([]int64{1}))
	assert.Equal(t, []int{3, 2, 1}, ReverseCopy([]int{1, 2, 3}))

	s := []int{-1, -2, 0, 1, 2}
	s2 := ReverseCopy(s)
	assert.Equal(t, []int{-1, -2, 0, 1, 2}, s)
	assert.Equal(t, []int{2, 1, 0, -2, -1}, s2)
}

func Test_Shuffle(t *testing.T) {
	// Empty input slice
	s1 := Shuffle[int]([]int(nil))
	assert.Equal(t, []int{}, s1)
	// One item input
	s2 := Shuffle([]float32{1.1})
	assert.Equal(t, []float32{1.1}, s2)
	// Multiple items input (with using custom rand function)
	s3 := Shuffle([]string{"1", "2", "3"}, rand.Intn)
	Sort(s3)
	assert.Equal(t, []string{"1", "2", "3"}, s3)
}

func Test_Chunk(t *testing.T) {
	// Empty input
	chunks := Chunk([]int{}, 5)
	assert.True(t, len(chunks) == 0)

	// Nil input
	chunks = Chunk[int]([]int(nil), 5)
	assert.True(t, len(chunks) == 0)

	// Chunk size greater than input size
	chunks = Chunk([]int{1, 2, 3}, 5)
	assert.True(t, len(chunks) == 1 && reflect.DeepEqual(chunks[0], []int{1, 2, 3}))

	// Normal case
	chunks = Chunk([]int{1, 2, 3, 4, 5}, 2)
	assert.True(t, len(chunks) == 3 &&
		len(chunks[0]) == 2 && reflect.DeepEqual(chunks[0], []int{1, 2}) &&
		len(chunks[1]) == 2 && reflect.DeepEqual(chunks[1], []int{3, 4}) &&
		len(chunks[2]) == 1 && reflect.DeepEqual(chunks[2], []int{5}))
}

func Test_ChunkByPieces(t *testing.T) {
	// Empty input
	chunks := ChunkByPieces([]int{}, 5)
	assert.True(t, len(chunks) == 0)

	// Nil input
	chunks = ChunkByPieces[int]([]int(nil), 5)
	assert.True(t, len(chunks) == 0)

	// Chunk count is zero
	chunks = ChunkByPieces([]int{1, 2, 3}, 0)
	assert.Equal(t, [][]int{}, chunks)

	// Chunk count greater than input size
	chunks = ChunkByPieces([]int{1, 2, 3}, 5)
	assert.True(t, len(chunks) == 3 &&
		len(chunks[0]) == 1 && reflect.DeepEqual(chunks[0], []int{1}) &&
		len(chunks[1]) == 1 && reflect.DeepEqual(chunks[1], []int{2}) &&
		len(chunks[2]) == 1 && reflect.DeepEqual(chunks[2], []int{3}))

	// Normal case
	chunks = ChunkByPieces([]int{1, 2, 3, 4, 5}, 2)
	assert.True(t, len(chunks) == 2 &&
		len(chunks[0]) == 3 && reflect.DeepEqual(chunks[0], []int{1, 2, 3}) &&
		len(chunks[1]) == 2 && reflect.DeepEqual(chunks[1], []int{4, 5}))
}

func Test_Reduce(t *testing.T) {
	// Empty slice
	assert.Equal(t, 0, Reduce[int]([]int{}, func(acc, v int) int { return acc + v }))
	// Slice has 1 element
	assert.Equal(t, 1, Reduce[int]([]int{1}, func(acc, v int) int { return acc + v }))

	assert.Equal(t, 6, Reduce[int]([]int{1, 2, 3}, func(acc, v int) int { return acc + v }))
	assert.Equal(t, 8, Reduce[int]([]int{1, 2, 4}, func(acc, v int) int { return acc * v }))
	assert.Equal(t, 0, Reduce[int]([]int{1, 2, 0}, func(acc, v int) int { return acc * v }))
}

func Test_ReduceEx(t *testing.T) {
	assert.Equal(t, 7, ReduceEx[int]([]int{1, 2, 3}, func(acc, v, i int) int { return acc + v }, 1))
	assert.Equal(t, 8, ReduceEx[int]([]int{1, 2, 4}, func(acc, v, i int) int { return acc * v }, 1))
	assert.Equal(t, 0, ReduceEx[int]([]int{1, 2, 0}, func(acc, v, i int) int { return acc * v }, 1))
}

func Test_ReduceRight(t *testing.T) {
	// Empty slice
	assert.Equal(t, 0, ReduceRight[int]([]int{}, func(acc, v int) int { return acc + v }))
	// Slice has 1 element
	assert.Equal(t, 1, ReduceRight[int]([]int{1}, func(acc, v int) int { return acc + v }))

	assert.Equal(t, 6, ReduceRight[int]([]int{1, 2, 3}, func(acc, v int) int { return acc + v }))
	assert.Equal(t, 8, ReduceRight[int]([]int{1, 2, 4}, func(acc, v int) int { return acc * v }))
	assert.Equal(t, 0, ReduceRight[int]([]int{1, 2, 0}, func(acc, v int) int { return acc * v }))
}

func Test_ReduceRightEx(t *testing.T) {
	assert.Equal(t, 7, ReduceRightEx[int]([]int{1, 2, 3}, func(acc, v, i int) int { return acc + v }, 1))
	assert.Equal(t, 8, ReduceRightEx[int]([]int{1, 2, 4}, func(acc, v, i int) int { return acc * v }, 1))
	assert.Equal(t, 0, ReduceRightEx[int]([]int{1, 2, 0}, func(acc, v, i int) int { return acc * v }, 1))
}

func Test_Partition(t *testing.T) {
	// Nil/empty input
	s1, s2 := Partition([]int(nil), func(i int, _ int) bool { return i%2 == 0 })
	assert.Equal(t, []int{}, s1)
	assert.Equal(t, []int{}, s2)
	s1, s2 = Partition([]int{}, func(i int, _ int) bool { return i%2 == 0 })
	assert.Equal(t, []int{}, s1)
	assert.Equal(t, []int{}, s2)

	s1, s2 = Partition([]int{3, 5, 7, 3}, func(i int, _ int) bool { return i%2 == 1 })
	assert.Equal(t, []int{3, 5, 7, 3}, s1)
	assert.Equal(t, []int{}, s2)

	s1, s2 = Partition([]int{3, 5, 7, 3}, func(i int, _ int) bool { return i%2 == 0 })
	assert.Equal(t, []int{}, s1)
	assert.Equal(t, []int{3, 5, 7, 3}, s2)

	s1, s2 = Partition([]int{3, 2, 5, 0, 0, 7, 3, 4}, func(i int, _ int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2, 0, 0, 4}, s1)
	assert.Equal(t, []int{3, 5, 7, 3}, s2)
}

func Test_PartitionN(t *testing.T) {
	// Zero partition
	p := PartitionN([]int{1, 2, 3}, 0, func(i int, _ int) int { return i % 2 })
	assert.Equal(t, 0, len(p))

	// Nil/empty input
	p = PartitionN([]int(nil), 3, func(i int, _ int) int { return i % 2 })
	assert.Equal(t, []int{}, p[0])
	assert.Equal(t, []int{}, p[1])
	assert.Equal(t, []int{}, p[2])
	p = PartitionN([]int{}, 3, func(i int, _ int) int { return i % 2 })
	assert.Equal(t, []int{}, p[0])
	assert.Equal(t, []int{}, p[1])
	assert.Equal(t, []int{}, p[2])

	p = PartitionN([]int{20, 30, 40, 50, 5, -1, 15, 10}, 3, func(i int, _ int) int { return i / 10 })
	assert.Equal(t, []int{5, -1}, p[0])
	assert.Equal(t, []int{15, 10}, p[1])
	assert.Equal(t, []int{20, 30, 40, 50}, p[2])
}

func Test_Union(t *testing.T) {
	assert.Equal(t, []int{}, Union[int]([]int(nil), nil))
	assert.Equal(t, []int{}, Union(nil, []int{}))
	assert.Equal(t, []int{1}, Union([]int{1}, nil))

	assert.Equal(t, []int{1, 2, 3, 4}, Union([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{1, 2, 3, 4}, Union([]int{1, 2, 3, 2}, []int{1, 2, 4, 3}))
	assert.Equal(t, []string{"1", "2", "3", "4"}, Union([]string{"1", "2", "3", "2"}, []string{"1", "2", "4", "3"}))
}

// nolint: forcetypeassert
func Test_UnionBy(t *testing.T) {
	assert.Equal(t, []any{}, UnionBy[any]([]any(nil), nil, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, UnionBy(nil, []any{}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1}, UnionBy([]any{1}, nil, func(t any) int { return t.(int) }))

	assert.Equal(t, []any{1, 2, 3, 4}, UnionBy([]any{1, 2}, []any{3, 4},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1, 2, 3, 4}, UnionBy([]any{1, 2, 3, 2}, []any{1, 2, 4, 3},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{"1", "2", "3", "4"},
		UnionBy([]any{"1", "2", "3", "2"}, []any{"1", "2", "4", "3"},
			func(t any) string { return t.(string) }))
}

// nolint: forcetypeassert
func Test_UnionPred_Deprecated(t *testing.T) {
	assert.Equal(t, []any{1, 2, 3, 4}, UnionPred([]any{1, 2}, []any{3, 4},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1, 2, 3, 4}, UnionPred([]any{1, 2, 3, 2}, []any{1, 2, 4, 3},
		func(t any) int { return t.(int) }))
}

func Test_Intersection(t *testing.T) {
	assert.Equal(t, []int{}, Intersection[int]([]int(nil), nil))
	assert.Equal(t, []int{}, Intersection(nil, []int{}))
	assert.Equal(t, []int{}, Intersection([]int{1}, nil))
	assert.Equal(t, []int{}, Intersection(nil, []int{1}))

	assert.Equal(t, []int{}, Intersection([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, []int{1, 2, 3}, Intersection([]int{1, 2, 3, 2}, []int{1, 2, 4, 3}))
	assert.Equal(t, []string{"1", "2", "3"}, Intersection([]string{"1", "2", "3", "2"}, []string{"1", "2", "4", "3"}))
}

// nolint: forcetypeassert
func Test_IntersectionBy(t *testing.T) {
	assert.Equal(t, []any{}, IntersectionBy[any]([]any(nil), nil, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, IntersectionBy(nil, []any{}, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, IntersectionBy([]any{1}, nil, func(t any) int { return t.(int) }))
	assert.Equal(t, []any{}, IntersectionBy(nil, []any{1}, func(t any) int { return t.(int) }))

	assert.Equal(t, []any{}, IntersectionBy([]any{1, 2}, []any{3, 4},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1, 2, 3}, IntersectionBy([]any{1, 2, 3, 2}, []any{1, 2, 4, 3},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{"1", "2", "3"},
		IntersectionBy([]any{"1", "2", "3", "2"}, []any{"1", "2", "4", "3"},
			func(t any) string { return t.(string) }))
}

// nolint: forcetypeassert
func Test_IntersectionPred_Deprecated(t *testing.T) {
	assert.Equal(t, []any{}, IntersectionPred([]any{1, 2}, []any{3, 4},
		func(t any) int { return t.(int) }))
	assert.Equal(t, []any{1, 2, 3}, IntersectionPred([]any{1, 2, 3, 2}, []any{1, 2, 4, 3},
		func(t any) int { return t.(int) }))
}

func Test_Difference(t *testing.T) {
	l, r := Difference[int]([]int(nil), nil)
	assert.Equal(t, []int{}, l)
	assert.Equal(t, []int{}, r)
	l, r = Difference([]int{}, nil)
	assert.Equal(t, []int{}, l)
	assert.Equal(t, []int{}, r)

	l, r = Difference([]int{1, 2}, []int{3, 4})
	assert.Equal(t, []int{1, 2}, l)
	assert.Equal(t, []int{3, 4}, r)

	l, r = Difference([]int{1, 2, 3, 2}, []int{1, 4, 4, 3})
	assert.Equal(t, []int{2, 2}, l)
	assert.Equal(t, []int{4, 4}, r)

	l2, r2 := Difference([]string{"1", "2", "3", "2"}, []string{"1", "4", "2", "", "3"})
	assert.Equal(t, []string{}, l2)
	assert.Equal(t, []string{"4", ""}, r2)
}

// nolint: forcetypeassert
func Test_DifferenceBy(t *testing.T) {
	l, r := DifferenceBy[any]([]any(nil), nil, func(t any) int { return t.(int) })
	assert.Equal(t, []any{}, l)
	assert.Equal(t, []any{}, r)
	l, r = DifferenceBy([]any{}, nil, func(t any) int { return t.(int) })
	assert.Equal(t, []any{}, l)
	assert.Equal(t, []any{}, r)

	l, r = DifferenceBy([]any{1, 2}, []any{3, 4}, func(t any) int { return t.(int) })
	assert.Equal(t, []any{1, 2}, l)
	assert.Equal(t, []any{3, 4}, r)

	l, r = DifferenceBy([]any{1, 2, 3, 2}, []any{1, 4, 4, 3}, func(t any) int { return t.(int) })
	assert.Equal(t, []any{2, 2}, l)
	assert.Equal(t, []any{4, 4}, r)

	l, r = DifferenceBy([]any{"1", "2", "3", "2"}, []any{"1", "4", "2", "", "3"},
		func(t any) string { return t.(string) })
	assert.Equal(t, []any{}, l)
	assert.Equal(t, []any{"4", ""}, r)
}

// nolint: forcetypeassert
func Test_DifferencePred_Deprecated(t *testing.T) {
	l, r := DifferencePred([]any{1, 2}, []any{3, 4}, func(t any) int { return t.(int) })
	assert.Equal(t, []any{1, 2}, l)
	assert.Equal(t, []any{3, 4}, r)

	l, r = DifferencePred([]any{1, 2, 3, 2}, []any{1, 4, 4, 3}, func(t any) int { return t.(int) })
	assert.Equal(t, []any{2, 2}, l)
	assert.Equal(t, []any{4, 4}, r)
}

func Test_Flatten(t *testing.T) {
	assert.True(t, reflect.DeepEqual([]int{}, Flatten([]int{}, []int{})))
	assert.True(t, reflect.DeepEqual([]int{1}, Flatten([]int{}, []int{1})))
	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, Flatten([]int{1, 2}, []int{3})))
	assert.True(t, reflect.DeepEqual([]int64{1, 2, 3, 4, 0}, Flatten([]int64{1, 2}, []int64{3, 4}, []int64{0})))
}

func Test_Flatten3(t *testing.T) {
	assert.True(t, reflect.DeepEqual([]int{}, Flatten3([][]int{}, [][]int{})))
	assert.True(t, reflect.DeepEqual([]int{1}, Flatten3([][]int{{1}}, [][]int{})))
	assert.True(t, reflect.DeepEqual([]int{1, 2, 3, 4}, Flatten3([][]int{{1, 2}, {3, 4}})))
	assert.True(t, reflect.DeepEqual([]int{1, 2, 3, 4, 0}, Flatten3([][]int{{1, 2}, {3, 4}}, [][]int{{0}})))
}

func Test_Zip(t *testing.T) {
	assert.Equal(t, []*Tuple2[int, bool]{}, Zip([]int{}, []bool{}))
	assert.Equal(t, []*Tuple2[int, bool]{}, Zip([]int{1, 2, 3}, []bool{}))
	assert.Equal(t, []*Tuple2[int, string]{{1, "1"}, {2, "2"}}, Zip([]int{1, 2, 3}, []string{"1", "2"}))
}

func Test_Zip3(t *testing.T) {
	assert.Equal(t, []*Tuple3[int, bool, int32]{}, Zip3([]int{}, []bool{}, []int32{}))
	assert.Equal(t, []*Tuple3[int, bool, int32]{}, Zip3([]int{1, 2, 3}, []bool{}, []int32{4, 5}))
	assert.Equal(t, []*Tuple3[int, int, int]{{1, 11, 111}, {2, 22, 222}},
		Zip3([]int{1, 2, 3, 4}, []int{11, 22, 33}, []int{111, 222}))
	assert.Equal(t, []*Tuple3[int, string, bool]{{1, "1", true}, {2, "2", false}},
		Zip3([]int{1, 2, 3}, []string{"1", "2"}, []bool{true, false, false}))
}
