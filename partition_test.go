package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
