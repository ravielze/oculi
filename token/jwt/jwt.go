package jwt

import "github.com/ravielze/oculi/token"

type (
	jwtImpl struct {
		token.Encoder
		token.Decoder
	}
)

var (
	runningServerId string = ""
)

func New(key string, alg string, serverId string) token.Tokenizer {
	encoder := NewEncoder(key, alg, serverId)
	decoder := NewDecoder(key)
	runningServerId = serverId

	return &jwtImpl{
		Encoder: encoder,
		Decoder: decoder,
	}
}
