package errors

import "errors"

var (
	ErrBcryptInvalidCost = errors.New("bcrypt invalid cost")
	ErrPasswordMismatch  = errors.New("password mismatch")
)
