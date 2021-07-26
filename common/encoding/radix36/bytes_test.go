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
		assert.Equal(t, "0", data.String())
	})

	t.Run("Positive Integer", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{0, 0, 0, 0, 0, 0, 0, 100})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(100), data.ToInt())
		assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 100}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-0000-000000000064"), data.ToUUID())
		assert.Equal(t, "2S", data.String())
	})

	t.Run("Negative Integer", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{255, 255, 255, 255, 255, 255, 255, 156})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(-100), data.ToInt())
		assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255, 156}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-ffff-ffffffffff9c"), data.ToUUID())
		assert.Equal(t, "3W5E11264SGPO", data.String())
		var i uint64
		j := int64(-100)
		i = uint64(j)
		assert.Equal(t, uint64(0xffffffffffffff9c), i)
	})

	t.Run("Any Bytes (1)", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{128})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(128), data.ToInt())
		assert.Equal(t, []byte{128}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-0000-000000000080"), data.ToUUID())
		assert.Equal(t, "3K", data.String())
	})
	t.Run("Any Bytes (2)", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{11, 11, 11, 11, 11, 11, 128})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(3108366801636224), data.ToInt())
		assert.Equal(t, []byte{11, 11, 11, 11, 11, 11, 128}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-000b-0b0b0b0b0b80"), data.ToUUID())
		assert.Equal(t, "ULTNZU1I0W", data.String())
	})
	t.Run("Any Bytes (3)", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{12, 13, 14, 15, 11, 11, 11, 11, 11, 11, 128})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(1083972277370555264), data.ToInt())
		assert.Equal(t, []byte{12, 13, 14, 15, 11, 11, 11, 11, 11, 11, 128}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-000c-0d0e-0f0b-0b0b0b0b0b80"), data.ToUUID())
		assert.Equal(t, "1TWEC005IVI4029A8", data.String())
	})
	t.Run("Any Bytes (4)", func(t *testing.T) {
		var data radix36
		data.FromBytes([]byte{176, 202, 111, 156, 203, 132, 12, 13, 14, 15, 11, 11, 11, 11, 11, 11, 128})

		assert.Equal(t, bytes, data.lastType)
		assert.Equal(t, int64(1083972277370555264), data.ToInt())
		assert.Equal(t, []byte{176, 202, 111, 156, 203, 132, 12, 13, 14, 15, 11, 11, 11, 11, 11, 11, 128}, data.ToBytes())
		assert.Equal(t, createUUID("b0ca6f9c-cb84-0c0d-0e0f-0b0b0b0b0b0b"), data.ToUUID())
		assert.Equal(t, "22FEPZS758WJ4FDSMAZ313C3VR4", data.String())
	})
}
