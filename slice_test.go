package gofn

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Equal(t *testing.T) {
	assert.True(t, Equal([]int{}, []int{}))
	assert.True(t, Equal([]int{}, nil))
	assert.True(t, Equal(nil, []int{}))
	assert.True(t, Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.True(t, Equal([]string{"3", "1", "2"}, []string{"3", "1", "2"}))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, Equal([]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{1, "1"}, {2, "2"}, {3, "3"}}))

	assert.False(t, Equal([]int{}, []int{1}))
	assert.False(t, Equal([]int{1}, nil))
	assert.False(t, Equal([]int{1, 2, 3}, []int{1, 2, 3, 4}))
	assert.False(t, Equal([]int{1, 2, 3}, []int{3, 2, 1}))
	assert.False(t, Equal([]string{"3", "1", "2"}, []string{"1", "2", "3"}))
	assert.False(t, Equal([]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{1, "1"}, {2, "2"}}))
	assert.False(t, Equal([]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{3, "3"}, {1, "1"}, {2, "2"}}))
}

func Test_EqualPred(t *testing.T) {
	assert.True(t, EqualPred([]interface{}{}, []interface{}{},
		func(a, b interface{}) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualPred([]interface{}{}, nil,
		func(a, b interface{}) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualPred(nil, []interface{}{},
		func(a, b interface{}) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualPred([]interface{}{1, 2, 3}, []interface{}{1, 2, 3},
		func(a, b interface{}) bool { return a.(int) == b.(int) }))
	assert.True(t, EqualPred([]interface{}{"3", "1", "2"}, []interface{}{"3", "1", "2"},
		func(a, b interface{}) bool { return a.(string) == b.(string) }))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, EqualPred(
		[]interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(a, b interface{}) bool { return a.(St) == b.(St) }))

	assert.False(t, EqualPred([]interface{}{}, []interface{}{1},
		func(a, b interface{}) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualPred([]interface{}{1}, nil,
		func(a, b interface{}) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualPred([]interface{}{1, 2, 3}, []interface{}{1, 2, 3, 4},
		func(a, b interface{}) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualPred([]interface{}{1, 2, 3}, []interface{}{3, 2, 1},
		func(a, b interface{}) bool { return a.(int) == b.(int) }))
	assert.False(t, EqualPred([]interface{}{"3", "1", "2"}, []interface{}{"1", "2", "3"},
		func(a, b interface{}) bool { return a.(string) == b.(string) }))
	assert.False(t, EqualPred(
		[]interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []interface{}{St{1, "1"}, St{2, "2"}},
		func(a, b interface{}) bool { return a.(St) == b.(St) }))
	assert.False(t, EqualPred(
		[]interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []interface{}{St{3, "3"}, St{2, "2"}, St{1, "1"}},
		func(a, b interface{}) bool { return a.(St) == b.(St) }))
}

func Test_ContentEqual(t *testing.T) {
	assert.True(t, ContentEqual([]int{}, []int{}))
	assert.True(t, ContentEqual([]int{}, nil))
	assert.True(t, ContentEqual(nil, []int{}))
	assert.True(t, ContentEqual([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.True(t, ContentEqual([]int{3, 1, 2}, []int{1, 2, 3}))
	assert.True(t, ContentEqual([]string{"3", "1", "2"}, []string{"1", "2", "3"}))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, ContentEqual([]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{3, "3"}, {1, "1"}, {2, "2"}}))

	assert.False(t, ContentEqual([]int{}, []int{1}))
	assert.False(t, ContentEqual([]int{1}, nil))
	assert.False(t, ContentEqual([]int{1, 2, 3}, []int{1, 2, 3, 4}))
	assert.False(t, ContentEqual([]St{{1, "1"}, {2, "2"}, {3, "3"}}, []St{{3, "3"}, {1, "1"}, {3, "3"}}))
}

func Test_ContentEqualPred(t *testing.T) {
	assert.True(t, ContentEqualPred([]interface{}{}, []interface{}{},
		func(t interface{}) int { return t.(int) }))
	assert.True(t, ContentEqualPred([]interface{}{}, nil,
		func(t interface{}) int { return t.(int) }))
	assert.True(t, ContentEqualPred(nil, []interface{}{},
		func(t interface{}) int { return t.(int) }))
	assert.True(t, ContentEqualPred([]interface{}{1, 2, 3}, []interface{}{1, 2, 3},
		func(t interface{}) int { return t.(int) }))
	assert.True(t, ContentEqualPred([]interface{}{3, 1, 2}, []interface{}{1, 2, 3},
		func(t interface{}) int { return t.(int) }))
	assert.True(t, ContentEqualPred([]interface{}{"3", "1", "2"}, []interface{}{"1", "2", "3"},
		func(t interface{}) string { return t.(string) }))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, ContentEqualPred(
		[]interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []interface{}{St{3, "3"}, St{1, "1"}, St{2, "2"}},
		func(t interface{}) St { return t.(St) }))

	assert.False(t, ContentEqualPred([]interface{}{}, []interface{}{1},
		func(t interface{}) int { return t.(int) }))
	assert.False(t, ContentEqualPred([]interface{}{1}, nil,
		func(t interface{}) int { return t.(int) }))
	assert.False(t, ContentEqualPred([]interface{}{1, 2, 3}, []interface{}{1, 2, 3, 4},
		func(t interface{}) int { return t.(int) }))
	assert.False(t, ContentEqualPred(
		[]interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}}, []interface{}{St{3, "3"}, St{1, "1"}, St{3, "3"}},
		func(t interface{}) St { return t.(St) }))
}

func Test_ContentEqualPtr(t *testing.T) {
	assert.True(t, ContentEqualPtr([]*int{}, []*int{}))
	assert.True(t, ContentEqualPtr([]*int{}, nil))
	assert.True(t, ContentEqualPtr(nil, []*int{}))
	i1, i2, i3 := New(1), New(2), New(3)
	assert.True(t, ContentEqual([]*int{i3, i1, i2}, []*int{i1, i2, i3}))

	type St struct {
		Int int
		Str string
	}
	assert.True(t, ContentEqualPtr([]*St{{1, "1"}, {2, "2"}, {3, "3"}}, []*St{{3, "3"}, {1, "1"}, {2, "2"}}))

	assert.False(t, ContentEqualPtr([]*int{}, []*int{i1}))
	assert.False(t, ContentEqualPtr([]*int{i1}, nil))
	assert.False(t, ContentEqualPtr([]*int{i1, i2, i3}, []*int{i1, i2}))
	assert.False(t, ContentEqualPtr([]*St{{1, "1"}, {2, "2"}, {3, "3"}}, []*St{{3, "3"}, {1, "1"}, {3, "3"}}))
}

func Test_Concat(t *testing.T) {
	assert.Equal(t, []int{}, Concat[int](nil, nil, nil))
	assert.Equal(t, []bool{}, Concat([]bool{}, []bool{}))
	assert.Equal(t, []float64{1.1}, Concat([]float64{}, []float64{}, []float64{1.1}))
	assert.Equal(t, []string{"", "1", "2", "3"}, Concat([]string{""}, []string{"1", "2"}, []string{}, []string{"3"}))
}

func Test_Contain(t *testing.T) {
	assert.False(t, Contain([]int{}, 1))
	assert.False(t, Contain([]string{"one"}, "One"))
	assert.False(t, Contain([]string{"one", "two"}, ""))
	assert.False(t, Contain([]int64{1, 2, 3}, 4))
	assert.False(t, Contain([]float32{1.1, 2.2, 3.3}, 3.35))

	type St struct {
		Int int
		Str string
	}
	assert.False(t, Contain([]St{{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "4"}))

	assert.True(t, Contain([]int64{1}, 1))
	assert.True(t, Contain([]uint{1, 2, 3, 1, 2, 3}, 2))
	assert.True(t, Contain([]string{"one", "two"}, "two"))
	assert.True(t, Contain([]string{"one", "two", ""}, ""))
	assert.True(t, Contain([]float64{1.1, 2.2, 3.3}, 2.2))
	assert.True(t, Contain([]St{{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "3"}))
}

func Test_ContainPred(t *testing.T) {
	assert.False(t, ContainPred([]interface{}{},
		func(i interface{}) bool { return i.(int) == 1 }))
	assert.False(t, ContainPred([]interface{}{"one"},
		func(i interface{}) bool { return i == "One" }))
	assert.False(t, ContainPred([]interface{}{"one", "two"},
		func(i interface{}) bool { return i == "" }))
	assert.False(t, ContainPred([]interface{}{1, 2, 3},
		func(i interface{}) bool { return i == 4 }))
	assert.False(t, ContainPred([]interface{}{1.1, 2.2, 3.3},
		func(i interface{}) bool { return i == 3.35 }))

	type St struct {
		Int int
		Str string
	}
	assert.False(t, ContainPred([]interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i interface{}) bool { return i == St{3, "4"} }))

	assert.True(t, ContainPred([]interface{}{1},
		func(i interface{}) bool { return i == 1 }))
	assert.True(t, ContainPred([]interface{}{1, 2, 3, 1, 2, 3},
		func(i interface{}) bool { return i == 2 }))
	assert.True(t, ContainPred([]interface{}{"one", "two"},
		func(i interface{}) bool { return i == "two" }))
	assert.True(t, ContainPred([]interface{}{"one", "two", ""},
		func(i interface{}) bool { return i == "" }))
	assert.True(t, ContainPred([]interface{}{1.1, 2.2, 3.3},
		func(i interface{}) bool { return i == 2.2 }))
	assert.True(t, ContainPred([]interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i interface{}) bool { return i == St{3, "3"} }))
}

func Test_ContainAll(t *testing.T) {
	assert.False(t, ContainAll([]int{}, 0))
	assert.False(t, ContainAll([]string{"one"}, "one", "two"))
	assert.False(t, ContainAll([]string{"one", "two"}, "two", "one", ""))
	assert.False(t, ContainAll([]int64{1, 2, 3}, 3, 3, 2, 1, 0))
	assert.False(t, ContainAll([]float32{1.1, 2.2, 3.3}, 3.35, 2.2))

	assert.True(t, ContainAll([]int64{1}, 1))
	assert.True(t, ContainAll([]uint{1, 2, 3, 1, 2, 3}, 2, 1, 3, 3, 2))
	assert.True(t, ContainAll([]string{"one", "two"}, "two"))
	assert.True(t, ContainAll([]string{"one", "two", ""}, "", "two"))
	assert.True(t, ContainAll([]float64{1.1, 2.2, 3.3}, 2.2, 3.3, 1.1))

	// More than 10 items in slice
	assert.False(t, ContainAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 1, 2, 3, 0))
	assert.False(t, ContainAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 12))
	assert.True(t, ContainAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 1, 2, 3, 3, 2, 1))
	assert.True(t, ContainAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11))
	assert.True(t, ContainAll([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
		"1", "2", "3", "3", "2", "1"))
}

