package token

import "github.com/ravielze/oculi/common/model/dto/user"

type (
	Encoder interface {
		Encode(claims Claims) (string, error)
		CreateClaims(credentials user.CredentialsDTO, exp int64) Claims
	}

	Decoder interface {
		Decode(token string) (Claims, error)
	}

	Claims interface {
		Credentials() user.CredentialsDTO
		Valid() error
	}
)
