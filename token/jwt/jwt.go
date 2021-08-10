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
	runningIdentifier string = ""
)

// Deprecated: use static identifier instead.
func GenerateIdentifier(state, restartIntegrityMinute int, name string) string {
	if restartIntegrityMinute <= 0 || restartIntegrityMinute > 60 {
		panic("restartIntegrityMinute is only allowed between 1 and 60")
	}
	now := time.Now()
	x := now.Hour()*100 + (now.Minute() / restartIntegrityMinute)
	identifier := fmt.Sprintf("%d/%s/%d/%s", state, name, x, now.Format("02012006"))
	return identifier
}

func New(key string, alg string, identifier string) token.Tokenizer {
	encoder := NewEncoder(key, alg, identifier)
	decoder := NewDecoder(key)
	runningIdentifier = identifier

	return &jwtImpl{
		Encoder: encoder,
		Decoder: decoder,
	}
}
