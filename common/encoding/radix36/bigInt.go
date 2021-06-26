package radix36

import (
	"math/big"
)

func (r *Radix36) BigInt(value big.Int) {
	r.data = value.Bytes()
	var signByte byte
	switch value.Sign() {
	case -1:
		signByte = 0
	default:
		signByte = 1
	}
	r.data = append(r.data, signByte)
	r.lastType = biginteger
}

func (r *Radix36) ToBigInt() big.Int {
	var result big.Int
	switch r.lastType {
	case integer:
		result.SetInt64(r.ToInt())
	case biginteger:
		x := len(r.data) - 1
		result.SetBytes(r.data[:x])
		if r.data[x] == byte(0) {
			result.Mul(&result, big.NewInt(-1))
		}
	}
	return result
}
