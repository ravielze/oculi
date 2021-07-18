package validator

import (
	"reflect"

	ut "github.com/go-playground/universal-translator"
)

type (
	Validator interface {
		Validate(obj interface{}) error
		ValidateVar(obj interface{}, tag string) error

		RegisterValidation(tag string, fn interface{})
		RegisterCustomTypeFunc(fn CustomTypeFunc, types ...interface{})
		RegisterStructValidation(fn interface{}, types ...interface{})

		Translator() *ut.Translator

		//Translate error into custom error message for the provided tag with the provided format.
		//{0} will be the field name.
		//{1} will be the validator param.
		//{2}, {3}, ... will extraParams[0], extraParams[1] and so on.
		AddTranslation(tag string, format string, extraParams ...string) error

		Register(tag string, vr ValidatorRegisterable) error
	}

	FormatOnError         string
	ExtraParamsOnFormat   []string
	ValidatorRegisterable func() (interface{}, FormatOnError, ExtraParamsOnFormat)

	CustomTypeFunc func(field reflect.Value) interface{}
)

func NewFormat(val string) FormatOnError {
	return FormatOnError(val)
}

func NewExtraParams(val ...string) ExtraParamsOnFormat {
	if len(val) == 0 {
		return ExtraParamsOnFormat([]string{})
	}
	return ExtraParamsOnFormat(val)
}
