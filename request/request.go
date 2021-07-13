package request

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/persistent/sql"
)

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
	}
)
