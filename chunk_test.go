package gofn

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Chunk(t *testing.T) {
	// Empty input
	chunks := Chunk([]int{}, 5)
	assert.True(t, len(chunks) == 0)

	// Nil input
	chunks = Chunk[int](nil, 5)
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
	chunks = ChunkByPieces[int](nil, 5)
	assert.True(t, len(chunks) == 0)

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
