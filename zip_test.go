package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Zip(t *testing.T) {
	assert.Equal(t, []*Tuple2[int, bool]{}, Zip([]int{}, []bool{}))
	assert.Equal(t, []*Tuple2[int, bool]{}, Zip([]int{1, 2, 3}, []bool{}))
	assert.Equal(t, []*Tuple2[int, string]{{1, "1"}, {2, "2"}}, Zip([]int{1, 2, 3}, []string{"1", "2"}))
}

func Test_Zip3(t *testing.T) {
	assert.Equal(t, []*Tuple3[int, bool, int32]{}, Zip3([]int{}, []bool{}, []int32{}))
	assert.Equal(t, []*Tuple3[int, bool, int32]{}, Zip3([]int{1, 2, 3}, []bool{}, []int32{4, 5}))
	assert.Equal(t, []*Tuple3[int, string, bool]{{1, "1", true}, {2, "2", false}},
		Zip3([]int{1, 2, 3}, []string{"1", "2"}, []bool{true, false, false}))
}