func Test_ContainAny(t *testing.T) {
	assert.False(t, ContainAny([]int{}, 0))
	assert.False(t, ContainAny([]string{"one"}, "One", "three"))
	assert.False(t, ContainAny([]string{"one", "two"}, ""))
	assert.False(t, ContainAny([]int64{1, 2, 3}, 0, 4, 5))
	assert.False(t, ContainAny([]float32{1.1, 2.2, 3.3}, 3.35, 2.22))

	assert.True(t, ContainAny([]int64{1}, 1, 1))
	assert.True(t, ContainAny([]uint{1, 2, 3, 1, 2, 3}, 10, 11, 1))
	assert.True(t, ContainAny([]string{"one", "two"}, "two", "three"))
	assert.True(t, ContainAny([]string{"one", "two", ""}, "", "two"))
	assert.True(t, ContainAny([]float64{1.1, 2.2, 3.3}, 2.2, 3.33, 1.11))

	// More than 10 items in slice
	assert.False(t, ContainAny([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 0, -1, 100))
	assert.False(t, ContainAny([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 13))
	assert.True(t, ContainAny([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 1, 2, 3, 3, 2, 1))
	assert.True(t, ContainAny([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 6, 100, 101, 102))
	assert.True(t, ContainAny([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}, "100", "0", "1", "2"))
}

func Test_IndexOf(t *testing.T) {
	assert.Equal(t, -1, IndexOf([]int{}, 1))
	assert.Equal(t, -1, IndexOf([]string{"one"}, "One"))
	assert.Equal(t, -1, IndexOf([]string{"one", "two"}, ""))
	assert.Equal(t, -1, IndexOf([]int64{1, 2, 3}, 4))
	assert.Equal(t, -1, IndexOf([]float32{1.1, 2.2, 3.3}, 3.35))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, -1, IndexOf([]St{{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "4"}))

	assert.Equal(t, 0, IndexOf([]int64{1}, 1))
	assert.Equal(t, 1, IndexOf([]uint{1, 2, 3, 1, 2, 3}, 2))
	assert.Equal(t, 1, IndexOf([]string{"one", "two"}, "two"))
	assert.Equal(t, 2, IndexOf([]string{"one", "two", ""}, ""))
	assert.Equal(t, 2, IndexOf([]float64{1.1, 2.2, 3.3}, 3.3))
	assert.Equal(t, 2, IndexOf([]St{St{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "3"}))
}

func Test_IndexOfPred(t *testing.T) {
	assert.Equal(t, -1, IndexOfPred([]interface{}{},
		func(i interface{}) bool { return i == 1 }))
	assert.Equal(t, -1, IndexOfPred([]interface{}{"one"},
		func(i interface{}) bool { return i == "One" }))
	assert.Equal(t, -1, IndexOfPred([]interface{}{"one", "two"},
		func(i interface{}) bool { return i == "" }))
	assert.Equal(t, -1, IndexOfPred([]interface{}{1, 2, 3},
		func(i interface{}) bool { return i == 4 }))
	assert.Equal(t, -1, IndexOfPred([]interface{}{1.1, 2.2, 3.3},
		func(i interface{}) bool { return i == 3.35 }))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, -1, IndexOfPred([]interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i interface{}) bool { return i == St{3, "4"} }))

	assert.Equal(t, 0, IndexOfPred([]interface{}{1},
		func(i interface{}) bool { return i == 1 }))
	assert.Equal(t, 1, IndexOfPred([]interface{}{1, 2, 3, 1, 2, 3},
		func(i interface{}) bool { return i == 2 }))
	assert.Equal(t, 1, IndexOfPred([]interface{}{"one", "two"},
		func(i interface{}) bool { return i == "two" }))
	assert.Equal(t, 2, IndexOfPred([]interface{}{"one", "two", ""},
		func(i interface{}) bool { return i == "" }))
	assert.Equal(t, 2, IndexOfPred([]interface{}{1.1, 2.2, 3.3},
		func(i interface{}) bool { return i == 3.3 }))
	assert.Equal(t, 2, IndexOfPred([]interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i interface{}) bool { return i == St{3, "3"} }))
}

func Test_LastIndexOf(t *testing.T) {
	assert.Equal(t, -1, LastIndexOf([]int{}, 1))
	assert.Equal(t, -1, LastIndexOf([]string{"one"}, "One"))
	assert.Equal(t, -1, LastIndexOf([]string{"one", "two"}, ""))
	assert.Equal(t, -1, LastIndexOf([]int64{1, 2, 3}, 4))
	assert.Equal(t, -1, LastIndexOf([]float32{1.1, 2.2, 3.3}, 3.35))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, -1, LastIndexOf([]St{{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "4"}))

	assert.Equal(t, 0, LastIndexOf([]int64{1}, 1))
	assert.Equal(t, 4, LastIndexOf([]uint{1, 2, 3, 1, 2, 3}, 2))
	assert.Equal(t, 1, LastIndexOf([]string{"one", "two"}, "two"))
	assert.Equal(t, 3, LastIndexOf([]string{"one", "", "two", ""}, ""))
	assert.Equal(t, 0, LastIndexOf([]float64{1.1, 2.2, 3.3}, 1.1))
	assert.Equal(t, 2, LastIndexOf([]St{{1, "1"}, {2, "2"}, {3, "3"}}, St{3, "3"}))
}

func Test_LastIndexOfPred(t *testing.T) {
	assert.Equal(t, -1, LastIndexOfPred([]interface{}{},
		func(i interface{}) bool { return i == 1 }))
	assert.Equal(t, -1, LastIndexOfPred([]interface{}{"one"},
		func(i interface{}) bool { return i == "One" }))
	assert.Equal(t, -1, LastIndexOfPred([]interface{}{"one", "two"},
		func(i interface{}) bool { return i == "" }))
	assert.Equal(t, -1, LastIndexOfPred([]interface{}{1, 2, 3},
		func(i interface{}) bool { return i == 4 }))
	assert.Equal(t, -1, LastIndexOfPred([]interface{}{1.1, 2.2, 3.3},
		func(i interface{}) bool { return i == 3.35 }))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, -1, LastIndexOfPred([]interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i interface{}) bool { return i == St{3, "4"} }))

	assert.Equal(t, 0, LastIndexOfPred([]interface{}{1},
		func(i interface{}) bool { return i == 1 }))
	assert.Equal(t, 4, LastIndexOfPred([]interface{}{1, 2, 3, 1, 2, 3},
		func(i interface{}) bool { return i == 2 }))
	assert.Equal(t, 1, LastIndexOfPred([]interface{}{"one", "two"},
		func(i interface{}) bool { return i == "two" }))
	assert.Equal(t, 3, LastIndexOfPred([]interface{}{"one", "", "two", ""},
		func(i interface{}) bool { return i == "" }))
	assert.Equal(t, 1, LastIndexOfPred([]interface{}{1.1, 1.1, 2.2, 3.3},
		func(i interface{}) bool { return i == 1.1 }))
	assert.Equal(t, 2, LastIndexOfPred([]interface{}{St{1, "1"}, St{2, "2"}, St{3, "3"}},
		func(i interface{}) bool { return i == St{3, "3"} }))
}

