package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ravielze/oculi/common/model/dto/user"
	"github.com/ravielze/oculi/token"
)

type encImpl struct {
	key      []byte
	alg      string
	serverId string
}

func NewEncoder(key string, alg string, serverId string) token.Encoder {
	return &encImpl{
		key:      []byte(key),
		alg:      alg,
		serverId: serverId,
	}
}

func (e *encImpl) Encode(claims token.Claims) (string, error) {
	newToken := jwt.NewWithClaims(jwt.GetSigningMethod(e.alg), claims)

	signedToken, err := newToken.SignedString(e.key)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (e *encImpl) CreateClaims(credentials user.CredentialsDTO, exp time.Duration) token.Claims {
	return &claims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp).Unix(),
		},
		credentials,
		e.serverId,
	}
}

func (e *encImpl) CreateAndEncode(credentials user.CredentialsDTO, exp time.Duration) (string, error) {
	tokenClaims := e.CreateClaims(credentials, exp)
	return e.Encode(tokenClaims)
}
