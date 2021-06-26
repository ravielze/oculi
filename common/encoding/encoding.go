package encoding

import (
	"math/big"

	"github.com/gofrs/uuid"
)

type (
	BasicEncoding interface {
		IntEncodeDecode
		BigIntEncodeDecode
		//BytesEncodeDecode
		//UUIDEncodeDecode
		//Randomize()
	}

	IntEncodeDecode interface {
		Int(value int64)
		ToInt() int64
	}

	BigIntEncodeDecode interface {
		BigInt(value big.Int)
		ToBigInt() big.Int
	}

	BytesEncodeDecode interface {
		EncodeBytes(value []byte)
		DecodeBytes() ([]byte, error)
	}

	UUIDEncodeDecode interface {
		EncodeUUID(value uuid.UUID)
		DecodeUUID() (uuid.UUID, error)
	}
)
