package gofn

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MapEqual(t *testing.T) {
	// Empty maps
	assert.True(t, MapEqual(map[int]bool{}, map[int]bool{}))

	// One is nil, one is empty
	assert.True(t, MapEqual(nil, map[int]int{}))
	assert.False(t, MapEqual(map[int]bool{}, map[int]bool{1: false}))

	assert.True(t, MapEqual(map[int]int{1: 1, 2: 2}, map[int]int{2: 2, 1: 1}))
	assert.False(t, MapEqual(map[int]int{1: 1, 2: 2}, map[int]int{2: 2, 1: 1, 3: 3}))

	type st struct {
		Int int
		Str string
	}
	assert.True(t, MapEqual(map[int]st{1: {1, "1"}, 2: {2, "2"}},
		map[int]st{2: {2, "2"}, 1: {1, "1"}}))
	assert.False(t, MapEqual(map[int]st{1: {1, "1"}, 2: {2, "2"}},
		map[int]st{2: {2, "2"}, 1: {1, "1"}, 3: {3, "3"}}))
}

func Test_MapEqualPred(t *testing.T) {
	// Empty maps
	assert.True(t, MapEqualPred(map[int]bool{}, map[int]bool{},
		func(v1, v2 bool) bool { return v1 == v2 }))

	// One is nil, one is empty
	assert.True(t, MapEqualPred(nil, map[int]int{},
		func(v1, v2 int) bool { return v1 == v2 }))
	assert.False(t, MapEqualPred(map[int]bool{}, map[int]bool{1: false},
		func(v1, v2 bool) bool { return v1 == v2 }))

	assert.True(t, MapEqualPred(map[int]int{1: 1, 2: 2}, map[int]int{2: 2, 1: 1},
		func(v1, v2 int) bool { return v1 == v2 }))
	assert.False(t, MapEqualPred(map[int]int{1: 1, 2: 2}, map[int]int{2: 2, 1: 1, 3: 3},
		func(v1, v2 int) bool { return v1 == v2 }))

	type st struct {
		Int int
		Str string
	}
	assert.True(t, MapEqualPred(map[int]st{1: {1, "1"}, 2: {2, "2"}}, map[int]st{2: {2, "2"}, 1: {1, "1"}},
		func(v1, v2 st) bool { return v1 == v2 }))
	assert.False(t, MapEqualPred(map[int]st{1: {1, "1"}, 2: {2, "2"}}, map[int]st{2: {2, "2"}, 1: {1, "1"}, 3: {3, "3"}},
		func(v1, v2 st) bool { return v1 == v2 }))

	// Value is also a map
	assert.True(t, MapEqualPred(map[int]map[int]int{1: {1: 1}, 2: {2: 2}},
		map[int]map[int]int{1: {1: 1}, 2: {2: 2}},
		func(v1, v2 map[int]int) bool { return MapEqual(v1, v2) }))
	assert.False(t, MapEqualPred(map[int]map[int]int{1: {1: 1}, 2: {2: 2}},
		map[int]map[int]int{1: {1: 1}, 2: {2: 2}, 3: {3: 3}},
		func(v1, v2 map[int]int) bool { return MapEqual(v1, v2) }))
}

func Test_MapContainKeys(t *testing.T) {
	assert.False(t, MapContainKeys(map[bool]int{}, false, true))

	assert.True(t, MapContainKeys(map[int]int{1: 1, 2: 2, 3: 3}, 1, 2))
	assert.True(t, MapContainKeys(map[int]int{1: 1, 2: 2, 3: 3}, 1, 2, 3))
	assert.False(t, MapContainKeys(map[int]int{1: 1, 2: 2, 3: 3}, 1, 2, 3, 4))
}

func Test_MapContainValues(t *testing.T) {
	assert.False(t, MapContainValues(map[int]bool{}, false, true))

	assert.True(t, MapContainValues(map[int]int{1: 1, 2: 2, 3: 3}, 1, 2))
	assert.True(t, MapContainValues(map[int]int{1: 1, 2: 2, 3: 3}, 1, 2, 3))
	assert.False(t, MapContainValues(map[int]int{1: 1, 2: 2, 3: 3}, 1, 2, 3, 4))

	type st struct {
		Int int
		Str string
	}
	assert.True(t, MapContainValues(map[int]st{1: {1, "1"}, 2: {2, "2"}}, st{1, "1"}, st{2, "2"}))
	assert.False(t, MapContainValues(map[int]st{1: {1, "1"}, 2: {2, "2"}}, st{1, "1"}, st{2, "2"}, st{}))
}

func Test_MapKeys(t *testing.T) {
	assert.Equal(t, []int{}, MapKeys[int, bool](nil))
	assert.Equal(t, []int{}, MapKeys(map[int]bool{}))

	assert.Equal(t, []int{1}, MapKeys(map[int]int{1: 11}))

	// NOTE: order of result slice is non-deterministic
	assert.True(t, ContentEqual([]int{1, 2}, MapKeys(map[int]int{1: 11, 2: 22})))
}