func Test_RemoveAt(t *testing.T) {
	s1 := []int{1}
	RemoveAt(&s1, 0)
	assert.Equal(t, []int{}, s1)

	s1 = []int{1, 2, 3}
	RemoveAt(&s1, 1)
	assert.Equal(t, []int{1, 3}, s1)

	s1 = []int{1, 2, 3}
	RemoveAt(&s1, 2)
	assert.Equal(t, []int{1, 2}, s1)

	s2 := []string{"", "one", "two"}
	RemoveAt(&s2, 0)
	assert.Equal(t, []string{"one", "two"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}}
	RemoveAt(&s3, 1)
	assert.Equal(t, []St{{1, "1"}, {3, "3"}}, s3)
}

func Test_FastRemoveAt(t *testing.T) {
	s1 := []int{1}
	FastRemoveAt(&s1, 0)
	assert.Equal(t, []int{}, s1)

	s1 = []int{1, 2, 3}
	FastRemoveAt(&s1, 0)
	assert.Equal(t, []int{3, 2}, s1)

	s1 = []int{1, 2, 3}
	FastRemoveAt(&s1, 2)
	assert.Equal(t, []int{1, 2}, s1)

	s2 := []string{"", "one", "two"}
	FastRemoveAt(&s2, 0)
	assert.Equal(t, []string{"two", "one"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}}
	FastRemoveAt(&s3, 1)
	assert.Equal(t, []St{{1, "1"}, {3, "3"}}, s3)
}

