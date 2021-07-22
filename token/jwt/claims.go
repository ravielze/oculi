package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ravielze/oculi/common/model/dto/user"
	"github.com/ravielze/oculi/constant/oculiTime"
)

type (
	claims struct {
		*jwt.StandardClaims
		Crd user.CredentialsDTO `json:"credentials"`
	}
)

func (c *claims) Credentials() user.CredentialsDTO {
	return c.Crd
}

func (c *claims) Valid() error {
	jwt.TimeFunc = oculiTime.Now

	return c.StandardClaims.Valid()
}
