package jwt

import "github.com/ravielze/oculi/token"

type (
	jwtImpl struct {
		Encoder token.Encoder
		Decoder token.Decoder
	}

	JWT interface {
		token.Encoder
		token.Decoder
	}
)

func NewJWT(key string, alg string) (JWT, error) {
	return nil, nil
}
