package context

import (
	"net/http"
	"reflect"

	"github.com/ravielze/oculi/request"
)

var (
	_errType = reflect.TypeOf((*error)(nil)).Elem()
)

func (ctx *Context) BindValidate(obj interface{}) {
	if ctx.HasError() {
		return
	}
	if err := ctx.Bind(obj); err != nil {
		ctx.AddError(http.StatusBadRequest, err)
		return
	}

	if err := ctx.Validate(obj); err != nil {
		ctx.AddError(http.StatusBadRequest, err)
		return
	}
}

func (ctx *Context) Process(usecaseResult ...interface{}) interface{} {
	if ctx.HasError() {
		return nil
	}
	resultLength := len(usecaseResult)
	if resultLength > 0 {
		last := reflect.ValueOf(usecaseResult[resultLength-1])
		if usecaseResult[resultLength-1] == nil {
			return nil
		}
		if last.Type().Implements(_errType) {
			if err, _ := last.Interface().(error); err != nil {
				ctx.AddError(http.StatusBadRequest, err)
				return nil
			}
		}
	}
	return usecaseResult[0]
}

func (ctx *Context) Merge(req request.Context) {
	if ctx.HasError() {
		return
	}
	if req.HasError() {
		ctx.AddError(req.ResponseCode(), req.Error())
	}
}