func Test_Remove(t *testing.T) {
	s1 := []int{1}
	Remove(&s1, 2)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1}
	Remove(&s1, 1)
	assert.Equal(t, []int{}, s1)

	s1 = []int{1, 2, 3}
	Remove(&s1, 1)
	assert.Equal(t, []int{2, 3}, s1)

	s1 = []int{1, 2, 3}
	Remove(&s1, 2)
	assert.Equal(t, []int{1, 3}, s1)

	s2 := []string{"", "one", "two"}
	Remove(&s2, "two")
	assert.Equal(t, []string{"", "one"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}}
	Remove(&s3, St{2, "2"})
	assert.Equal(t, []St{{1, "1"}, {3, "3"}}, s3)
}

func Test_FastRemove(t *testing.T) {
	s1 := []int{1}
	FastRemove(&s1, 2)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1}
	FastRemove(&s1, 1)
	assert.Equal(t, []int{}, s1)

	s1 = []int{1, 2, 3}
	FastRemove(&s1, 1)
	assert.Equal(t, []int{3, 2}, s1)

	s1 = []int{1, 2, 3}
	FastRemove(&s1, 3)
	assert.Equal(t, []int{1, 2}, s1)

	s2 := []string{"", "one", "two"}
	FastRemove(&s2, "")
	assert.Equal(t, []string{"two", "one"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}}
	FastRemove(&s3, St{2, "2"})
	assert.Equal(t, []St{{1, "1"}, {3, "3"}}, s3)
}

