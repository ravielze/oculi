package context

import (
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/ravielze/oculi/common/radix36"
	stderr "github.com/ravielze/oculi/standard/errors"

	stdcode "github.com/ravielze/oculi/standard/code"
)

// Get parameter with any value
func (ctx *Context) Param(param string) *Context {
	if !ctx.IsError() {

		p := ctx.ginCtx.Param(param)
		if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			ctx.Error(stderr.NewSpecific(param, "missing"), http.StatusBadRequest, stdcode.PARAMETER_ERROR)
		} else {
			ctx.params[param] = p
		}

	}
	return ctx
}

// Get parameter with UUID string value
func (ctx *Context) ParamUUID(param string) *Context {
	if !ctx.IsError() {

		p := ctx.ginCtx.Param(param)
		if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			ctx.Error(
				stderr.NewSpecific(param, "missing"),
				http.StatusBadRequest,
				stdcode.PARAMETER_ERROR,
			)
		} else {
			uuidParsed := uuid.FromStringOrNil(p)
			if strings.EqualFold(p, "default") {
				ctx.params[param] = "default"
			} else if uuidParsed == uuid.Nil {
				ctx.Error(
					stderr.NewSpecific(param, "not_uuid"),
					http.StatusBadRequest,
					stdcode.PARAMETER_ERROR,
				)
			} else {
				ctx.params[param] = uuidParsed.String()
			}
		}

	}
	return ctx
}

// Get parameter with Base36 value
func (ctx *Context) Param36(param string) *Context {
	if !ctx.IsError() {

		p := strings.ToUpper(ctx.ginCtx.Param(param))
		if strings.EqualFold(p, "default") {
			ctx.params[param] = "default"
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			ctx.Error(
				stderr.NewSpecific(param, "missing"),
				http.StatusBadRequest,
				stdcode.PARAMETER_ERROR,
			)
		} else {
			if !radix36.ValidateRadix36(p) {
				ctx.Error(
					stderr.NewSpecific(param, "not_radix36"),
					http.StatusBadRequest,
					stdcode.PARAMETER_ERROR,
				)
			} else {
				ctx.params[param] = p
			}
		}

	}
	return ctx
}

// Get parameter with UUID string value and convert to Base36
func (ctx *Context) ParamUUID36(param string) *Context {
	if !ctx.IsError() {

		p := ctx.ginCtx.Param(param)
		if strings.EqualFold(p, "default") {
			ctx.params[param] = "default"
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			ctx.Error(
				stderr.NewSpecific(param, "missing"),
				http.StatusBadRequest,
				stdcode.PARAMETER_ERROR,
			)
		} else {
			radix36Parsed, err := radix36.EncodeUUID(p)
			if err != nil {
				ctx.Error(
					stderr.NewSpecific(param, "not_uuid"),
					http.StatusBadRequest,
					stdcode.PARAMETER_ERROR,
				)
			} else {
				ctx.params[param] = radix36Parsed
			}
		}

	}
	return ctx
}

// Get parameter with Base36 value and convert to UUID string
func (ctx *Context) Param36UUID(param string) *Context {
	if !ctx.IsError() {

		p := ctx.ginCtx.Param(param)
		if strings.EqualFold(p, "default") {
			ctx.params[param] = "default"
		} else if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
			ctx.Error(
				stderr.NewSpecific(param, "missing"),
				http.StatusBadRequest,
				stdcode.PARAMETER_ERROR,
			)
		} else {
			uuidParsed := radix36.DecodeUUID(strings.ToUpper(p))
			if uuidParsed == uuid.Nil {
				ctx.Error(
					stderr.NewSpecific(param, "not_radix36"),
					http.StatusBadRequest,
					stdcode.PARAMETER_ERROR,
				)
			} else {
				ctx.params[param] = uuidParsed.String()
			}
		}

	}
	return ctx
}
