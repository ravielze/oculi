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

		result   interface{}
		errors   []error
		httpCode int
		skipAuth bool
	}
)

func New(ec echo.Context) *Context {
	return &Context{
		ctx:      context.Background(),
		ec:       ec,
		errors:   make([]error, 0),
		httpCode: http.StatusOK,
		skipAuth: false,
	}
}

func (ctx *Context) Context() context.Context {
	return ctx.ctx
}

func (ctx *Context) NeedAuth() bool {
	return !ctx.skipAuth
}

func (ctx *Context) SkipAuth(state bool) {
	ctx.skipAuth = state
}
