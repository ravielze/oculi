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
		Bytes() []byte
	}

	IntEncodeDecode interface {
		FromInt(value int64)
		ToInt() int64
	}

	BytesEncodeDecode interface {
		FromBytes(value []byte)
		ToBytes() []byte
	}

	UUIDEncodeDecode interface {
		FromUUID(value uuid.UUID)
		FromUUIDString(value string) error
		ToUUID() uuid.UUID
		Randomize() UUIDEncodeDecode
	}
)
