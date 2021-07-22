package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/ravielze/oculi/token"
)

var (
	ErrUnclaimedToken = errors.New("unclaimed token")
)

type (
	decImpl struct {
		key []byte
	}
)

func NewDecoder(key string) token.Decoder {
	return &decImpl{
		key: []byte(key),
	}
}

func (d *decImpl) Decode(token string) (token.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token, &claims{},
		func(t *jwt.Token) (interface{}, error) {
			return d.key, nil
		},
	)

	if err != nil || tokenClaims == nil || tokenClaims.Claims == nil {
		return nil, ErrUnclaimedToken
	}

	c := tokenClaims.Claims.(*claims)
	return c, nil
}
