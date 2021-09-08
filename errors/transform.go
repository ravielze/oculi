package errors

import (
	"errors"

	"github.com/minio/minio-go/v7"
	consts "github.com/ravielze/oculi/constant/errors"
	"gorm.io/gorm"
)

func Convert(err error) error {
	if convertedErr, ok := err.(minio.ErrorResponse); ok {
		storageErr, ok := consts.StorageCodeErrorMapping[convertedErr.Code]
		if ok {
			return storageErr
		}
	}

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return consts.ErrRecordNotFound
	case errors.Is(err, gorm.ErrInvalidTransaction):
		return consts.ErrInvalidTransaction
	case errors.Is(err, gorm.ErrNotImplemented):
		return consts.ErrNotImplemented
	case errors.Is(err, gorm.ErrMissingWhereClause):
		return consts.ErrMissingWhereClause
	case errors.Is(err, gorm.ErrUnsupportedRelation):
		return consts.ErrUnsupportedRelation
	case errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return consts.ErrPrimaryKeyRequired
	case errors.Is(err, gorm.ErrModelValueRequired):
		return consts.ErrModelValueRequired
	case errors.Is(err, gorm.ErrInvalidData):
		return consts.ErrInvalidData
	case errors.Is(err, gorm.ErrUnsupportedDriver):
		return consts.ErrUnsupportedDriver
	case errors.Is(err, gorm.ErrRegistered):
		return consts.ErrRegistered
	case errors.Is(err, gorm.ErrInvalidField):
		return consts.ErrInvalidField
	case errors.Is(err, gorm.ErrEmptySlice):
		return consts.ErrEmptySlice
	case errors.Is(err, gorm.ErrDryRunModeUnsupported):
		return consts.ErrDryRunModeUnsupported
	case errors.Is(err, gorm.ErrInvalidDB):
		return consts.ErrInvalidDB
	case errors.Is(err, gorm.ErrInvalidValue):
		return consts.ErrInvalidValue
	case errors.Is(err, gorm.ErrInvalidValueOfLength):
		return consts.ErrInvalidValueOfLength
	}

	return err
}
