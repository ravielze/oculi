package baseX

import (
	"github.com/gofrs/uuid"
)

type (
	BasicTransforming interface {
		IntTransformer
		BytesTransformer
		UUIDTransformer
		String() string
		Bytes() []byte
	}

	IntTransformer interface {
		FromInt(value int64)
		ToInt() int64
	}

	BytesTransformer interface {
		FromBytes(value []byte)
		ToBytes() []byte
	}

	UUIDTransformer interface {
		FromUUID(value uuid.UUID)
		FromUUIDString(value string) error
		ToUUID() uuid.UUID
		Randomize() UUIDTransformer
	}
)
