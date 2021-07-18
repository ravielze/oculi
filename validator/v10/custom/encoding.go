package custom

import (
	v10 "github.com/go-playground/validator/v10"
	"github.com/ravielze/oculi/common/encoding/radix36"
	"github.com/ravielze/oculi/validator"
)

func Base36() (interface{}, validator.FormatOnError, validator.ExtraParamsOnFormat) {
	v := func(fl v10.FieldLevel) bool {
		value, ok := fl.Field().Interface().(string)
		if ok {
			return radix36.Validate(value)
		}
		return true
	}
	extraParams := validator.NewExtraParams()
	format := validator.NewFormat("{0} must be in base36 format.")
	return v, format, extraParams
}
