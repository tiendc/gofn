package gofn

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_MinTime(t *testing.T) {
	t0 := time.Time{}
	t1 := time.Now()
	t2 := t1.Add(time.Second)
	t3 := t1.Add(-time.Minute)

	assert.Equal(t, t0, MinTime(t0, t1, t2, t3))
	assert.Equal(t, t3, MinTime(t1, t2, t3))
}

func Test_MaxTime(t *testing.T) {
	t0 := time.Time{}
	t1 := time.Now()
	t2 := t1.Add(time.Second)
	t3 := t1.Add(-time.Minute)

	assert.Equal(t, t2, MaxTime(t0, t1, t2, t3))
	assert.Equal(t, t2, MaxTime(t1, t2, t3))
}

func Test_MinMaxTime(t *testing.T) {
	t0 := time.Time{}
	t1 := time.Now()
	t2 := t1.Add(time.Second)
	t3 := t1.Add(-time.Minute)

	m1, m2 := MinMaxTime(t0, t1, t2, t3)
	assert.Equal(t, t0, m1)
	assert.Equal(t, t2, m2)

	m1, m2 = MinMaxTime(t1, t2, t3)
	assert.Equal(t, t3, m1)
	assert.Equal(t, t2, m2)
}
