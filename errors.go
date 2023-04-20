package helper

import "errors"

var (
	ErrIntegerInvalid = errors.New("integer is invalid")
	ErrNilScanner     = errors.New("nil scanner found")
	ErrInvalidTimes   = errors.New("invalid number of times")
)
