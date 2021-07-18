package custom

import (
	"time"

	v10 "github.com/go-playground/validator/v10"
	"github.com/ravielze/oculi/validator"
)

func AfterNow() (interface{}, validator.FormatOnError, validator.ExtraParamsOnFormat) {
	v := func(fl v10.FieldLevel) bool {
		date, ok := fl.Field().Interface().(time.Time)
		if ok {
			return date.After(time.Now())
		}
		return true
	}
	extraParams := validator.NewExtraParams()
	format := validator.NewFormat("{0} must be after current time.")
	return v, format, extraParams
}

func BeforeNow() (interface{}, validator.FormatOnError, validator.ExtraParamsOnFormat) {
	v := func(fl v10.FieldLevel) bool {
		date, ok := fl.Field().Interface().(time.Time)
		if ok {
			return date.Before(time.Now())
		}
		return true
	}
	extraParams := validator.NewExtraParams()
	format := validator.NewFormat("{0} must be before current time.")
	return v, format, extraParams
}
