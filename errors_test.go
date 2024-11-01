package gofn

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrWrap(t *testing.T) {
	e := errors.New("err")
	assert.Equal(t, "err: msg", ErrWrap(e, "msg").Error())
	assert.Equal(t, "msg: err", ErrWrapL("msg", e).Error())
}

type wrappedErrs struct { //nolint:errname
	errs []error
}

func (we *wrappedErrs) Error() string {
	return ""
}

func (we *wrappedErrs) Unwrap() []error {
	return we.errs
}

func Test_ErrUnwrap(t *testing.T) {
	e1 := errors.New("e1")
	e2 := fmt.Errorf("e2: %w", e1)
	e3 := &wrappedErrs{errs: []error{e1, e2}}

	assert.Equal(t, []error(nil), ErrUnwrap(nil))
	assert.Equal(t, []error(nil), ErrUnwrap(e1))
	assert.Equal(t, []error{e1}, ErrUnwrap(e2))
	assert.Equal(t, []error{e1, e2}, ErrUnwrap(e3))
}

func Test_ErrUnwrapToRoot(t *testing.T) {
	e1 := errors.New("e1")
	e2 := fmt.Errorf("e2: %w", e1)
	e3 := fmt.Errorf("e3: %w", e2)

	assert.Equal(t, nil, ErrUnwrapToRoot(nil))
	assert.Equal(t, e1, ErrUnwrapToRoot(e2))
	assert.Equal(t, e1, ErrUnwrapToRoot(e3))
}
