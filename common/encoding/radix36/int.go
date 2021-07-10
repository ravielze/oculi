package radix36

import (
	"encoding/binary"
)

func (r *radix36) FromInt(value int64) {
	r.data = make([]byte, 8)
	binary.BigEndian.PutUint64(r.data, uint64(value))
	r.lastType = integer
}

func (r *radix36) ToInt() int64 {
	switch r.lastType {
	case bytes, t_uuid:
		b := len(r.data)
		if b >= 8 {
			return int64(binary.BigEndian.Uint64(r.data[len(r.data)-8:]))
		} else {
			zero := make([]byte, 8-b)
			zero = append(zero, r.data...)
			return int64(binary.BigEndian.Uint64(zero))
		}
	case integer:
		return int64(binary.BigEndian.Uint64(r.data))
	}
	return int64(0)
}
