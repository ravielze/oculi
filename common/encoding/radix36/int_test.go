package radix36

import (
	"math"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func createUUID(uuidString string) uuid.UUID {
	return uuid.FromStringOrNil(uuidString)
}

func TestRadix36_EncodeDecodeInt(t *testing.T) {

	t.Run("Zero Integer", func(t *testing.T) {
		var data radix36
		data.FromInt(0)
		assert.Equal(t, integer, data.lastType)
		assert.Equal(t, int64(0), data.ToInt())
		assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-0000-000000000000"), data.ToUUID())
		assert.Equal(t, "0", data.String())
	})

	t.Run("Positive Integer", func(t *testing.T) {
		var data radix36
		data.FromInt(100)
		assert.Equal(t, integer, data.lastType)
		assert.Equal(t, int64(100), data.ToInt())
		assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 100}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-0000-000000000064"), data.ToUUID())
		assert.Equal(t, "2S", data.String())
	})

	t.Run("Negative Integer", func(t *testing.T) {
		var data radix36
		data.FromInt(-100)
		assert.Equal(t, integer, data.lastType)
		assert.Equal(t, int64(-100), data.ToInt())
		assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 255, 156}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-ffff-ffffffffff9c"), data.ToUUID())
		assert.Equal(t, "3W5E11264SGPO", data.String())
	})

	t.Run("Max Integer", func(t *testing.T) {
		var data radix36
		data.FromInt(math.MaxInt64)
		assert.Equal(t, integer, data.lastType)
		assert.Equal(t, int64(math.MaxInt64), data.ToInt())
		assert.Equal(t, []byte{127, 255, 255, 255, 255, 255, 255, 255}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-7fff-ffffffffffff"), data.ToUUID())
		assert.Equal(t, "1Y2P0IJ32E8E7", data.String())
	})

	t.Run("Min Integer", func(t *testing.T) {
		var data radix36
		data.FromInt(math.MinInt64)
		assert.Equal(t, integer, data.lastType)
		assert.Equal(t, int64(math.MinInt64), data.ToInt())
		assert.Equal(t, []byte{128, 0, 0, 0, 0, 0, 0, 0}, data.ToBytes())
		assert.Equal(t, createUUID("00000000-0000-0000-8000-000000000000"), data.ToUUID())
		assert.Equal(t, "1Y2P0IJ32E8E8", data.String())
	})

	// t.Run("Big Integer But Still In Range Int64", func(t *testing.T) {
	// 	var data radix36
	// 	data.BigInt(*big.NewInt(18512))
	// 	result := data.ToInt()
	// 	assert.Equal(t, int64(18512), result)
	// 	assert.Equal(t, biginteger, data.lastType)
	// })

	// t.Run("Big Integer But Still In Range Int64 (2)", func(t *testing.T) {
	// 	var data radix36
	// 	data.BigInt(*big.NewInt(math.MaxInt64))
	// 	result := data.ToInt()
	// 	assert.Equal(t, int64(math.MaxInt64), result)
	// 	assert.Equal(t, biginteger, data.lastType)
	// })

	// t.Run("Big Integer Max Int64", func(t *testing.T) {
	// 	var data radix36
	// 	data.BigInt(*big.NewInt(0).

	// 		// 3 * MaxInt64 + 1 * MaxInt64
	// 		Add(
	// 			// 3 * MaxInt64
	// 			big.NewInt(0).Mul(
	// 				big.NewInt(math.MaxInt64), big.NewInt(3),
	// 			),

	// 			// MaxInt64
	// 			big.NewInt(math.MaxInt64),
	// 		),
	// 	)
	// 	result := data.ToInt()
	// 	assert.Equal(t, int64(0), result)
	// 	assert.Equal(t, biginteger, data.lastType)
	// })

	// t.Run("Big Integer Min Int64", func(t *testing.T) {
	// 	var data radix36
	// 	data.BigInt(*big.NewInt(0).Neg(big.NewInt(0).Mul(
	// 		big.NewInt(3),
	// 		big.NewInt(math.MaxInt64),
	// 	)))
	// 	result := data.ToInt()
	// 	assert.Equal(t, int64(0), result)
	// 	assert.Equal(t, biginteger, data.lastType)
	// })
}
