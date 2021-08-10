package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ravielze/oculi/common/model/dto/user"
	consts "github.com/ravielze/oculi/constant/errors"
	"github.com/ravielze/oculi/constant/oculiTime"
)

type (
	accessClaims struct {
		*jwt.StandardClaims
		Crd        user.CredentialsDTO `json:"credentials"`
		Identifier string              `json:"identifier"`
	}

	refreshClaims struct {
		*jwt.StandardClaims
		UserID     uint64 `json:"user_id"`
		Identifier string `json:"identifier"`
		Token      string `json:"token"`
	}
)

func (c *refreshClaims) Credentials() user.CredentialsDTO {
	return user.CredentialsDTO{ID: c.UserID, Metadata: c.Token}
}

func (c *refreshClaims) Valid() error {
	jwt.TimeFunc = oculiTime.Now
	if c.Identifier != runningIdentifier {
		return consts.ErrExpiredToken
	}

	return c.StandardClaims.Valid()
}

func (c *accessClaims) Credentials() user.CredentialsDTO {
	return c.Crd
}

func (c *accessClaims) Valid() error {
	jwt.TimeFunc = oculiTime.Now
	if c.Identifier != runningIdentifier {
		return consts.ErrExpiredToken
	}

	return c.StandardClaims.Valid()
}
