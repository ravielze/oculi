package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/ravielze/oculi/common/model/dto/user"
	"github.com/ravielze/oculi/token"
)

type encImpl struct {
	key        []byte
	alg        string
	identifier string
}

func NewEncoder(key string, alg string, identifier string) token.Encoder {
	return &encImpl{
		key:        []byte(key),
		alg:        alg,
		identifier: identifier,
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

func (e *encImpl) CreateAccessClaims(credentials user.CredentialsDTO, exp time.Duration) token.Claims {
	return &accessClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp).Unix(),
		},
		Crd:        credentials,
		Identifier: e.identifier,
	}
}

func (e *encImpl) CreateAccessAndEncode(credentials user.CredentialsDTO, exp time.Duration) (string, error) {
	tokenClaims := e.CreateAccessClaims(credentials, exp)
	return e.Encode(tokenClaims)
}

func (e *encImpl) CreateRefreshClaims(userId uint64, exp time.Duration) (token.Claims, uuid.UUID) {
	uuidToken := uuid.Must(uuid.NewV4())
	return &refreshClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp).Unix(),
		},
		UserID:     userId,
		Identifier: e.identifier,
		Token:      uuidToken.String(),
	}, uuidToken
}

func (e *encImpl) CreateRefreshAndEncode(userId uint64, exp time.Duration) (string, uuid.UUID, error) {
	tokenClaims, uuidToken := e.CreateRefreshClaims(userId, exp)
	jwt, err := e.Encode(tokenClaims)
	if err != nil {
		return "", uuid.UUID{}, err
	}
	return jwt, uuidToken, nil
}
