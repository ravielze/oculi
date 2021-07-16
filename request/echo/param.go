package request

import (
	"github.com/ravielze/oculi/request"
)

type (
	Stringer interface {
		String() string
	}
)

// Get parameter with any value
func (ctx *reqCtx) Param(param string) request.EchoContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.ParseString("parameter."+param, ctx.ec.Param(param))
	return ctx
}

// Get parameter with UUID string value
func (ctx *reqCtx) ParamUUID(param string) request.EchoContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.ParseUUID("parameter."+param, ctx.ec.Param(param))
	return ctx
}

// Get parameter with Base36 value
func (ctx *reqCtx) Param36(param string) request.EchoContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.Parse36("parameter."+param, ctx.ec.Param(param))
	return ctx
}

// Get parameter with UUID string value and convert to Base36
func (ctx *reqCtx) ParamUUID36(param string) request.EchoContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.ParseUUID36("parameter."+param, ctx.ec.Param(param))
	return ctx
}

// Get parameter with Base36 value and convert to UUID string
func (ctx *reqCtx) Param36UUID(param string) request.EchoContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.Parse36UUID("parameter."+param, ctx.ec.Param(param))
	return ctx
}
