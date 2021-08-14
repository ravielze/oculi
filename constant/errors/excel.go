package consts

import "errors"

var (
	ErrCellOutOfRange = errors.New("cell out of range")
	ErrStyleNotFound  = errors.New("style id not found")
)
