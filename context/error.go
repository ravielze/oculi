package context

import std "github.com/ravielze/oculi/standard"

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
