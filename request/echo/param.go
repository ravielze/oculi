package request

import (
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/ravielze/oculi/common/encoding/radix36"
	"github.com/ravielze/oculi/request"
)

type (
	Stringer interface {
		String() string
	}
)

// Get parameter with any value
func (ctx *reqCtx) Param(param string) request.Context {
	if ctx.ec == nil {
		return ctx
	}
	if !ctx.HasError() {

		p := ctx.ec.Param(param)
		if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			ctx.AddError(http.StatusBadRequest, request.ErrMissingParam)
		} else {
			ctx.ec.Set(request.ParameterKey(param), p)
		}

	}
	return ctx
}

// Get parameter with UUID string value
func (ctx *reqCtx) ParamUUID(param string) request.Context {
	if ctx.ec == nil {
		return ctx
	}
	if !ctx.HasError() {
		p := ctx.ec.Param(param)
		if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			ctx.AddError(http.StatusBadRequest, request.ErrMissingParam)
		} else {
			uuidParsed := uuid.FromStringOrNil(p)
			if strings.EqualFold(p, "default") {
				ctx.ec.Set(request.ParameterKey(param), "default")
			} else if uuidParsed == uuid.Nil {
				ctx.AddError(http.StatusBadRequest, request.ErrParamNotUUID)
			} else {
				ctx.ec.Set(request.ParameterKey(param), uuidParsed.String())
			}
		}
	}
	return ctx
}

// Get parameter with Base36 value
func (ctx *reqCtx) Param36(param string) request.Context {
	if ctx.ec == nil {
		return ctx
	}
	if !ctx.HasError() {
		p := strings.ToUpper(ctx.ec.Param(param))
		if strings.EqualFold(p, "default") {
			ctx.ec.Set(request.ParameterKey(param), "default")
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			ctx.AddError(http.StatusBadRequest, request.ErrMissingParam)
		} else {
			if data, err := radix36.NewRadix36(p); err != nil {
				ctx.AddError(http.StatusBadRequest, request.ErrParamNotBase36)
			} else {
				ctx.ec.Set(request.ParameterKey(param), data.String())
			}
		}

	}
	return ctx
}

// Get parameter with UUID string value and convert to Base36
func (ctx *reqCtx) ParamUUID36(param string) request.Context {
	if ctx.ec == nil {
		return ctx
	}
	if !ctx.HasError() {

		p := ctx.ec.Param(param)
		if strings.EqualFold(p, "default") {
			ctx.ec.Set(request.ParameterKey(param), "default")
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			ctx.AddError(http.StatusBadRequest, request.ErrMissingParam)
		} else {
			data, err := radix36.NewFromUUIDString(p)
			if err != nil {
				ctx.AddError(http.StatusBadRequest, request.ErrParamNotUUID)
			} else {
				ctx.ec.Set(request.ParameterKey(param), data.String())
			}
		}

	}
	return ctx
}

// Get parameter with Base36 value and convert to UUID string
func (ctx *reqCtx) Param36UUID(param string) request.Context {
	if ctx.ec == nil {
		return ctx
	}
	if !ctx.HasError() {

		p := ctx.ec.Param(param)
		if strings.EqualFold(p, "default") {
			ctx.ec.Set(request.ParameterKey(param), "default")
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			ctx.AddError(http.StatusBadRequest, request.ErrMissingParam)
		} else {
			if data, err := radix36.NewRadix36(p); err != nil {
				ctx.AddError(http.StatusBadRequest, request.ErrParamNotBase36)
			} else {
				ctx.ec.Set(request.ParameterKey(param), data.ToUUID().String())
			}
		}

	}
	return ctx
}

func (ctx *reqCtx) AccessParam(param string) string {
	if ctx.ec == nil || ctx.HasError() {
		return ""
	}
	val := ctx.ec.Get(request.ParameterKey(param))
	if val == nil {
		return ""
	}
	if str, ok := val.(string); !ok {
		if stringer, ok := val.(Stringer); ok {
			return stringer.String()
		}
		return ""
	} else {
		return str
	}
}
