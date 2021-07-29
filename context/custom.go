package context

import (
	"net/http"

	"github.com/ravielze/oculi/request"
)

func (ctx *Context) BindValidate(obj interface{}) {
	if ctx.HasError() {
		return
	}
	if err := ctx.Bind(obj); err != nil {
		ctx.AddError(http.StatusUnprocessableEntity, err)
		return
	}

	if err := ctx.Validate(obj); err != nil {
		ctx.AddError(http.StatusUnprocessableEntity, err)
		return
	}
}

func (ctx *Context) Merge(req request.Context) {
	if ctx.HasError() {
		return
	}
	if req.HasError() {
		ctx.AddError(req.ResponseCode(), req.Error())
	}
}
