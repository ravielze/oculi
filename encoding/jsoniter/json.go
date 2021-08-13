package jsoniter

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/ravielze/oculi/encoding"
)

type (
	impl struct {
		json jsoniter.API
	}
)

var instance encoding.Encoding = &impl{json: jsoniter.ConfigCompatibleWithStandardLibrary}

func (i *impl) Marshal(val interface{}) ([]byte, error) {
	return i.json.Marshal(val)
}

func (i *impl) Unmarshal(data []byte, val interface{}) error {
	return i.json.Unmarshal(data, val)
}

func New() encoding.Encoding {
	return instance
}
