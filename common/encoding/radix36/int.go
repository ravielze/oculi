package radix36

import (
	"encoding/binary"
	"math"
	"math/big"
)

func (r *radix36) Int(value int64) {
	r.data = make([]byte, 8)
	binary.BigEndian.PutUint64(r.data, uint64(value))
	r.lastType = integer
}

func (r *radix36) ToInt() int64 {
	switch r.lastType {
	case bytes:
		b := len(r.data)
		if b >= 8 {
			return int64(binary.BigEndian.Uint64(r.data[:8]))
		} else {
			zero := make([]byte, 8-b)
			zero = append(zero, r.data...)
			return int64(binary.BigEndian.Uint64(zero))
		}
	case integer:
		return int64(binary.BigEndian.Uint64(r.data))
	case biginteger:
		var i big.Int
		x := len(r.data) - 1
		i.SetBytes(r.data[:x])
		if r.data[x] == byte(0) {
			if i.CmpAbs(big.NewInt(math.MinInt64)) <= 0 {
				return i.Int64() * -1
			}
		} else {
			if i.CmpAbs(big.NewInt(math.MaxInt64)) <= 0 {
				return i.Int64()
			}
		}
	case t_uuid:
		j := r.ToBigInt()
		return j.Int64()
	}
	return int64(0)
}
