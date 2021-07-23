package external

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ravielze/oculi/example/config"
	"github.com/ravielze/oculi/token"
	oculiJWT "github.com/ravielze/oculi/token/jwt"
)

func NewTokenizer(config *config.Env) token.Tokenizer {
	return oculiJWT.New(config.JWTKey, jwt.SigningMethodES256.Name)
}
