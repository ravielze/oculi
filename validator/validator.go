package validator

import "reflect"

type (
	Validator interface {
		Validate(obj interface{}) error
		ValidateVar(obj interface{}, tag string) error

		RegisterValidation(tag string, fn interface{})
		RegisterCustomTypeFunc(fn CustomTypeFunc, types ...interface{})
		RegisterStructValidation(fn interface{}, types ...interface{})
	}

	CustomTypeFunc func(field reflect.Value) interface{}
)
