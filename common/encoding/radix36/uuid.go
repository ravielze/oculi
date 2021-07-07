package radix36

import (
	"github.com/gofrs/uuid"
)


func (r *radix36) UUID(value uuid.UUID) {
	r.data = value.Bytes()
	r.lastType = t_uuid
}

func (r *radix36) ToUUID() uuid.UUID {
	data := r.data
	if r.lastType == biginteger {
		x := r.ToBigInt()
		data = x.Bytes()
	}
	b := len(data)
	if b >= 16 {
		return uuid.Must(uuid.FromBytes(data[:16]))
	} else {
		zero := make([]byte, 16-b)
		zero = append(zero, data...)
		return uuid.Must(uuid.FromBytes(zero))
	}
}

func (r *radix36) Randomize() {
	r.data = uuid.Must(uuid.NewV4()).Bytes()
	r.lastType = t_uuid
}