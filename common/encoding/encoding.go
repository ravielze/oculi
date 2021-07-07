package encoding

import (
	"math/big"

	"github.com/gofrs/uuid"
)

type (
	BasicEncoding interface {
		IntEncodeDecode
		BigIntEncodeDecode
		BytesEncodeDecode
		UUIDEncodeDecode
	}

	IntEncodeDecode interface {
		Int(value int64)
		ToInt() int64
	}

	BigIntEncodeDecode interface {
		BigInt(value big.Int)
		BigIntFromInt64(value int64)
		ToBigInt() big.Int
	}

	BytesEncodeDecode interface {
		Bytes(value []byte)
		ToBytes() []byte
	}

	UUIDEncodeDecode interface {
		UUID(value uuid.UUID)
		ToUUID() uuid.UUID
		Randomize()
	}
)