func Test_MapValues(t *testing.T) {
	assert.Equal(t, []bool{}, MapValues[int, bool](nil))
	assert.Equal(t, []int{}, MapValues(map[int]int{}))

	assert.Equal(t, []int{11}, MapValues(map[int]int{1: 11}))
	// NOTE: order of result slice is non-deterministic
	assert.True(t, ContentEqual([]int{11, 22}, MapValues(map[int]int{1: 11, 2: 22})))
}

func Test_MapEntries(t *testing.T) {
	assert.Equal(t, []*Tuple2[int, bool]{}, MapEntries[int, bool](nil))
	assert.Equal(t, []*Tuple2[int, int]{}, MapEntries(map[int]int{}))

	assert.Equal(t, []*Tuple2[int, int]{{1, 11}}, MapEntries(map[int]int{1: 11}))
	assert.True(t, ContentEqualPtr([]*Tuple2[int, int]{{1, 11}, {2, 22}}, MapEntries(map[int]int{1: 11, 2: 22})))
}

func Test_MapUpdate(t *testing.T) {
	// Update with a nil map
	assert.Equal(t, map[int]bool{1: false, 2: true},
		MapUpdate(map[int]bool{1: false, 2: true}, nil))
	// Update with an empty map
	assert.Equal(t, map[int]bool{1: false, 2: true},
		MapUpdate(map[int]bool{1: false, 2: true}, map[int]bool{}))

	// Merge 2 maps
	assert.Equal(t, map[int]string{1: "one", 2: "two", 3: "three"},
		MapUpdate(map[int]string{2: "two"}, map[int]string{1: "one", 3: "three"}))
	// Merge 2 maps with override
	assert.Equal(t, map[int]string{1: "one", 2: "TWO", 3: "three"},
		MapUpdate(map[int]string{2: "two"}, map[int]string{1: "one", 2: "TWO", 3: "three"}))
}

func Test_MapGet(t *testing.T) {
	assert.Equal(t, true, MapGet[int, bool](nil, 1, true))
	assert.Equal(t, 11, MapGet(map[int]int{}, 1, 11))

	assert.Equal(t, 11, MapGet(map[int]int{1: 11}, 1, 100))
	assert.Equal(t, 100, MapGet(map[int]int{1: 11, 2: 22}, 3, 100))
}

func Test_MapPop(t *testing.T) {
	assert.Equal(t, true, MapPop[int, bool](nil, 1, true))
	assert.Equal(t, 11, MapPop(map[int]int{}, 1, 11))

	m1 := map[int]int{1: 11}
	assert.True(t, 11 == MapPop(m1, 1, 100) && reflect.DeepEqual(m1, map[int]int{}))
	m2 := map[int]int{1: 11, 2: 22}
	assert.True(t, 100 == MapPop(m2, 3, 100) && reflect.DeepEqual(m2, map[int]int{1: 11, 2: 22}))
}

func Test_MapSetDefault(t *testing.T) {
	assert.Equal(t, true, MapSetDefault[int, bool](nil, 1, true))
	assert.Equal(t, 11, MapSetDefault(map[int]int{}, 1, 11))

	m1 := map[int]int{1: 11}
	assert.True(t, 11 == MapSetDefault(m1, 1, 100) && reflect.DeepEqual(m1, map[int]int{1: 11}))
	m2 := map[int]int{1: 11}
	assert.True(t, 22 == MapSetDefault(m2, 2, 22) && reflect.DeepEqual(m2, map[int]int{1: 11, 2: 22}))
}

func Test_MapUnionKeys(t *testing.T) {
	assert.Equal(t, []int{}, MapUnionKeys[int, int](nil, nil))
	assert.Equal(t, []int{}, MapUnionKeys(nil, map[int]int{}))
	assert.Equal(t, []int{1}, MapUnionKeys(map[int]int{1: 11}, nil))

	assert.True(t, ContentEqual([]int{1, 2, 3, 4},
		MapUnionKeys(map[int]int{1: 11, 2: 22}, map[int]int{3: 33, 4: 44})))
	assert.True(t, ContentEqual([]string{"1", "2", "3", "4"},
		MapUnionKeys(map[string]int{"1": 11, "2": 22}, map[string]int{"3": 33, "4": 44})))
}

func Test_MapIntersectionKeys(t *testing.T) {
	assert.Equal(t, []int{}, MapIntersectionKeys[int, int](nil, nil))
	assert.Equal(t, []int{}, MapIntersectionKeys(nil, map[int]int{}))
	assert.Equal(t, []int{}, MapIntersectionKeys(map[int]int{1: 11}, nil))

	assert.True(t, ContentEqual([]int{},
		MapIntersectionKeys(map[int]int{1: 11, 2: 22}, map[int]int{3: 33, 4: 44})))
	assert.True(t, ContentEqual([]string{"2"},
		MapIntersectionKeys(map[string]int{"1": 11, "2": 22}, map[string]int{"3": 33, "2": 22})))
}
