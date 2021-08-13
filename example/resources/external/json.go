package external

import (
	"github.com/ravielze/oculi/encoding"
	"github.com/ravielze/oculi/encoding/jsoniter"
)

func NewJsonEncoding() encoding.Encoding {
	return jsoniter.New()
}
