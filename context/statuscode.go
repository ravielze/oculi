package context

import (
	"net/http"

	"github.com/ravielze/oculi/errors"
)

func (ctx *Context) SetHttpCode(httpCode int) {
	ctx.httpCode = httpCode
}

func (ctx *Context) ResponseCode() int {
	return ctx.httpCode
}

func (ctx *Context) AddError(httpCode int, err ...error) {
	if ctx.httpCode < 400 {
		ctx.httpCode = httpCode
	}
	ctx.errors = append(ctx.errors, err...)
}

func (ctx *Context) ClearErrors() {
	ctx.errors = make([]error, 0)
	ctx.httpCode = http.StatusOK
}

func (ctx *Context) HasError() bool {
	return len(ctx.errors) > 0
}

func (ctx *Context) Errors() []error {
	return ctx.errors
}

func (ctx *Context) TransformError(errorMap errors.Mappers) {
	if ctx.HasError() {
		err := errors.Transform(ctx.errors[0], errorMap)
		ctx.errors = append([]error{err}, ctx.errors...)
		ctx.httpCode = err.Code
	}
}
