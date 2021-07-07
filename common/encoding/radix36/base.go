package radix36

import (
	"github.com/ravielze/oculi/common/encoding"
)

type (
	radix36 struct {
		data     []byte
		lastType R36Type
	}

	R36Type int
)

const (
	none R36Type = iota + 1
	integer
	biginteger
	bytes
	t_uuid
	r36
)

func New() encoding.BasicEncoding {
	return &radix36{
		data:     nil,
		lastType: none,
	}
}

func NewRandomize() encoding.BasicEncoding {
	x := New()
	x.Randomize()
	return x
}
