package enum

import (
	"database/sql/driver"
	"testing"

	"github.com/ravielze/oculi/constant/errors"
	"github.com/stretchr/testify/assert"
)

type (
	testEnumClass struct {
		name string
		code string
	}

	TestEnum  int
	WrongEnum int
)

const (
	Waiting TestEnum = iota + 1
	Confirmed
	OnGoing
)

var enumData = []testEnumClass{
	{"Waiting", "waiting"},
	{"Confirmed", "confirmed"},
	{"On Going", "on_going"},
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
	return enumData[int(t)-1].name
}

func (t TestEnum) Code() string {
	if int(t) < 1 || int(t) > len(enumData) {
		return ""
	}
	return enumData[int(t)-1].code
}

func (t TestEnum) MarshalJSON() ([]byte, error) {
	return MarshalJSON(t)
}

func (t *TestEnum) UnmarshalJSON(val []byte) error {
	idx, err := UnmarshalJSON(val, "test_enum")
	if err != nil {
		return err
	}
	*t = TestEnum(idx)
	return nil
}

func (t *TestEnum) Scan(val interface{}) error {
	idx, err := Scan(val, "test_enum")
	if err != nil {
		return err
	}
	*t = TestEnum(idx)
	return nil
}

func (t TestEnum) Value() (driver.Value, error) {
	return Value(t)
}

func TestRegister(t *testing.T) {

	t.Run("normal enum registration", func(t *testing.T) {
		enum := TestEnum(0)
		err := Register("test_enum", enumData, enum, &enum)
		assert.Nil(t, err)
	})

	t.Run("when register same key", func(t *testing.T) {
		enum := TestEnum(0)
		Register("test_enum", enumData, enum, &enum)
		err := Register("test_enum", enumData, enum, &enum)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrEnumKeyRegistered, err)
	})

	t.Run("when register non-integer enum", func(t *testing.T) {
		wrongEnum := "a"
		err := Register("wrong_enum_type", enumData, wrongEnum, &wrongEnum)
		assert.Equal(t, errors.ErrEnumNotInt, err)
	})

	t.Run("when register not-fully-implemented RegisterableEnum", func(t *testing.T) {
		wrongEnum := WrongEnum(0)
		err := Register("wrong_enum_type", enumData, wrongEnum, &wrongEnum)
		assert.Equal(t, errors.ErrEnumImplRegisterable, err)
	})

	t.Run("when wrong param", func(t *testing.T) {
		enum := TestEnum(0)
		err := Register("wrong_test_enum_param", enumData, enum, enum)
		assert.Equal(t, errors.ErrEnumNotIntPointer, err)
	})
	t.Run("when wrong param", func(t *testing.T) {
		enum := TestEnum(0)
		enum2 := WrongEnum(0)
		err := Register("wrong_test_enum_param", enumData, enum, &enum2)
		assert.Equal(t, errors.ErrEnumImplRegisterablePtr, err)
	})
}
