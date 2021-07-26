package context

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Context struct {
		ec  echo.Context
		ctx context.Context

		errors   []error
		httpCode int
	}
)

func New(ec echo.Context) *Context {
	return &Context{
		ctx:      context.Background(),
		ec:       ec,
		errors:   make([]error, 0),
		httpCode: http.StatusOK,
	}
}

func (ctx *Context) Context() context.Context {
	return ctx.ctx
}
