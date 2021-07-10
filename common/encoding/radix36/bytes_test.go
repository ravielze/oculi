package radix36

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadix36_EncodeDecodeBytes(t *testing.T) {
	t.Run("Zero Integer", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{0, 0, 0, 0, 0, 0, 0, 0})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(0), data.ToInt())
		assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-0000-000000000000"), data.ToUUID())
	})

	t.Run("Positive Integer", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{0, 0, 0, 0, 0, 0, 0, 100})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(100), data.ToInt())
		assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 100}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-0000-000000000064"), data.ToUUID())
	})

	t.Run("Negative Integer", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{255, 255, 255, 255, 255, 255, 255, 156})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(-100), data.ToInt())
		assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255, 156}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-ffff-ffffffffff9c"), data.ToUUID())
	})

	t.Run("Any Bytes (1)", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{128})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(128), data.ToInt())
		assert.Equal(t, []byte{128}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-0000-000000000080"), data.ToUUID())
	})
	t.Run("Any Bytes (2)", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{11, 11, 11, 11, 11, 11, 128})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(3108366801636224), data.ToInt())
		assert.Equal(t, []byte{11, 11, 11, 11, 11, 11, 128}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-000b-0b0b0b0b0b80"), data.ToUUID())
	})
	t.Run("Any Bytes (3)", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{12, 13, 14, 15, 11, 11, 11, 11, 11, 11, 128})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(1083972277370555264), data.ToInt())
		assert.Equal(t, []byte{12, 13, 14, 15, 11, 11, 11, 11, 11, 11, 128}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-000c-0d0e-0f0b-0b0b0b0b0b80"), data.ToUUID())
	})
	t.Run("Any Bytes (4)", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{176, 202, 111, 156, 203, 132, 12, 13, 14, 15, 11, 11, 11, 11, 11, 11, 128})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(1083972277370555264), data.ToInt())
		assert.Equal(t, []byte{176, 202, 111, 156, 203, 132, 12, 13, 14, 15, 11, 11, 11, 11, 11, 11, 128}, data.ToBytes())
		assert.Equal(t, createUUID("b0ca6f9c-cb84-0c0d-0e0f-0b0b0b0b0b0b"), data.ToUUID())
	})
}
