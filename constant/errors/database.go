package errors

import "errors"

var (
	ErrRecordNotFound        = errors.New("record not found")
	ErrInvalidTransaction    = errors.New("no valid transaction")
	ErrNotImplemented        = errors.New("not implemented")
	ErrMissingWhereClause    = errors.New("WHERE conditions required")
	ErrUnsupportedRelation   = errors.New("unsupported relations")
	ErrPrimaryKeyRequired    = errors.New("primary key required")
	ErrModelValueRequired    = errors.New("model value required")
	ErrInvalidData           = errors.New("unsupported data")
	ErrUnsupportedDriver     = errors.New("unsupported driver")
	ErrRegistered            = errors.New("registered")
	ErrInvalidField          = errors.New("invalid field")
	ErrEmptySlice            = errors.New("empty slice found")
	ErrConflictData          = errors.New("data is conflicted")
	ErrDryRunModeUnsupported = errors.New("dry run mode unsupported")
	ErrInvalidDB             = errors.New("invalid database")
	ErrInvalidValue          = errors.New("invalid value, should be pointer to struct or slice")
	ErrInvalidValueOfLength  = errors.New("invalid association values, length doesn't match")
)
