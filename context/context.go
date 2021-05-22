package context

import (
	"github.com/gin-gonic/gin"
	std "github.com/ravielze/oculi/standard"
)

type (
	Parameters map[string]string
	Queries    map[string]string
	Context    struct {
		ginCtx *gin.Context

		err      error
		params   Parameters
		code     string
		httpCode int
		query    Queries
		isError  bool
	}
)

func New(ctx *gin.Context) *Context {
	return &Context{
		ginCtx:   ctx,
		isError:  false,
		err:      nil,
		params:   Parameters{},
		code:     "",
		httpCode: -1,
		query:    Queries{},
	}
}

func (ctx *Context) IsError() bool {
	return ctx.isError
}

func (ctx *Context) Error(err error, httpCode int, code std.Code) {
	if !ctx.IsError() {
		ctx.err = err
		ctx.httpCode = httpCode
		ctx.code = string(code)
		ctx.isError = true
	}
}
