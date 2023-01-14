package errors

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
	ErrInvalid  = errors.New("invalid")
	ErrInternal = errors.New("internal error")

	ErrUnknown = errors.New("unknown error")
)
