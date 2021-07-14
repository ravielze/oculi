package request

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/persistent/sql"
)

var (
	ErrMissingParam   = errors.New("required parameter is missing")
	ErrParamNotUUID   = errors.New("required parameter is not uuid")
	ErrParamNotBase36 = errors.New("required parameter is not base36")
)

func ParameterKey(key string) string {
	return "parameter-" + key
}

type (
	Context interface {
		Context() context.Context
		Echo() echo.Context

		HasError() bool
		AddError(responseCode int, err ...error)

		SetResponseCode(code int)
		ResponseCode() int
		Error() error

		HasTransaction() bool
		Transaction() sql.API
		NewTransaction() sql.API
		CommitTransaction() sql.API
		RollbackTransaction() sql.API

		Param(param string) Context
		ParamUUID(param string) Context
		Param36(param string) Context
		ParamUUID36(param string) Context
		Param36UUID(param string) Context
	}
)
