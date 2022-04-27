package adapter

import "errors"

var (
	ErrDuplicate    = errors.New("duplicate error")
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
	ErrQuery        = errors.New("query failed")
)
