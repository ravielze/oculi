package radix36

import (
	"math/big"
)

func (r *Radix36) BigInt(value big.Int) {
	r.data = value.Bytes()
	r.lastType = biginteger
}

func (r *Radix36) ToBigInt() big.Int {
	var result big.Int
	switch r.lastType {
	case integer:
		result.SetInt64(r.ToInt())
	case biginteger:
		result.SetBytes(r.data)
	}
	return result
}
