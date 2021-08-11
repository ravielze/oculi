package functions

import (
	"fmt"
	"reflect"
	"strconv"
)

// Ascii TO Integer
func Atoi(val string, defaultValue uint64) uint64 {
	v, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

func Stringify(val reflect.Value) string {
	stringValue := ""

	switch val.Kind() {
	case reflect.String:
		stringValue = val.Interface().(string)
	case reflect.Bool:
		stringValue = strconv.FormatBool(val.Interface().(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		stringValue = fmt.Sprintf("%d", val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		stringValue = fmt.Sprintf("%d", val.Uint())
	default:
		stringValue = fmt.Sprint(val)
	}

	return stringValue
}
