package radix36

import (
	"github.com/ravielze/oculi/common/encoding"
)

type (
	Radix36 struct {
		data     []byte
		lastType R36Type
	}

	R36Type int
)

const (
	none R36Type = iota + 1
	integer
	biginteger
	uuid
	r36
)

func New() encoding.BasicEncoding {
	return &Radix36{
		data:     nil,
		lastType: none,
	}
}
