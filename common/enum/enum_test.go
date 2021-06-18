package enum

import (
	"database/sql/driver"
	"testing"
)

type (
	testEnumClass struct {
		name string
		code string
	}

	TestEnum int
)

const (
	Waiting TestEnum = iota + 1
	Confirmed
	OnGoing
)

var enumData = []IEnum{
	testEnumClass{"Waiting", "waiting"},
	testEnumClass{"Confirmed", "confirmed"},
	testEnumClass{"On Going", "on_going"},
}

func (t testEnumClass) Name() string {
	return t.name
}

func (t testEnumClass) Code() string {
	return t.code
}

func (t TestEnum) Name() string {
	if int(t) < 1 || int(t) > len(enumData) {
		return ""
	}
	return enumData[int(1)-1].Name()
}

func (t TestEnum) Code() string {
	if int(t) < 1 || int(t) > len(enumData) {
		return ""
	}
	return enumData[int(1)-1].Code()
}

func (t TestEnum) MarshalJSON() ([]byte, error) {
	return MarshalJSON(t)
}

func (t *TestEnum) UnmarshalJSON(val []byte) error {
	idx, err := UnmarshalJSON(val, enumData)
	if err != nil {
		return err
	}
	*t = TestEnum(idx)
	return nil
}

func (t *TestEnum) Scan(val interface{}) error {
	idx, err := Scan(val, enumData)
	if err != nil {
		return err
	}
	*t = TestEnum(idx)
	return nil
}

func (t TestEnum) Value() (driver.Value, error) {
	return Value(t)
}

func TestScan(t *testing.T) {

}
