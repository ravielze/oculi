package context

import (
	"net/http"

	stdcode "github.com/ravielze/oculi/standard/code"
)

func (ctx *Context) BindJSON(obj interface{}) *Context {
	if !ctx.IsError() {
		if err := ctx.ginCtx.ShouldBindJSON(obj); err != nil {
			ctx.Error(
				err,
				http.StatusUnprocessableEntity,
				stdcode.UNPROCESSABLE_ENTITY,
			)
		}
	}
	return ctx
}

func (ctx *Context) BindForm(obj interface{}) *Context {
	if !ctx.IsError() {
		if err := ctx.ginCtx.ShouldBind(obj); err != nil {
			ctx.Error(
				err,
				http.StatusUnprocessableEntity,
				stdcode.UNPROCESSABLE_ENTITY,
			)
		}
	}
	return ctx
}
