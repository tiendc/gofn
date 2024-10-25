package gofn

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrUnwrap(t *testing.T) {
	e1 := errors.New("e1")
	e2 := fmt.Errorf("e2: %w", e1)
	e3 := fmt.Errorf("e3: %w - %w", e1, e2) // NOTE: errors.Join() is unavailable in Go prior to 1.20

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
