package gofn

import (
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
	assert.False(t, MapEqual(map[int]int{1: 1, 2: 2}, map[int]int{2: 2, 1: 11}))

	type st struct {
		Int int
		Str string
	}
	assert.True(t, MapEqual(map[int]st{1: {1, "1"}, 2: {2, "2"}},
		map[int]st{2: {2, "2"}, 1: {1, "1"}}))
	assert.False(t, MapEqual(map[int]st{1: {1, "1"}, 2: {2, "2"}},
		map[int]st{2: {2, "2"}, 1: {1, "1"}, 3: {3, "3"}}))
}

// nolint: gocritic
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
	assert.False(t, MapEqualPred(map[int]int{1: 1, 2: 2}, map[int]int{2: 2, 1: 11},
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
	assert.Equal(t, []int{}, MapKeys[int, bool, map[int]bool](nil))
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
	assert.Equal(t, []*Tuple2[int, bool]{}, MapEntries[int, bool, map[int]bool](nil))
	assert.Equal(t, []*Tuple2[int, int]{}, MapEntries(map[int]int{}))

	assert.Equal(t, []*Tuple2[int, int]{{1, 11}}, MapEntries(map[int]int{1: 11}))
	assert.True(t, ContentEqualPtr([]*Tuple2[int, int]{{1, 11}, {2, 22}}, MapEntries(map[int]int{1: 11, 2: 22})))
}

func Test_MapUpdate(t *testing.T) {
	// Update with a nil map
	assert.Equal(t, map[int]bool{}, MapUpdate(nil, map[int]bool{}))
	assert.Equal(t, map[int]bool{1: false, 2: true}, MapUpdate(map[int]bool{1: false, 2: true}, nil))
	// Update with an empty map
	assert.Equal(t, map[int]bool{1: false, 2: true}, MapUpdate(map[int]bool{1: false, 2: true}, map[int]bool{}))

	// Merge 2 maps
	assert.Equal(t, map[int]string{1: "one", 2: "two", 3: "three"},
		MapUpdate(map[int]string{2: "two"}, map[int]string{1: "one", 3: "three"}))
	// Merge 2 maps with override
	assert.Equal(t, map[int]string{1: "one", 2: "TWO", 3: "three"},
		MapUpdate(map[int]string{2: "two"}, map[int]string{1: "one", 2: "TWO", 3: "three"}))

	// Derived type
	type Map map[int]string
	assert.Equal(t, Map{1: "one", 2: "two", 3: "three"},
		MapUpdate(Map{2: "two"}, map[int]string{1: "one", 3: "three"}))
}

func Test_MapUpdateExistingOnly(t *testing.T) {
	// Update with a nil map
	assert.Equal(t, map[int]bool{}, MapUpdateExistingOnly(nil, map[int]bool{}))
	assert.Equal(t, map[int]bool{}, MapUpdateExistingOnly(map[int]bool{}, nil))
	assert.Equal(t, map[int]int{}, MapUpdateExistingOnly(map[int]int{}, map[int]int{1: 1}))
	assert.Equal(t, map[int]int{1: 1}, MapUpdateExistingOnly(map[int]int{1: 1}, map[int]int{}))
	assert.Equal(t, map[int]int{1: 1, 2: 2}, MapUpdateExistingOnly(map[int]int{1: 1, 2: 2}, map[int]int{3: 3, 4: 4}))
	assert.Equal(t, map[int]int{1: 1, 2: 22}, MapUpdateExistingOnly(map[int]int{1: 1, 2: 2}, map[int]int{3: 3, 2: 22}))
}

func Test_MapUpdateNewOnly(t *testing.T) {
	// Update with a nil map
	assert.Equal(t, map[int]bool{}, MapUpdateNewOnly(nil, map[int]bool{}))
	assert.Equal(t, map[int]bool{}, MapUpdateNewOnly(map[int]bool{}, nil))
	assert.Equal(t, map[int]int{}, MapUpdateNewOnly(map[int]int{}, map[int]int{}))
	assert.Equal(t, map[int]int{1: 1}, MapUpdateNewOnly(map[int]int{1: 1}, map[int]int{1: 11}))
	assert.Equal(t, map[int]int{1: 1, 2: 2}, MapUpdateNewOnly(map[int]int{1: 1, 2: 2}, map[int]int{1: 11, 2: 22}))
	assert.Equal(t, map[int]int{1: 1, 2: 2, 3: 3}, MapUpdateNewOnly(map[int]int{1: 1, 2: 2}, map[int]int{3: 3, 2: 22}))
}

func Test_MapGet(t *testing.T) {
	assert.Equal(t, true, MapGet[int, bool, map[int]bool](nil, 1, true))
	assert.Equal(t, 11, MapGet(map[int]int{}, 1, 11))

	assert.Equal(t, 11, MapGet(map[int]int{1: 11}, 1, 100))
	assert.Equal(t, 100, MapGet(map[int]int{1: 11, 2: 22}, 3, 100))
}

func Test_MapPop(t *testing.T) {
	assert.Equal(t, true, MapPop[int, bool, map[int]bool](nil, 1, true))
	assert.Equal(t, 11, MapPop(map[int]int{}, 1, 11))

	m1 := map[int]int{1: 11}
	assert.Equal(t, 11, MapPop(m1, 1, 100))
	assert.Equal(t, map[int]int{}, m1)
	m2 := map[int]int{1: 11, 2: 22}
	assert.Equal(t, 100, MapPop(m2, 3, 100))
	assert.Equal(t, map[int]int{1: 11, 2: 22}, m2)
}

func Test_MapSetDefault(t *testing.T) {
	assert.Equal(t, false, MapSetDefault[int, bool, map[int]bool](nil, 1, true))
	assert.Equal(t, 11, MapSetDefault(map[int]int{}, 1, 11))

	m1 := map[int]int{1: 11}
	assert.Equal(t, 11, MapSetDefault(m1, 1, 100))
	assert.Equal(t, map[int]int{1: 11}, m1)
	m2 := map[int]int{1: 11}
	assert.Equal(t, 22, MapSetDefault(m2, 2, 22))
	assert.Equal(t, map[int]int{1: 11, 2: 22}, m2)
}

func Test_MapUnionKeys(t *testing.T) {
	assert.Equal(t, []int{}, MapUnionKeys[int, int, map[int]int](nil, nil))
	assert.Equal(t, []int{}, MapUnionKeys(nil, map[int]int{}))
	assert.Equal(t, []int{1}, MapUnionKeys(map[int]int{1: 11}, nil))

	assert.True(t, ContentEqual([]int{1, 2, 3, 4},
		MapUnionKeys(map[int]int{1: 11, 2: 22}, map[int]int{3: 33, 4: 44})))
	assert.True(t, ContentEqual([]string{"1", "2", "3", "4"},
		MapUnionKeys(map[string]int{"1": 11, "2": 22}, map[string]int{"3": 33, "4": 44})))
}

func Test_MapIntersectionKeys(t *testing.T) {
	assert.Equal(t, []int{}, MapIntersectionKeys[int, int, map[int]int](nil, nil))
	assert.Equal(t, []int{}, MapIntersectionKeys(nil, map[int]int{}))
	assert.Equal(t, []int{}, MapIntersectionKeys(map[int]int{1: 11}, nil))

	assert.True(t, ContentEqual([]int{},
		MapIntersectionKeys(map[int]int{1: 11, 2: 22}, map[int]int{3: 33, 4: 44})))
	assert.True(t, ContentEqual([]string{"2"},
		MapIntersectionKeys(map[string]int{"1": 11, "2": 22}, map[string]int{"3": 33, "2": 22})))
}

func Test_MapDifferenceKeys(t *testing.T) {
	l, r := MapDifferenceKeys[int, int, map[int]int](nil, nil)
	assert.Equal(t, []int{}, l)
	assert.Equal(t, []int{}, r)
	l, r = MapDifferenceKeys(nil, map[int]int{})
	assert.Equal(t, []int{}, l)
	assert.Equal(t, []int{}, r)
	l, r = MapDifferenceKeys(map[int]int{1: 11}, nil)
	assert.Equal(t, []int{1}, l)
	assert.Equal(t, []int{}, r)

	// Derived types
	type MapII map[int]int
	type MapSI map[string]int

	l, r = MapDifferenceKeys(map[int]int{1: 11, 2: 22}, MapII{3: 33, 4: 44})
	assert.True(t, ContentEqual([]int{1, 2}, l))
	assert.True(t, ContentEqual([]int{3, 4}, r))

	l2, r2 := MapDifferenceKeys(MapSI{"1": 11, "2": 22}, map[string]int{"3": 33, "2": 22})
	assert.True(t, ContentEqual([]string{"1"}, l2))
	assert.True(t, ContentEqual([]string{"3"}, r2))
}

func Test_MapCopy(t *testing.T) {
	// Nil/empty maps
	assert.Equal(t, map[int]int{}, MapCopy[int, int](map[int]int(nil)))
	assert.Equal(t, map[int]int{}, MapCopy[int, int](map[int]int{}))

	assert.True(t, MapEqual(map[int]int{1: 11, 2: 22}, MapCopy(map[int]int{1: 11, 2: 22})))
}

func Test_MapPick(t *testing.T) {
	// Nil/empty maps
	assert.Equal(t, map[int]int{}, MapPick[int, int](map[int]int(nil)))
	assert.Equal(t, map[int]int{}, MapPick[int, int](map[int]int{}, 1, 2, 3))

	assert.True(t, MapEqual(map[int]int{}, MapPick(map[int]int{1: 11, 2: 22})))
	assert.True(t, MapEqual(map[int]int{2: 22}, MapPick(map[int]int{1: 11, 2: 22}, 2, 3, 2)))
}

func Test_MapOmitCopy(t *testing.T) {
	// Nil/empty maps
	assert.Equal(t, map[int]int{}, MapOmitCopy[int, int, map[int]int](nil))
	assert.Equal(t, map[int]int{}, MapOmitCopy[int, int](map[int]int{}, 1, 2, 3))

	assert.True(t, MapEqual(map[int]int{1: 11, 2: 22}, MapOmitCopy(map[int]int{1: 11, 2: 22})))
	assert.True(t, MapEqual(map[int]int{1: 11}, MapOmitCopy(map[int]int{1: 11, 2: 22}, 2, 3, 2)))
}

func Test_MapCopyExcludeKeys(t *testing.T) {
	m := MapCopyExcludeKeys(map[int]int{1: 11, 2: 22}, 2, 3)
	assert.True(t, MapEqual(map[int]int{1: 11}, m))
}
