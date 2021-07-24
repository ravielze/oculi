package context

import "net/http"

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
