package bcrypt

import (
	"github.com/ravielze/oculi/constant/errors"
	"github.com/ravielze/oculi/hash"
	bcryptLib "golang.org/x/crypto/bcrypt"
)

type (
	bcrypt struct {
		cost int
	}
)

const (
	MinCost     = bcryptLib.MinCost
	MaxCost     = bcryptLib.MaxCost
	DefaultCost = bcryptLib.DefaultCost
)

func NewHash() (hash.Hash, error) {
	return NewHashWithCost(DefaultCost)
}

func NewHashWithCost(cost int) (hash.Hash, error) {
	if cost < MinCost {
		cost = DefaultCost
	}

	if cost > MaxCost {
		return nil, errors.ErrBcryptInvalidCost
	}

	return &bcrypt{cost: cost}, nil
}

func (b *bcrypt) Hash(raw string) (string, error) {
	buff := []byte(raw)

	hash, err := bcryptLib.GenerateFromPassword(buff, b.cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (b *bcrypt) Verify(raw string, hashed string) (bool, error) {
	buffHash := []byte(hashed)
	buffRaw := []byte(raw)

	err := bcryptLib.CompareHashAndPassword(buffHash, buffRaw)
	if err != nil {
		//Todo convert error
		return false, errors.ErrPasswordMismatch
	}

	return true, nil
}
