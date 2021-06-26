package radix36

import (
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadix36_EncodeDecodeBigInt(t *testing.T) {
	t.Run("Zero Integer to Big Int and Viceversa", func(t *testing.T) {
		var data Radix36
		data.Int(0)
		result := data.ToBigInt()

		var data2 Radix36
		data2.BigInt(*big.NewInt(0))
		result2 := data2.ToInt()

		assert.Equal(t, *big.NewInt(0), result)
		assert.Equal(t, integer, data.lastType)

		assert.Equal(t, int64(0), result2)
		assert.Equal(t, biginteger, data2.lastType)
	})

	t.Run("Positive Integer to Big Int and Viceversa", func(t *testing.T) {
		var data Radix36
		data.Int(130451)
		result := data.ToBigInt()

		var data2 Radix36
		data2.BigInt(*big.NewInt(130451))
		result2 := data2.ToInt()

		assert.Equal(t, *big.NewInt(130451), result)
		assert.Equal(t, integer, data.lastType)

		assert.Equal(t, int64(130451), result2)
		assert.Equal(t, biginteger, data2.lastType)
	})

	t.Run("Negative Integer to Big Int and Viceversa", func(t *testing.T) {
		var data Radix36
		data.Int(-12345678)
		result := data.ToBigInt()

		var data2 Radix36
		data2.BigInt(*big.NewInt(-12345678))
		result2 := data2.ToInt()

		assert.Equal(t, *big.NewInt(-12345678), result)
		assert.Equal(t, integer, data.lastType)

		assert.Equal(t, int64(-12345678), result2)
		assert.Equal(t, biginteger, data2.lastType)
	})

	t.Run("Min Integer to Big Int and Viceversa", func(t *testing.T) {
		var data Radix36
		data.Int(math.MinInt64)
		result := data.ToBigInt()

		var data2 Radix36
		data2.BigInt(*big.NewInt(math.MinInt64))
		result2 := data2.ToInt()

		assert.Equal(t, *big.NewInt(math.MinInt64), result)
		assert.Equal(t, integer, data.lastType)

		assert.Equal(t, int64(math.MinInt64), result2)
		assert.Equal(t, biginteger, data2.lastType)
	})

	t.Run("Max Integer to Big Int and Viceversa", func(t *testing.T) {
		var data Radix36
		data.Int(math.MaxInt64)
		result := data.ToBigInt()

		var data2 Radix36
		data2.BigInt(*big.NewInt(math.MaxInt64))
		result2 := data2.ToInt()

		assert.Equal(t, *big.NewInt(math.MaxInt64), result)
		assert.Equal(t, integer, data.lastType)

		assert.Equal(t, int64(math.MaxInt64), result2)
		assert.Equal(t, biginteger, data2.lastType)
	})
}
