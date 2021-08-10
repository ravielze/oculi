package request

import (
	"reflect"
	"unsafe"

	"github.com/labstack/echo/v4"
	uDto "github.com/ravielze/oculi/common/model/dto/user"
	consts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/persistent/sql"
	"github.com/ravielze/oculi/request"
)

type (
	reqCtx struct {
		request.ReqContext
		ec echo.Context
	}
)

func New(ec echo.Context, db sql.API) request.EchoReqContext {
	r := &reqCtx{
		ReqContext: request.NewBase(db),
		ec:         ec,
	}
	if item := ec.Get(consts.KeyCredentials); item != nil {
		if c, ok := item.(uDto.CredentialsDTO); ok {
			r.WithIdentifier(c.ID)
		}
	}
	return r
}

func (r *reqCtx) Echo() echo.Context {
	return r.ec
}

func getUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

func (r *reqCtx) Transform() request.ReqContext {
	result := r.ReqContext
	// r.ec is echo.Context (implemented in oculi.Context)
	// Elem() get the implemented (oculi.Context)
	// FieldByName("ec") get field ec from echo.Context (implemented in *echo.context)
	// Elem() transform echo.Context to *echo.context
	// Second Elem() transform *echo.context to echo.context
	// FieldByName("store") is where echo context store located
	store := reflect.ValueOf(r.ec).Elem().FieldByName("ec").Elem().Elem().FieldByName("store")
	data := *r.Data()
	if !store.IsNil() {
		echoStore := getUnexportedField(store).(echo.Map)
		for k, v := range echoStore {
			data[consts.EchoPrefix(k)] = v
		}
	}
	data[consts.EchoContext] = r.ec
	return result
}
