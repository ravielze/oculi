package errors

import "errors"

var (
	ErrUnknown     = errors.New("unknown error has occured")
	ErrUnclasified = errors.New("unclasified error has occured")
)
