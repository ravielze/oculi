package errors

import "errors"

var (
	ErrEnumKeyRegistered       = errors.New("enum's key already registered")
	ErrEnumNotInt              = errors.New("enum is not int")
	ErrEnumNotIntPointer       = errors.New("enum is not int pointer")
	ErrEnumImplRegisterable    = errors.New("enum need implement registerable")
	ErrEnumImplRegisterablePtr = errors.New("enum need implement registerable pointer")
	ErrEnumNotSlice            = errors.New("enum data is not slice")
	ErrEnumNotFound            = errors.New("enum not found")
	ErrEnumParsing             = errors.New("failed to parse enum")
)
