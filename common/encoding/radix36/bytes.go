package radix36

func (r *radix36) FromBytes(value []byte) {
	r.data = value
	r.lastType = bytes
}
func (r *radix36) ToBytes() []byte {
	return r.data
}
