package bcrypt

import (
	"errors"

	"github.com/ravielze/oculi/hash"
	bcryptLib "golang.org/x/crypto/bcrypt"
)

type (
	bcrypt struct {
		cost int
	}
)

var (
	ErrBcryptInvalidCost = errors.New("bcrypt invalid cost")
	ErrPasswordMismatch  = errors.New("password mismatch")
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
		return nil, ErrBcryptInvalidCost
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

func (b *bcrypt) Verify(raw string, hashed string) error {
	buffHash := []byte(hashed)
	buffRaw := []byte(raw)

	err := bcryptLib.CompareHashAndPassword(buffHash, buffRaw)
	if err != nil {
		//Todo convert error
		return ErrPasswordMismatch
	}

	return nil
}
