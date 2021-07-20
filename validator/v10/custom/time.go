package custom

import (
	"time"

	v10 "github.com/go-playground/validator/v10"
	"github.com/ravielze/oculi/validator"
)

type (
	AfterNowValidator struct {
		*validator.CustomValidator
	}
)

func AfterNow(tag string) AfterNowValidator {
	return AfterNowValidator{
		CustomValidator: validator.NewCustomValidator(tag, "{0} must be after current time."),
	}
}

func (AfterNowValidator) Validate(fl v10.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		return date.After(time.Now())
	}
	return true
}

// func BeforeNow() (interface{}, validator.FormatOnError, validator.ExtraParamsOnFormat) {
// 	v := func(fl v10.FieldLevel) bool {
// 		date, ok := fl.Field().Interface().(time.Time)
// 		if ok {
// 			return date.Before(time.Now())
// 		}
// 		return true
// 	}
// 	extraParams := validator.NewExtraParams()
// 	format := validator.NewFormat("{0} must be before current time.")
// 	return v, format, extraParams
// }
