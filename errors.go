package gofn

import "errors"

var (
	ErrEmpty           = errors.New("container is empty")
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrPanic           = errors.New("panic occurred")
)

var (
	// Deprecated: use ErrEmpty instead
	ErrSliceEmpty = ErrEmpty
)
