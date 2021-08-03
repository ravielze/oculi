package request

import (
	consts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/request"
)

type (
	Stringer interface {
		String() string
	}
)

// Get parameter with any value
func (ctx *reqCtx) Param(param string) request.EchoReqContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.ParseString(consts.ParameterPrefix(param), ctx.ec.Param(param))
	return ctx
}

// Get parameter with UUID string value
func (ctx *reqCtx) ParamUUID(param string) request.EchoReqContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.ParseUUID(consts.ParameterPrefix(param), ctx.ec.Param(param))
	return ctx
}

// Get parameter with Base36 value
func (ctx *reqCtx) Param36(param string) request.EchoReqContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.Parse36(consts.ParameterPrefix(param), ctx.ec.Param(param))
	return ctx
}

// Get parameter with UUID string value and convert to Base36
func (ctx *reqCtx) ParamUUID36(param string) request.EchoReqContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.ParseUUID36(consts.ParameterPrefix(param), ctx.ec.Param(param))
	return ctx
}

// Get parameter with Base36 value and convert to UUID string
func (ctx *reqCtx) Param36UUID(param string) request.EchoReqContext {
	if ctx.ec == nil {
		return ctx
	}
	ctx.Parse36UUID(consts.ParameterPrefix(param), ctx.ec.Param(param))
	return ctx
}
