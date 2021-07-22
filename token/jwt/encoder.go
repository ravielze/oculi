package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ravielze/oculi/common/model/dto/user"
	"github.com/ravielze/oculi/token"
)

type encImpl struct {
	key []byte
	alg string
}

func NewEncoder(key string, alg string) token.Encoder {
	return &encImpl{
		key: []byte(key),
		alg: alg,
	}
}

func (e *encImpl) Encode(claims token.Claims) (string, error) {
	newToken := jwt.New(jwt.GetSigningMethod(e.alg))
	newToken.Claims = claims

	signedToken, err := newToken.SignedString(e.key)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (e *encImpl) CreateClaims(credentials user.CredentialsDTO, exp int64) token.Claims {
	return &claims{
		&jwt.StandardClaims{
			ExpiresAt: exp,
		},
		credentials,
	}
}