func Test_RemoveLastOf(t *testing.T) {
	s1 := []int{1}
	RemoveLastOf(&s1, 2)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1, 1}
	RemoveLastOf(&s1, 1)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1, 2, 1}
	RemoveLastOf(&s1, 1)
	assert.Equal(t, []int{1, 2}, s1)

	s1 = []int{1, 2, 3, 2, 2}
	RemoveLastOf(&s1, 2)
	assert.Equal(t, []int{1, 2, 3, 2}, s1)

	s2 := []string{"", "one", "", "two"}
	RemoveLastOf(&s2, "")
	assert.Equal(t, []string{"", "one", "two"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}, {2, "2"}, {3, "3"}}
	RemoveLastOf(&s3, St{2, "2"})
	assert.Equal(t, []St{{1, "1"}, {2, "2"}, {3, "3"}, {3, "3"}}, s3)
}

func Test_FastRemoveLastOf(t *testing.T) {
	s1 := []int{1}
	FastRemoveLastOf(&s1, 2)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1, 1}
	FastRemoveLastOf(&s1, 1)
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1, 2, 1, 2, 3}
	FastRemoveLastOf(&s1, 1)
	assert.Equal(t, []int{1, 2, 3, 2}, s1)

	s1 = []int{1, 2, 3, 2, 2}
	FastRemoveLastOf(&s1, 2)
	assert.Equal(t, []int{1, 2, 3, 2}, s1)

	s2 := []string{"", "one", "", "two", "three"}
	FastRemoveLastOf(&s2, "")
	assert.Equal(t, []string{"", "one", "three", "two"}, s2)

	type St struct {
		Int int
		Str string
	}
	s3 := []St{{1, "1"}, {2, "2"}, {3, "3"}, {2, "2"}, {3, "3"}}
	FastRemoveLastOf(&s3, St{2, "2"})
	assert.Equal(t, []St{{1, "1"}, {2, "2"}, {3, "3"}, {3, "3"}}, s3)
}

