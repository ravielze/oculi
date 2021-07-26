package radix36

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadix36_EncodeDecodeUUID(t *testing.T) {
	t.Run("Zero Integer", func(t *testing.T) {
		var data radix36
		data.FromUUIDString("00000000-0000-0000-0000-000000000000")
		assert.Equal(t, t_uuid, data.lastType)
		assert.Equal(t, int64(0), data.ToInt())
		assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-0000-000000000000"), data.ToUUID())
		assert.Equal(t, "0", data.String())
	})

	t.Run("Positive Integer", func(t *testing.T) {
		var data radix36
		data.FromUUIDString("00000000-0000-0000-0000-000000000064")
		assert.Equal(t, t_uuid, data.lastType)
		assert.Equal(t, int64(100), data.ToInt())
		assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-0000-000000000064"), data.ToUUID())
		assert.Equal(t, "2S", data.String())
	})

	t.Run("Negative Integer", func(t *testing.T) {
		var data radix36
		data.FromUUIDString("00000000-0000-0000-ffff-ffffffffff9c")
		assert.Equal(t, t_uuid, data.lastType)
		assert.Equal(t, int64(-100), data.ToInt())
		assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 255, 156}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-ffff-ffffffffff9c"), data.ToUUID())
		assert.Equal(t, "3W5E11264SGPO", data.String())
	})
}
