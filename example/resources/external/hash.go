package external

import (
	"github.com/ravielze/oculi/hash"
	"github.com/ravielze/oculi/hash/bcrypt"
)

func NewHash() (hash.Hash, error) {
	return bcrypt.NewHash()
}
