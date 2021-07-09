package encoding

import (
	"github.com/gofrs/uuid"
)

type (
	BasicEncoding interface {
		IntEncodeDecode
		BytesEncodeDecode
		UUIDEncodeDecode
		String() string
	}

	IntEncodeDecode interface {
		Int(value int64)
		ToInt() int64
	}

	BytesEncodeDecode interface {
		Bytes(value []byte)
		ToBytes() []byte
	}

	UUIDEncodeDecode interface {
		UUID(value uuid.UUID)
		UUIDString(value string) error
		ToUUID() uuid.UUID
		Randomize() UUIDEncodeDecode
	}
)
