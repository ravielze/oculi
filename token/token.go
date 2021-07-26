package token

import (
	"net/http"
	"time"

	"github.com/ravielze/oculi/common/model/dto/user"
)

type (
	Encoder interface {
		Encode(claims Claims) (string, error)
		CreateClaims(credentials user.CredentialsDTO, exp time.Duration) Claims
		CreateAndEncode(credentials user.CredentialsDTO, exp time.Duration) (string, error)
	}

	Decoder interface {
		Decode(tkn string) (Claims, error)
		DecodeHttpRequest(req *http.Request) (Claims, error)
	}

	Claims interface {
		Credentials() user.CredentialsDTO
		Valid() error
	}

	Tokenizer interface {
		Encoder
		Decoder
	}
)
