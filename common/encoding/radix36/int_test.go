package radix36

import (
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadix36_EncodeDecodeInt(t *testing.T) {

	t.Run("Zero Integer", func(t *testing.T) {
		var data Radix36
		data.Int(0)
		result := data.ToInt()
		assert.Equal(t, int64(0), result)
		assert.Equal(t, integer, data.lastType)
	})

	t.Run("Positive Integer", func(t *testing.T) {
		var data Radix36
		data.Int(100)
		result := data.ToInt()
		assert.Equal(t, int64(100), result)
		assert.Equal(t, integer, data.lastType)
	})

	t.Run("Negative Integer", func(t *testing.T) {
		var data Radix36
		data.Int(-250)
		result := data.ToInt()
		assert.Equal(t, int64(-250), result)
		assert.Equal(t, integer, data.lastType)
	})

	t.Run("Max Integer", func(t *testing.T) {
		var data Radix36
		data.Int(math.MaxInt64)
		result := data.ToInt()
		assert.Equal(t, int64(math.MaxInt64), result)
		assert.Equal(t, integer, data.lastType)
	})

	t.Run("Min Integer", func(t *testing.T) {
		var data Radix36
		data.Int(math.MinInt64)
		result := data.ToInt()
		assert.Equal(t, int64(math.MinInt64), result)
		assert.Equal(t, integer, data.lastType)
	})

	t.Run("Big Integer But Still In Range Int64", func(t *testing.T) {
		var data Radix36
		data.BigInt(*big.NewInt(18512))
		result := data.ToInt()
		assert.Equal(t, int64(18512), result)
		assert.Equal(t, biginteger, data.lastType)
	})

	t.Run("Big Integer But Still In Range Int64 (2)", func(t *testing.T) {
		var data Radix36
		data.BigInt(*big.NewInt(math.MaxInt64))
		result := data.ToInt()
		assert.Equal(t, int64(math.MaxInt64), result)
		assert.Equal(t, biginteger, data.lastType)
	})

	t.Run("Big Integer Max Int64", func(t *testing.T) {
		var data Radix36
		data.BigInt(*big.NewInt(0).

			// 3 * MaxInt64 + 1 * MaxInt64
			Add(
				// 3 * MaxInt64
				big.NewInt(0).Mul(
					big.NewInt(math.MaxInt64), big.NewInt(3),
				),

				// MaxInt64
				big.NewInt(math.MaxInt64),
			),
		)
		result := data.ToInt()
		assert.Equal(t, int64(0), result)
		assert.Equal(t, biginteger, data.lastType)
	})

	t.Run("Big Integer Min Int64", func(t *testing.T) {
		var data Radix36
		data.BigInt(*big.NewInt(0).Neg(big.NewInt(0).Mul(
			big.NewInt(3),
			big.NewInt(math.MaxInt64),
		)))
		result := data.ToInt()
		assert.Equal(t, int64(0), result)
		assert.Equal(t, biginteger, data.lastType)
	})
}
