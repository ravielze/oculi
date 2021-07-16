package request

import (
	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/persistent/sql"
	"github.com/ravielze/oculi/request"
)

type (
	reqCtx struct {
		request.Context
		ec echo.Context
	}
)

func New(ec echo.Context, db sql.API) request.EchoContext {
	return &reqCtx{
		Context: request.NewBase(db),
		ec:      ec,
	}
}

func NewWithIdentifier(ec echo.Context, db sql.API, identifier uint64) request.EchoContext {
	return &reqCtx{
		Context: request.NewBaseWithIdentifier(identifier, db),
		ec:      ec,
	}
}

func (r *reqCtx) Echo() echo.Context {
	return r.ec
}
