package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Product(t *testing.T) {
	assert.Equal(t, 0, Product[int]())
	assert.Equal(t, -6, Product(1, 2, 3, -1))
	assert.Equal(t, 0, Product(1, 2, 3, -1, 0))
	assert.Equal(t, int8(-6), Product[int8](1, 2, 3, -1))
}

func Test_ProductAs(t *testing.T) {
	assert.Equal(t, 0, ProductAs[int, int]())
	assert.Equal(t, -6, ProductAs[int](1, 2, 3, -1))
	assert.Equal(t, 0, ProductAs[int](1, 2, 3, -1, 0))
	assert.Equal(t, int64(6000000000), ProductAs[int64](int32(1000), int32(2000), int32(3000)))
	// Overflow
	assert.Equal(t, int32(1705032704), ProductAs[int32](int32(1000), int32(2000), int32(3000)))
}
