package jwt

import (
	"fmt"
	"time"

	"github.com/ravielze/oculi/token"
)

type (
	jwtImpl struct {
		token.Encoder
		token.Decoder
	}
)

var (
	runningServerId string = ""
)

func GenerateIdentifier(state, restartIntegrityMinute int, name string) string {
	if restartIntegrityMinute <= 0 || restartIntegrityMinute > 60 {
		panic("restartIntegrityMinute is only allowed between 1 and 60")
	}
	now := time.Now()
	x := now.Hour()*100 + (now.Minute() / restartIntegrityMinute)
	identifier := fmt.Sprintf("%d/%s/%d/%s", state, name, x, now.Format("02012006"))
	return identifier
}

func New(key string, alg string, serverId string) token.Tokenizer {
	encoder := NewEncoder(key, alg, serverId)
	decoder := NewDecoder(key)
	runningServerId = serverId

	return &jwtImpl{
		Encoder: encoder,
		Decoder: decoder,
	}
}
