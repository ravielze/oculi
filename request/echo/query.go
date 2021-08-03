package request

import (
	consts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/request"
)

func (ctx *reqCtx) Query(query, def string) request.EchoReqContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.ParseStringOrDefault(consts.QueryPrefix(query), ctx.ec.QueryParam(query), def)
	return ctx
}

func (ctx *reqCtx) QueryBoolean(query string, def bool) request.EchoReqContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.ParseBoolean(consts.QueryPrefix(query), ctx.ec.QueryParam(query), def)
	return ctx
}
