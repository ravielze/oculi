package radix36

import (
	"github.com/gofrs/uuid"
	"github.com/ravielze/oculi/common/encoding"
)

func (r *radix36) FromUUID(value uuid.UUID) {
	r.data = value.Bytes()
	r.lastType = t_uuid
}

func (r *radix36) FromUUIDString(value string) error {
	val, err := uuid.FromString(value)
	if err != nil {
		return err
	}
	r.FromUUID(val)
	return nil
}

func (r *radix36) ToUUID() uuid.UUID {
	data := r.data
	b := len(data)
	if b >= 16 {
		return uuid.Must(uuid.FromBytes(data[:16]))
	} else {
		zero := make([]byte, 16-b)
		zero = append(zero, data...)
		return uuid.Must(uuid.FromBytes(zero))
	}
}

func (r *radix36) Randomize() encoding.UUIDEncodeDecode {
	r.data = uuid.Must(uuid.NewV4()).Bytes()
	r.lastType = t_uuid
	return r
}
