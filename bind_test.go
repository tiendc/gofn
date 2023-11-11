package gofn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bind1Arg(t *testing.T) {
	fn1Arg0Ret := func(a1 int) {
		assert.True(t, a1 == 1)
	}
	fn1Arg1Ret := func(a1 int) int {
		assert.True(t, a1 == 1)
		return 1
	}
	fn1Arg2Ret := func(a1 int) (int, int) {
		assert.True(t, a1 == 1)
		return 1, 2
	}
	fn1Arg3Ret := func(a1 int) (int, int, int) {
		assert.True(t, a1 == 1)
		return 1, 2, 3
	}

	Bind1Arg0Ret(fn1Arg0Ret, 1)()
	r1 := Bind1Arg1Ret(fn1Arg1Ret, 1)()
	assert.True(t, r1 == 1)
	r1, r2 := Bind1Arg2Ret(fn1Arg2Ret, 1)()
	assert.True(t, r1 == 1 && r2 == 2)
	r1, r2, r3 := Bind1Arg3Ret(fn1Arg3Ret, 1)()
	assert.True(t, r1 == 1 && r2 == 2 && r3 == 3)
}

func Test_Bind2Arg(t *testing.T) {
	fn2Arg0Ret := func(a1 int, a2 int) {
		assert.True(t, a1 == 1 && a2 == 2)
	}
	fn2Arg1Ret := func(a1 int, a2 int) int {
		assert.True(t, a1 == 1 && a2 == 2)
		return 1
	}
	fn2Arg2Ret := func(a1 int, a2 int) (int, int) {
		assert.True(t, a1 == 1 && a2 == 2)
		return 1, 2
	}
	fn2Arg3Ret := func(a1 int, a2 int) (int, int, int) {
		assert.True(t, a1 == 1 && a2 == 2)
		return 1, 2, 3
	}

	Bind2Arg0Ret(fn2Arg0Ret, 1, 2)()
	r1 := Bind2Arg1Ret(fn2Arg1Ret, 1, 2)()
	assert.True(t, r1 == 1)
	r1, r2 := Bind2Arg2Ret(fn2Arg2Ret, 1, 2)()
	assert.True(t, r1 == 1 && r2 == 2)
	r1, r2, r3 := Bind2Arg3Ret(fn2Arg3Ret, 1, 2)()
	assert.True(t, r1 == 1 && r2 == 2 && r3 == 3)
}

func Test_Bind3Arg(t *testing.T) {
	fn3Arg0Ret := func(a1 int, a2 int, a3 int) {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3)
	}
	fn3Arg1Ret := func(a1 int, a2 int, a3 int) int {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3)
		return 1
	}
	fn3Arg2Ret := func(a1 int, a2 int, a3 int) (int, int) {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3)
		return 1, 2
	}
	fn3Arg3Ret := func(a1 int, a2 int, a3 int) (int, int, int) {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3)
		return 1, 2, 3
	}

	Bind3Arg0Ret(fn3Arg0Ret, 1, 2, 3)()
	r1 := Bind3Arg1Ret(fn3Arg1Ret, 1, 2, 3)()
	assert.True(t, r1 == 1)
	r1, r2 := Bind3Arg2Ret(fn3Arg2Ret, 1, 2, 3)()
	assert.True(t, r1 == 1 && r2 == 2)
	r1, r2, r3 := Bind3Arg3Ret(fn3Arg3Ret, 1, 2, 3)()
	assert.True(t, r1 == 1 && r2 == 2 && r3 == 3)
}

func Test_Bind4Arg(t *testing.T) {
	fn4Arg0Ret := func(a1 int, a2 int, a3 int, a4 int) {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3 && a4 == 4)
	}
	fn4Arg1Ret := func(a1 int, a2 int, a3 int, a4 int) int {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3 && a4 == 4)
		return 1
	}
	fn4Arg2Ret := func(a1 int, a2 int, a3 int, a4 int) (int, int) {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3 && a4 == 4)
		return 1, 2
	}
	fn4Arg3Ret := func(a1 int, a2 int, a3 int, a4 int) (int, int, int) {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3 && a4 == 4)
		return 1, 2, 3
	}

	Bind4Arg0Ret(fn4Arg0Ret, 1, 2, 3, 4)()
	r1 := Bind4Arg1Ret(fn4Arg1Ret, 1, 2, 3, 4)()
	assert.True(t, r1 == 1)
	r1, r2 := Bind4Arg2Ret(fn4Arg2Ret, 1, 2, 3, 4)()
	assert.True(t, r1 == 1 && r2 == 2)
	r1, r2, r3 := Bind4Arg3Ret(fn4Arg3Ret, 1, 2, 3, 4)()
	assert.True(t, r1 == 1 && r2 == 2 && r3 == 3)
}

func Test_Bind5Arg(t *testing.T) {
	fn5Arg0Ret := func(a1 int, a2 int, a3 int, a4 int, a5 int) {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3 && a4 == 4 && a5 == 5)
	}
	fn5Arg1Ret := func(a1 int, a2 int, a3 int, a4 int, a5 int) int {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3 && a4 == 4 && a5 == 5)
		return 1
	}
	fn5Arg2Ret := func(a1 int, a2 int, a3 int, a4 int, a5 int) (int, int) {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3 && a4 == 4 && a5 == 5)
		return 1, 2
	}
	fn5Arg3Ret := func(a1 int, a2 int, a3 int, a4 int, a5 int) (int, int, int) {
		assert.True(t, a1 == 1 && a2 == 2 && a3 == 3 && a4 == 4 && a5 == 5)
		return 1, 2, 3
	}

	Bind5Arg0Ret(fn5Arg0Ret, 1, 2, 3, 4, 5)()
	r1 := Bind5Arg1Ret(fn5Arg1Ret, 1, 2, 3, 4, 5)()
	assert.True(t, r1 == 1)
	r1, r2 := Bind5Arg2Ret(fn5Arg2Ret, 1, 2, 3, 4, 5)()
	assert.True(t, r1 == 1 && r2 == 2)
	r1, r2, r3 := Bind5Arg3Ret(fn5Arg3Ret, 1, 2, 3, 4, 5)()
	assert.True(t, r1 == 1 && r2 == 2 && r3 == 3)
}
