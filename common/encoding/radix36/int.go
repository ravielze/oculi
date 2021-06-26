package radix36

import (
	"encoding/binary"
	"math"
	"math/big"
)

func (r *Radix36) Int(value int64) {
	r.data = make([]byte, 8)
	binary.BigEndian.PutUint64(r.data, uint64(value))
	r.lastType = integer
}

func (r *Radix36) ToInt() int64 {
	switch r.lastType {
	case integer:
		return int64(binary.BigEndian.Uint64(r.data))
	case biginteger:
		var i big.Int
		i.SetBytes(r.data)
		if i.CmpAbs(big.NewInt(math.MaxInt64)) <= 0 {
			return i.Int64()
		}
	}
	return int64(0)
}
