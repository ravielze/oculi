package jwt

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	consts "github.com/ravielze/oculi/constant/errors"
	"github.com/ravielze/oculi/token"
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
	if err != nil {
		return nil, err
	}

	if tokenClaims == nil || tokenClaims.Claims == nil {
		return nil, consts.ErrUnclaimedToken
	}

	c := tokenClaims.Claims.(*claims)
	return c, nil
}

func extractToken(req *http.Request) string {
	bearToken := req.Header.Get("Authorization")
	if len(bearToken) == 0 {
		bearTokenQuery, ok := req.URL.Query()["Authorization"]
		if ok && len(bearTokenQuery) == 1 {
			return bearTokenQuery[0]
		}
		return ""
	}
	keys := strings.Split(bearToken, " ")
	if len(keys) == 2 {
		return keys[1]
	}
	return ""
}

func (d *decImpl) DecodeHttpRequest(req *http.Request) (token.Claims, error) {
	if ex := extractToken(req); ex != "" && len(ex) > 0 {
		return d.Decode(ex)
	}
	return nil, consts.ErrNoBearerToken
}
