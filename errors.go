package gofn

import (
	"errors"
	"fmt"
)

var (
	ErrEmpty           = errors.New("container is empty")
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrOverflow        = errors.New("overflow")
	ErrPanic           = errors.New("panic occurred")
)

var (
	// Deprecated: use ErrEmpty instead
	ErrSliceEmpty = ErrEmpty
)

// ErrWrap wraps an error with a message placed in the right
func ErrWrap(err error, msg string) error {
	return fmt.Errorf("%w: %s", err, msg)
}

// ErrWrapL wraps an error with a message placed in the left
func ErrWrapL(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

// ErrUnwrap unwraps an error to get a slice.
// This function can unwrap error created by errors.Join() and fmt.Errorf(<multiple errors passed>).
// In case there's only single item wrapped in the input, the slice has only 1 item.
func ErrUnwrap(err error) []error {
	if err == nil {
		return nil
	}
	u1, ok := err.(interface{ Unwrap() []error })
	if ok {
		return u1.Unwrap()
	}
	if we := errors.Unwrap(err); we != nil {
		return []error{we}
	}
	return nil
}

// ErrUnwrapToRoot unwraps an error until the deepest one
func ErrUnwrapToRoot(err error) error {
	rootErr := err
	for {
		e := errors.Unwrap(rootErr)
		if e == nil {
			return rootErr
		}
		rootErr = e
	}
}
