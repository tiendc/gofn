package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToSet(t *testing.T) {
	assert.Equal(t, []int{}, ToSet([]int{}))
	assert.Equal(t, []string{"one"}, ToSet([]string{"one"}))
	assert.Equal(t, []string{"one", "two", "Two"}, ToSet([]string{"one", "two", "one", "Two"}))
	assert.Equal(t, []int{1, 2, 3}, ToSet([]int{1, 2, 3, 1, 2}))
	assert.Equal(t, []float32{float32(1.1), float32(2.2), float32(3.3), float32(1.11)},
		ToSet([]float32{1.1, 1.1, 2.2, 3.3, 1.11}))

	type st struct {
		I int
		S string
	}
	assert.Equal(t, []st{{1, "one"}, {2, "two"}},
		ToSet([]st{{1, "one"}, {2, "two"}, {1, "one"}}))
}

// nolint: forcetypeassert
func Test_ToSetPred(t *testing.T) {
	// Comparable types
	assert.Equal(t, []int{},
		ToSetPred([]int{}, func(t int) int { return t }))
	assert.Equal(t, []string{"one"},
		ToSetPred([]string{"one"}, func(t string) string { return t }))
	assert.Equal(t, []string{"one", "two", "Two"},
		ToSetPred([]string{"one", "two", "one", "Two"}, func(t string) string { return t }))
	assert.Equal(t, []int{1, 2, 3},
		ToSetPred([]int{1, 2, 3, 1, 2}, func(t int) int { return t }))
	assert.Equal(t, []float32{float32(1.1), float32(2.2), float32(3.3), float32(1.11)},
		ToSetPred([]float32{1.1, 1.1, 2.2, 3.3, 1.11}, func(t float32) float32 { return t }))

	// Incomparable types
	assert.Equal(t, []interface{}{},
		ToSetPred([]interface{}{}, func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{"one"},
		ToSetPred([]interface{}{"one"}, func(t interface{}) string { return t.(string) }))
	assert.Equal(t, []interface{}{"one", "two", "Two"},
		ToSetPred([]interface{}{"one", "two", "one", "Two"}, func(t interface{}) string { return t.(string) }))
	assert.Equal(t, []interface{}{1, 2, 3},
		ToSetPred([]interface{}{1, 2, 3, 1, 2}, func(t interface{}) int { return t.(int) }))
	assert.Equal(t, []interface{}{1.1, 2.2, 3.3, 1.11},
		ToSetPred([]interface{}{1.1, 1.1, 2.2, 3.3, 1.11}, func(t interface{}) float64 { return t.(float64) }))
}
