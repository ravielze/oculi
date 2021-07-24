package rest

import (
	"reflect"

	"github.com/labstack/echo/v4"
)

type (
	Registerable interface {
		Register(ec *echo.Group) error
	}
)

func Register(ec *echo.Group, obj interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Type().Name() == "In" {
			continue
		}
		if r, ok := f.Interface().(Registerable); ok {
			if err := r.Register(ec); err != nil {
				return err
			}
		}
	}
	return nil
}
