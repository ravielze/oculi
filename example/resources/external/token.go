package external

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/ravielze/oculi/example/config"
	"github.com/ravielze/oculi/token"
	oculiJWT "github.com/ravielze/oculi/token/jwt"
)

func NewTokenizer(config *config.Env) token.Tokenizer {
	identifier := oculiJWT.GenerateIdentifier(config.ServiceState, 5, config.ServiceName)
	fmt.Println(identifier)
	return oculiJWT.New(config.JWTKey, jwt.SigningMethodHS256.Name, identifier)
}
