package request

import (
	"github.com/ravielze/oculi/request"
)

func (ctx *reqCtx) Query(query, def string) request.EchoContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.ParseStringOrDefault("query."+query, ctx.ec.QueryParam(query), def)
	return ctx
}

func (ctx *reqCtx) QueryBoolean(query string, def bool) request.EchoContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.ParseBoolean("query."+query, ctx.ec.QueryParam(query), def)
	return ctx
}
