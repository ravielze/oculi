package radix36

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadix36_EncodeDecodeBytes(t *testing.T) {
	t.Run("Zero Integer", func(t *testing.T) {
		var data radix36
		data.Bytes([]byte{0, 0, 0, 0, 0, 0, 0, 0})
		result := data.ToInt()
		assert.Equal(t, int64(0), result)
		assert.Equal(t, bytes, data.lastType)
	})

	t.Run("Positive Integer", func(t *testing.T) {
		var data radix36
		data.Bytes([]byte{0, 0, 0, 0, 0, 0, 0, 100})
		result := data.ToInt()
		assert.Equal(t, int64(100), result)
		assert.Equal(t, bytes, data.lastType)
	})

	t.Run("Negative Integer", func(t *testing.T) {
		var data radix36
		data.Bytes([]byte{255, 255, 255, 255, 255, 255, 255, 6})
		result := data.ToInt()
		assert.Equal(t, int64(-250), result)
		assert.Equal(t, bytes, data.lastType)
	})
}
