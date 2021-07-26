package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ravielze/oculi/common/model/dto/user"
	consts "github.com/ravielze/oculi/constant/errors"
	"github.com/ravielze/oculi/constant/oculiTime"
)

type (
	claims struct {
		*jwt.StandardClaims
		Crd      user.CredentialsDTO `json:"credentials"`
		ServerId string              `json:"server_id"`
	}
)

func (c *claims) Credentials() user.CredentialsDTO {
	return c.Crd
}

func (c *claims) Valid() error {
	jwt.TimeFunc = oculiTime.Now
	if c.ServerId != runningServerId {
		return consts.ErrExpiredToken
	}

	return c.StandardClaims.Valid()
}
