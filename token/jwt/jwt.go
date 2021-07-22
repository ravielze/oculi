package jwt

import "github.com/ravielze/oculi/token"

type (
	jwtImpl struct {
		token.Encoder
		token.Decoder
	}
)

func New(key string, alg string) (token.Tokenizer, error) {
	encoder := NewEncoder(key, alg)
	decoder := NewDecoder(key)

	return &jwtImpl{
		Encoder: encoder,
		Decoder: decoder,
	}, nil
}
