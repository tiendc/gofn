package gofn

import "errors"

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
