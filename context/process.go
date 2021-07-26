package context

import (
	"reflect"

	"github.com/ravielze/oculi/errors"
)

var (
	_errType = reflect.TypeOf((*error)(nil)).Elem()
)

type (
	Function struct {
		function interface{}
		args     []interface{}
	}
)

func NewFunction(function interface{}, args ...interface{}) *Function {
	return &Function{
		function: function,
		args:     args,
	}
}

// Process/Run a function f if context doesn't have error, map the error to setup the status code. Run function onError when error.
func (ctx *Context) Process(fn *Function, onError func(), errorMap errors.Mappers) interface{} {
	if ctx.HasError() {
		return nil
	}
	argsValue := make([]reflect.Value, len(fn.args))
	for i := range argsValue {
		argsValue[i] = reflect.ValueOf(fn.args[i])
	}
	fnValue := reflect.ValueOf(fn.function)
	fnResults := fnValue.Call(argsValue)
	resultLen := len(fnResults)
	if resultLen > 0 {
		last := fnResults[resultLen-1]
		if fnResults[resultLen-1].Interface() == nil {
			return fnResults[0].Interface()
		}
		if last.Type().Implements(_errType) {
			if err, _ := last.Interface().(error); err != nil {
				errTrans := errors.Transform(err, errorMap)
				ctx.AddError(errTrans.Code, errTrans)
				if onError != nil {
					onError()
				}
				return nil
			}
		}
		return fnResults[0].Interface()
	}
	return nil
}
