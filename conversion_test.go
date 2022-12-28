package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToIntfSlice(t *testing.T) {
	assert.Equal(t, []interface{}{}, ToIntfSlice([]int{}))
	assert.Equal(t, []interface{}{"one"}, ToIntfSlice([]string{"one"}))
	assert.Equal(t, []interface{}{1, 2, 3}, ToIntfSlice([]int{1, 2, 3}))
	assert.Equal(t, []interface{}{float32(1.1), float32(2.2), float32(3.3)}, ToIntfSlice([]float32{1.1, 2.2, 3.3}))
	assert.Equal(t, []interface{}{"one", 2, 3.3}, ToIntfSlice([]interface{}{"one", 2, 3.3}))
}

func Test_ToStringSlice(t *testing.T) {
	type StrType string
	assert.Equal(t, []string{}, ToStringSlice[string]([]StrType{}))
	assert.Equal(t, []string{"one", "two"}, ToStringSlice[string]([]StrType{"one", "two"}))
}

func Test_ToNumberSlice(t *testing.T) {
	type NumType int
	assert.Equal(t, []int16{}, ToNumberSlice[int16]([]NumType{}))
	assert.Equal(t, []int16{1, -1, 0, -32768, 32767}, ToNumberSlice[int16]([]NumType{1, 65535, 65536, 32768, 32767}))
	assert.Equal(t, []int64{-1, 65535}, ToNumberSlice[int64]([]NumType{-1, 65535}))
}

func Test_ToSlice(t *testing.T) {
	assert.Equal(t, []int{}, ToSlice[int]())
	assert.Equal(t, []int{1}, ToSlice[int](1))
	assert.Equal(t, []string{"1", "2", "3"}, ToSlice("1", "2", "3"))
}
