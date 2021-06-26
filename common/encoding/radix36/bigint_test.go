package radix36

import (
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
}
