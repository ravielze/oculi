package enum

import (
	"database/sql/driver"
	"reflect"

	consts "github.com/ravielze/oculi/constant/errors"
	"github.com/ravielze/oculi/encoding/jsoniter"
)

type (
	IEnum interface {
		Code() string
		Name() string
	}

	EnumRegisterable interface {
		Name() string
		Code() string
		MarshalJSON() ([]byte, error)
		Value() (driver.Value, error)
	}
	EnumRegisterablePtr interface {
		Scan(val interface{}) error
		UnmarshalJSON(val []byte) error
	}
)

var enumArrayMap = map[string][]IEnum{}
var json = jsoniter.New()

func Register(key string, slice interface{}, objStructPtr interface{}) error {
	if enumArrayMap[key] != nil {
		return consts.ErrEnumKeyRegistered
	}

	valPtr := reflect.ValueOf(objStructPtr)
	if valPtr.Kind() != reflect.Ptr {
		return consts.ErrEnumNotIntPointer
	}

	val := reflect.ValueOf(objStructPtr).Elem()
	if val.Kind() != reflect.Int {
		return consts.ErrEnumNotInt
	}

	if _, ok := objStructPtr.(EnumRegisterablePtr); !ok {
		return consts.ErrEnumImplRegisterablePtr
	}

	if _, ok := val.Interface().(EnumRegisterable); !ok {
		return consts.ErrEnumImplRegisterable
	}
	register(key, slice)
	return nil
}

func register(key string, data interface{}) {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		panic(consts.ErrEnumNotSlice)
	}
	if val.IsNil() {
		return
	}
	arr := make([]IEnum, val.Len())
	for i := 0; i < val.Len(); i++ {
		arr[i] = val.Index(i).Interface().(IEnum)
	}
	enumArrayMap[key] = arr
}

func Scan(val interface{}, key string) (int, error) {
	var dbValue string
	switch v := val.(type) {
	case string:
		dbValue = v
	case []byte:
		dbValue = string(v)
	default:
		return 0, consts.ErrEnumParsing
	}
	idx := findIndex(dbValue, key, func(e IEnum) string { return e.Code() })
	if idx == 0 {
		return 0, consts.ErrEnumNotFound
	}
	return idx, nil
}

func UnmarshalJSON(val []byte, key string) (int, error) {
	var rawValue string
	if err := json.Unmarshal(val, &rawValue); err != nil {
		return 0, err
	}

	idx := findIndex(rawValue, key, func(e IEnum) string { return e.Name() })
	if idx == 0 {
		return 0, consts.ErrEnumNotFound
	}
	return idx, nil
}

func Value(enum IEnum) (driver.Value, error) {
	return string(enum.Code()), nil
}

func findIndex(x string, key string, selector func(e IEnum) string) int {
	if enumArrayMap[key] == nil {
		return 0
	}
	for i, v := range enumArrayMap[key] {
		if selector(v) == x {
			return i + 1
		}
	}
	return 0
}

func MarshalJSON(enum IEnum) ([]byte, error) {
	return json.Marshal(enum.Name())
}