func Test_RemoveAll(t *testing.T) {
	s1 := []int{1}
	assert.Equal(t, 0, RemoveAll(&s1, 2))
	assert.Equal(t, []int{1}, s1)

	s1 = []int{1, 1}
	assert.Equal(t, 2, RemoveAll(&s1, 1))
	assert.Equal(t, []int{}, s1)

	s1 = []int{1, 2, 1, 2, 3, 2}
	assert.Equal(t, 3, RemoveAll(&s1, 2))
	assert.Equal(t, []int{1, 1, 3}, s1)

	s1 = []int{1, 2, 3, 1, 2, 3, 1}
	assert.Equal(t, 3, RemoveAll(&s1, 1))
	assert.Equal(t, []int{2, 3, 2, 3}, s1)

	s2 := []string{"", "one", "", "two", "three", ""}
	assert.Equal(t, 3, RemoveAll(&s2, ""))
	assert.Equal(t, []string{"one", "two", "three"}, s2)
}

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

func Test_Replace(t *testing.T) {
	// No replacement done
	s := []int{}
	assert.True(t, Replace(s, 1, 11) == false && reflect.DeepEqual(s, []int{}))
	s = []int{0, 1, 2, 3, 4, 5}
	assert.True(t, Replace(s, 100, 1000) == false && reflect.DeepEqual(s, []int{0, 1, 2, 3, 4, 5}))
	s2 := []string{"one", "two"}
	assert.True(t, Replace(s2, "One", "ONE") == false && reflect.DeepEqual(s2, []string{"one", "two"}))

	// Has replacement
	s = []int{0, 1, 2, 3, 0, 1, 2, 3}
	assert.True(t, Replace(s, 1, 11) == true && reflect.DeepEqual(s, []int{0, 11, 2, 3, 0, 1, 2, 3}))
	s2 = []string{"one", "two", "one"}
	assert.True(t, Replace(s2, "one", "1") == true && reflect.DeepEqual(s2, []string{"1", "two", "one"}))
}

