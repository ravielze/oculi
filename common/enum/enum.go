package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type (
	IEnum interface {
		Code() string
		Name() string
	}
)

func Scan(val interface{}, enumArray []IEnum) (int, error) {
	rawValue, ok := val.([]byte)
	if !ok {
		return 0, errors.New("enum error on parsing")
	}
	dbValue := string(rawValue)
	idx := findIndex(dbValue, enumArray, func(e IEnum) string { return e.Code() })
	if idx == 0 {
		return 0, errors.New("enum not found")
	}
	return idx, nil
}

func UnmarshalJSON(val []byte, enumArray []IEnum) (int, error) {
	var rawValue string
	if err := json.Unmarshal(val, &rawValue); err != nil {
		return 0, err
	}

	idx := findIndex(rawValue, enumArray, func(e IEnum) string { return e.Name() })
	if idx == 0 {
		return 0, errors.New("enum not found")
	}
	return idx, nil
}

func Value(enum IEnum) (driver.Value, error) {
	return string(enum.Code()), nil
}

func findIndex(code string, enumArray []IEnum, selector func(e IEnum) string) int {
	for i, v := range enumArray {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}

func MarshalJSON(enum IEnum) ([]byte, error) {
	return json.Marshal(enum.Name())
}
