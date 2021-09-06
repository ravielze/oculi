package errors

import "errors"

var (
	ErrRequestMissingValue   = "required value is missing on key "
	ErrRequestValueNotUUID   = "required value is not uuid on key "
	ErrRequestValueNotBase36 = "required value is not base36 on key "
	ErrKeyNotFound           = errors.New("key not found")
)
