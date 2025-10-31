package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tuple2_Unpack(t *testing.T) {
	t1 := Tuple2[int, string]{}
	elem1, elem2 := t1.Unpack()
	assert.Equal(t, 0, elem1)
	assert.Equal(t, "", elem2)

	t2 := &Tuple2[int, string]{Elem1: 1, Elem2: "a"}
	elem1, elem2 = t2.Unpack()
	assert.Equal(t, 1, elem1)
	assert.Equal(t, "a", elem2)
}

func Test_Tuple3_Unpack(t *testing.T) {
	t1 := Tuple3[int, string, *int]{}
	elem1, elem2, elem3 := t1.Unpack()
	assert.Equal(t, 0, elem1)
	assert.Equal(t, "", elem2)
	assert.Nil(t, elem3)

	t2 := &Tuple3[int, string, *int]{Elem1: 1, Elem2: "a", Elem3: ToPtr(2)}
	elem1, elem2, elem3 = t2.Unpack()
	assert.Equal(t, 1, elem1)
	assert.Equal(t, "a", elem2)
	assert.Equal(t, 2, *elem3)
}

func Test_Tuple4_Unpack(t *testing.T) {
	t1 := Tuple4[int, string, *int, bool]{}
	elem1, elem2, elem3, elem4 := t1.Unpack()
	assert.Equal(t, 0, elem1)
	assert.Equal(t, "", elem2)
	assert.Nil(t, elem3)
	assert.False(t, elem4)

	t2 := &Tuple4[int, string, *int, bool]{Elem1: 1, Elem2: "a", Elem3: ToPtr(2), Elem4: true}
	elem1, elem2, elem3, elem4 = t2.Unpack()
	assert.Equal(t, 1, elem1)
	assert.Equal(t, "a", elem2)
	assert.Equal(t, 2, *elem3)
	assert.True(t, elem4)
}

func Test_Tuple5_Unpack(t *testing.T) {
	t1 := Tuple5[int, string, *int, bool, any]{}
	elem1, elem2, elem3, elem4, elem5 := t1.Unpack()
	assert.Equal(t, 0, elem1)
	assert.Equal(t, "", elem2)
	assert.Nil(t, elem3)
	assert.False(t, elem4)
	assert.Equal(t, nil, elem5)

	t2 := &Tuple5[int, string, *int, bool, any]{Elem1: 1, Elem2: "a", Elem3: ToPtr(2), Elem4: true, Elem5: 123}
	elem1, elem2, elem3, elem4, elem5 = t2.Unpack()
	assert.Equal(t, 1, elem1)
	assert.Equal(t, "a", elem2)
	assert.Equal(t, 2, *elem3)
	assert.True(t, elem4)
	assert.Equal(t, 123, elem5)
}
