package gofn

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
