package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"reflect"
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

var (
	ErrEnumKeyRegistered       = errors.New("enum's key already registered")
	ErrEnumNotInt              = errors.New("enum is not int")
	ErrEnumNotIntPointer       = errors.New("enum is not int pointer")
	ErrEnumImplRegisterable    = errors.New("enum need implement registerable")
	ErrEnumImplRegisterablePtr = errors.New("enum need implement registerable pointer")
	ErrEnumNotSlice            = errors.New("enum data is not slice")
	ErrEnumNotFound            = errors.New("enum not found")
	ErrEnumParsing             = errors.New("failed to parse enum")
)

var enumArrayMap = map[string][]IEnum{}

func Register(key string, slice interface{}, objStruct interface{}, objStructPtr interface{}) error {
	if enumArrayMap[key] != nil {
		return ErrEnumKeyRegistered
	}

	val := reflect.ValueOf(objStruct)
	if val.Kind() != reflect.Int {
		return ErrEnumNotInt
	}
	if _, ok := objStruct.(EnumRegisterable); !ok {
		return ErrEnumImplRegisterable
	}

	valPtr := reflect.ValueOf(objStructPtr)
	if valPtr.Kind() != reflect.Ptr {
		return ErrEnumNotIntPointer
	}
	if _, ok := objStructPtr.(EnumRegisterablePtr); !ok {
		return ErrEnumImplRegisterablePtr
	}
	register(key, slice)
	return nil
}

func register(key string, data interface{}) {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		panic(ErrEnumNotSlice)
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
	rawValue, ok := val.([]byte)
	if !ok {
		return 0, ErrEnumParsing
	}
	dbValue := string(rawValue)
	idx := findIndex(dbValue, key, func(e IEnum) string { return e.Code() })
	if idx == 0 {
		return 0, ErrEnumNotFound
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
		return 0, ErrEnumNotFound
	}
	return idx, nil
}

func Value(enum IEnum) (driver.Value, error) {
	return string(enum.Code()), nil
}

func findIndex(code string, key string, selector func(e IEnum) string) int {
	if enumArrayMap[key] == nil {
		return 0
	}
	for i, v := range enumArrayMap[key] {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}

func MarshalJSON(enum IEnum) ([]byte, error) {
	return json.Marshal(enum.Name())
}
