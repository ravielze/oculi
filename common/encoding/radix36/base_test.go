package radix36

import (
	"math"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/martinlindhe/base36"
	"github.com/stretchr/testify/assert"
)

func Test_radix36_Bytes(t *testing.T) {
	t.Run("when empty", func(t *testing.T) {
		data := New()

		assert.Nil(t, data.Bytes())
	})
	t.Run("when not empty (1)", func(t *testing.T) {
		data := NewFromInt(1527392)

		assert.NotNil(t, data.Bytes())
		assert.Len(t, data.Bytes(), 8)
	})
	t.Run("when not empty (2)", func(t *testing.T) {
		data := NewFromUUID(uuid.Must(uuid.NewV4()))

		assert.NotNil(t, data.Bytes())
		assert.Len(t, data.Bytes(), 16)
	})
}

func TestNewRadix36(t *testing.T) {
	t.Run("when success", func(t *testing.T) {
		data, err := NewRadix36("ABCDE123")

		assert.NotNil(t, data)
		assert.Nil(t, err)
	})
	t.Run("when error", func(t *testing.T) {
		data, err := NewRadix36("ABCDE123!!!")

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Nil(t, data)
	})
}

func TestRadix36(t *testing.T) {
	t.Run("when success", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Radix36("ABCDE123")
		})
	})
	t.Run("when error", func(t *testing.T) {
		assert.Panics(t, func() {
			Radix36("ABCDE123!!!")
		})
	})
}

func TestNew(t *testing.T) {
	t.Run("any condition", func(t *testing.T) {
		data := New()
		assert.Nil(t, data.Bytes())
		assert.Equal(t, none, data.(*radix36).lastType)
	})
}

func TestNewFromInt(t *testing.T) {
	t.Run("any condition", func(t *testing.T) {
		data := NewFromInt(28263)
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, integer, data.(*radix36).lastType)
		assert.Len(t, data.Bytes(), 8)
	})
}

func TestNewFromUUID(t *testing.T) {
	t.Run("any condition", func(t *testing.T) {
		data := NewFromUUID(uuid.Must(uuid.NewV4()))
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, t_uuid, data.(*radix36).lastType)
		assert.Len(t, data.Bytes(), 16)
	})
}

func TestNewFromBytes(t *testing.T) {
	t.Run("any condition (1)", func(t *testing.T) {
		data := NewFromBytes([]byte{1, 2, 3})
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, bytes, data.(*radix36).lastType)
		assert.Len(t, data.Bytes(), 3)
	})
	t.Run("any condition (2)", func(t *testing.T) {
		data := NewFromBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, bytes, data.(*radix36).lastType)
		assert.Len(t, data.Bytes(), 10)
	})
	t.Run("any condition (3)", func(t *testing.T) {
		data := NewFromBytes([]byte{})
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, bytes, data.(*radix36).lastType)
		assert.Len(t, data.Bytes(), 0)
	})
}

func TestNewFromUUIDString(t *testing.T) {
	t.Run("when success", func(t *testing.T) {
		data, err := NewFromUUIDString(uuid.Must(uuid.NewV4()).String())
		assert.NotNil(t, data.Bytes())
		assert.Nil(t, err)
		assert.Equal(t, t_uuid, data.(*radix36).lastType)
		assert.Len(t, data.Bytes(), 16)
	})
	t.Run("when error", func(t *testing.T) {
		data, err := NewFromUUIDString("abc!!!")
		assert.Nil(t, data.Bytes())
		assert.Error(t, err)
		assert.NotNil(t, err)
		assert.Equal(t, none, data.(*radix36).lastType)
	})
}

func TestNewRandomize(t *testing.T) {
	t.Run("any condition", func(t *testing.T) {
		data := NewRandomize()
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, t_uuid, data.(*radix36).lastType)
		assert.Len(t, data.Bytes(), 16)
	})
}

func Test_radix36_String(t *testing.T) {
	t.Run("when not empty (1)", func(t *testing.T) {
		data := NewFromInt(623698779)
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, "ABC123", data.String())
	})

	t.Run("when not empty (2)", func(t *testing.T) {
		data := NewFromInt(math.MaxInt64)
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, "1Y2P0IJ32E8E7", data.String())
	})

	t.Run("when not empty (3)", func(t *testing.T) {
		data := NewFromInt(128)
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, "3K", data.String())
	})

	t.Run("when not empty (4)", func(t *testing.T) {
		data := Radix36("3K")
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, "3K", data.String())
		assert.Equal(t, int64(128), data.ToInt())
	})

	t.Run("when not empty (5)", func(t *testing.T) {
		data := Radix36("3W5E11264SGPO")
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, "3W5E11264SGPO", data.String())
		assert.Equal(t, int64(-100), data.ToInt())
	})
	t.Run("when not empty (6)", func(t *testing.T) {
		uuid := uuid.Must(uuid.NewV4())
		data := NewFromUUID(uuid)
		assert.NotNil(t, data.Bytes())
		assert.Equal(t, uuid.Bytes(), data.Bytes())
		assert.Equal(t, base36.EncodeBytes(uuid.Bytes()), data.String())
	})

	t.Run("when empty", func(t *testing.T) {
		var data radix36
		assert.Nil(t, data.Bytes())
		assert.Equal(t, "", data.String())
	})
}