func Test_ReplaceAll(t *testing.T) {
	// No replacement done
	s := []int{}
	assert.True(t, ReplaceAll(s, 1, 11) == 0 && reflect.DeepEqual(s, []int{}))
	s = []int{0, 1, 2, 3, 4, 5}
	assert.True(t, ReplaceAll(s, 100, 1000) == 0 && reflect.DeepEqual(s, []int{0, 1, 2, 3, 4, 5}))
	s2 := []string{"one", "two"}
	assert.True(t, ReplaceAll(s2, "One", "ONE") == 0 && reflect.DeepEqual(s2, []string{"one", "two"}))

	// Has replacement
	s = []int{0, 1, 2, 3, 0, 1, 2, 3}
	assert.True(t, ReplaceAll(s, 1, 11) == 2 && reflect.DeepEqual(s, []int{0, 11, 2, 3, 0, 11, 2, 3}))
	s2 = []string{"one", "two", "one"}
	assert.True(t, ReplaceAll(s2, "one", "1") == 2 && reflect.DeepEqual(s2, []string{"1", "two", "1"}))
}

func Test_Fill(t *testing.T) {
	s := make([]int, 5)
	Fill(s, 1)
	assert.Equal(t, []int{1, 1, 1, 1, 1}, s)
	Fill(s, 2)
	assert.Equal(t, []int{2, 2, 2, 2, 2}, s)
	sub := s[2:]
	Fill(sub, 5)
	assert.Equal(t, []int{5, 5, 5}, sub)
	assert.Equal(t, []int{2, 2, 5, 5, 5}, s)

	s2 := []int{}
	Fill(s2, 1)
	assert.Equal(t, []int{}, s2)
}

func Test_CountValue(t *testing.T) {
	assert.Equal(t, 0, CountValue([]int{1, 2, 3}, 4))
	assert.Equal(t, 1, CountValue([]int{1, 2, 3}, 2))
	assert.Equal(t, 2, CountValue([]int{1, 2, 3, 2}, 2))
	assert.Equal(t, 2, CountValue([]float32{1.1, 2.2, 3.3, 1.100001, 1.1}, 1.1))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, 2, CountValue([]St{{1, "1"}, {2, "2"}, {1, "1"}}, St{1, "1"}))
	assert.Equal(t, 0, CountValue([]St{{1, "1"}, {2, "2"}, {1, "1"}}, St{1, "2"}))
}

func Test_CountValuePred(t *testing.T) {
	assert.Equal(t, 0, CountValuePred([]interface{}{1, 2, 3},
		func(t interface{}) bool { return t == 4 }))
	assert.Equal(t, 1, CountValuePred([]interface{}{1, 2, 3},
		func(t interface{}) bool { return t == 2 }))
	assert.Equal(t, 2, CountValuePred([]interface{}{1, 2, 3, 2},
		func(t interface{}) bool { return t == 2 }))
	assert.Equal(t, 2, CountValuePred([]interface{}{1.1, 2.2, 3.3, 1.100001, 1.1},
		func(t interface{}) bool { return t == 1.1 }))

	type St struct {
		Int int
		Str string
	}
	assert.Equal(t, 2, CountValuePred([]interface{}{St{1, "1"}, St{2, "2"}, St{1, "1"}},
		func(t interface{}) bool { return t == St{1, "1"} }))
	assert.Equal(t, 0, CountValuePred([]interface{}{St{1, "1"}, St{2, "2"}, St{1, "1"}},
		func(t interface{}) bool { return t == St{1, "2"} }))
}

