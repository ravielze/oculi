package v10

import (
	"github.com/gin-gonic/gin/binding"
	v10 "github.com/go-playground/validator/v10"
	"github.com/ravielze/oculi/constant/errors"
	"github.com/ravielze/oculi/validator"
)

type (
	impl struct {
		instance *v10.Validate
	}
)

func New() (validator.Validator, error) {
	if v, ok := binding.Validator.Engine().(*v10.Validate); ok {
		return &impl{instance: v}, nil
	}
	return nil, errors.ErrValidatorEngineNotFound
}

func (i *impl) RegisterAlias(alias string, tags string) {
	i.instance.RegisterAlias(alias, tags)
}

func (i *impl) RegisterCustomTypeFunc(fn validator.CustomTypeFunc, types ...interface{}) {
	i.instance.RegisterCustomTypeFunc(v10.CustomTypeFunc(fn), types...)
}

func (i *impl) RegisterValidation(tag string, fn interface{}) {
	if fnConv, ok := fn.(v10.Func); !ok {
		panic("fn is not v10.Func")
	} else {
		i.instance.RegisterValidation(tag, fnConv)
	}
}

func (i *impl) RegisterStructValidation(fn interface{}, types ...interface{}) {
	if fnConv, ok := fn.(v10.StructLevelFunc); !ok {
		panic("fn is not v10.StructLevelFunc")
	} else {
		i.instance.RegisterStructValidation(fnConv, types...)
	}
}

func (i *impl) Validate(object interface{}) error {
	return i.instance.Struct(object)
}

func (i *impl) InstallDefault() {
}
