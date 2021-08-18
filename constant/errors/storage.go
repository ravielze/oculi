package errors

import "errors"

var (
	ErrBucketDeleted = errors.New("bucket is already deleted")
)
