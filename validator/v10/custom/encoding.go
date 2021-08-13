package custom

import (
	v10 "github.com/go-playground/validator/v10"
	"github.com/ravielze/oculi/common/baseX/radix36"
	"github.com/ravielze/oculi/validator"
)

type (
	Base36Validator struct {
		*validator.CustomValidator
	}
)

func Base36(tag string) Base36Validator {
	return Base36Validator{
		CustomValidator: validator.NewCustomValidator(tag, "{0} must be in base36 format."),
	}
}

func (Base36Validator) Validate(fl v10.FieldLevel) bool {
	value, ok := fl.Field().Interface().(string)
	if ok {
		return radix36.Validate(value)
	}
	return true
}