func Test_GetFirst(t *testing.T) {
	assert.Equal(t, 1, GetFirst([]int{1, 2, 3}, 4))
	assert.Equal(t, 11, GetFirst([]int{}, 11))
}

func Test_GetLast(t *testing.T) {
	assert.Equal(t, 3, GetLast([]int{1, 2, 3}, 4))
	assert.Equal(t, 11, GetLast([]int{}, 11))
}

func Test_IndexOfSlice(t *testing.T) {
	assert.Equal(t, -1, IndexOfSlice([]int{}, nil))
	assert.Equal(t, -1, IndexOfSlice([]string{"one"}, []string{}))
	assert.Equal(t, -1, IndexOfSlice([]string{"one", "two"}, []string{"Two"}))
	assert.Equal(t, -1, IndexOfSlice([]int64{1, 2, 3}, []int64{1, 2, 3, 4}))
	assert.Equal(t, -1, IndexOfSlice([]uint{0, 1, 2, 3, 4, 5}, []uint{3, 4, 5, 6}))
	assert.Equal(t, -1, IndexOfSlice([]float32{1.1, 2.2, 3.3}, []float32{2.2, 3.31}))

	assert.Equal(t, 0, IndexOfSlice([]int{1}, []int{1}))
	assert.Equal(t, 2, IndexOfSlice([]int{0, 1, 2}, []int{2}))
	assert.Equal(t, 0, IndexOfSlice([]int{0, 1, 2, 0, 1, 2, 3}, []int{0, 1, 2}))
	assert.Equal(t, 1, IndexOfSlice([]string{"one", ""}, []string{""}))
	assert.Equal(t, 0, IndexOfSlice([]string{"one", "two", "three"}, []string{"one", "two"}))
	assert.Equal(t, 0, IndexOfSlice([]int64{1, 2, 3}, []int64{1, 2, 3}))
	assert.Equal(t, 1, IndexOfSlice([]uint{0, 1, 1, 1, 1}, []uint{1}))
}

func Test_LastIndexOfSlice(t *testing.T) {
	assert.Equal(t, -1, LastIndexOfSlice([]int{}, nil))
	assert.Equal(t, -1, LastIndexOfSlice([]string{"one"}, []string{}))
	assert.Equal(t, -1, LastIndexOfSlice([]string{"one", "two"}, []string{"Two"}))
	assert.Equal(t, -1, LastIndexOfSlice([]int64{1, 2, 3}, []int64{1, 2, 3, 4}))
	assert.Equal(t, -1, LastIndexOfSlice([]uint{0, 1, 2, 3, 4, 5}, []uint{3, 4, 5, 6}))
	assert.Equal(t, -1, LastIndexOfSlice([]float32{1.1, 2.2, 3.3}, []float32{2.2, 3.31}))

	assert.Equal(t, 0, LastIndexOfSlice([]int{1}, []int{1}))
	assert.Equal(t, 2, LastIndexOfSlice([]int{0, 1, 2}, []int{2}))
	assert.Equal(t, 3, LastIndexOfSlice([]int{0, 1, 2, 0, 1, 2, 3}, []int{0, 1, 2}))
	assert.Equal(t, 2, LastIndexOfSlice([]string{"", "one", ""}, []string{""}))
	assert.Equal(t, 0, LastIndexOfSlice([]string{"one", "two", "three"}, []string{"one", "two"}))
	assert.Equal(t, 0, LastIndexOfSlice([]int64{1, 2, 3}, []int64{1, 2, 3}))
	assert.Equal(t, 4, LastIndexOfSlice([]uint{0, 1, 1, 1, 1}, []uint{1}))
}
