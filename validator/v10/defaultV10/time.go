package defaultV10

import (
	"time"

	v10 "github.com/go-playground/validator/v10"
)

func AfterNow(fl v10.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		return date.After(time.Now())
	}
	return true
}

func BeforeNow(fl v10.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		return date.Before(time.Now())
	}
	return true
}

func IsInFormat(fl v10.FieldLevel) bool {
	// fl.Param()
	// date, ok := fl.Field().Interface().(time.Time)
	// if ok {
	// 	return date.Before(time.Now())
	// }
	// return true
	//TODO
	return false
}
