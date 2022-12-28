package gofn

import "errors"

var (
	ErrSliceEmpty      = errors.New("slice empty")
	ErrIndexOutOfRange = errors.New("index out of range")
)
