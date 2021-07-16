package request

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/persistent/sql"
)

var (
	ErrMissingValue   = "required value is missing on key "
	ErrValueNotUUID   = "required value is not uuid on key "
	ErrValueNotBase36 = "required value is not base36 on key "
)

type (
	Context interface {
		SetContext(ctx context.Context) Context
		GetContext() context.Context
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

		ParseString(key, value string) Context
		ParseStringOrDefault(key, value, def string) Context
		ParseUUID(key, value string) Context
		Parse36(key, value string) Context
		ParseUUID36(key, value string) Context
		Parse36UUID(key, value string) Context
		ParseBoolean(key, value string, def bool) Context
		Data() *map[string]string

		Identifier() uint64
	}

	//TODO
	NonEchoContext interface {
		BindValidate(obj interface{})
	}

	EchoContext interface {
		Context

		Echo() echo.Context

		Param(param string) EchoContext
		ParamUUID(param string) EchoContext
		Param36(param string) EchoContext
		ParamUUID36(param string) EchoContext
		Param36UUID(param string) EchoContext

		// Get query with string value and set it to default if it's empty
		Query(query, def string) EchoContext

		// Get query with boolean value
		QueryBoolean(query string, def bool) EchoContext
	}
)
