package jwt

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	consts "github.com/ravielze/oculi/constant/errors"
	key "github.com/ravielze/oculi/constant/key"
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

func (d *decImpl) decode(jwtStr string, obj jwt.Claims) (*jwt.Token, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		jwtStr, obj,
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
	return tokenClaims, nil
}

func (d *decImpl) DecodeAccess(jwtStr string) (token.Claims, error) {
	tokenClaims, err := d.decode(jwtStr, &accessClaims{})
	if err != nil {
		return nil, err
	}

	c := tokenClaims.Claims.(*accessClaims)
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

func (d *decImpl) DecodeAccessHeader(req *http.Request) (token.Claims, error) {
	if ex := extractToken(req); ex != "" && len(ex) > 0 {
		return d.DecodeAccess(ex)
	}
	return nil, consts.ErrNoBearerToken
}

func (d *decImpl) DecodeAccessCookie(req *http.Request) (token.Claims, error) {
	cookie, err := req.Cookie(key.KeyAccessToken)
	if err != nil || cookie == nil || cookie.Value == "" {
		return nil, consts.ErrCookieNotFound
	}
	return d.DecodeAccess(cookie.Value)
}

func (d *decImpl) DecodeRefresh(jwtStr string) (token.Claims, error) {
	tokenClaims, err := d.decode(jwtStr, &refreshClaims{})
	if err != nil {
		return nil, err
	}
	c := tokenClaims.Claims.(*refreshClaims)
	return c, nil
}

func (d *decImpl) DecodeRefreshCookie(req *http.Request) (token.Claims, error) {
	cookie, err := req.Cookie(key.KeyRefreshToken)
	if err != nil || cookie == nil || cookie.Value == "" {
		return nil, consts.ErrCookieNotFound
	}
	return d.DecodeRefresh(cookie.Value)
}
