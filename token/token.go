package token

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/ravielze/oculi/common/model/dto/auth"
)

type (
	Encoder interface {
		AccessEncoder
		RefreshEncoder
		Encode(claims Claims) (string, error)
	}

	Decoder interface {
		AccessDecoder
		RefreshDecoder
	}

	AccessEncoder interface {
		CreateAccessClaims(credentials auth.StandardCredentials, exp time.Duration) (claim Claims)
		CreateAccessAndEncode(credentials auth.StandardCredentials, exp time.Duration) (tokenString string, err error)
	}

	RefreshEncoder interface {
		CreateRefreshClaims(userId uint64, exp time.Duration) (claim Claims, token uuid.UUID)
		CreateRefreshAndEncode(userId uint64, exp time.Duration) (tokenString string, token uuid.UUID, err error)
	}

	AccessDecoder interface {
		DecodeAccess(tokenString string) (claim Claims, err error)
		DecodeAccessHeader(req *http.Request) (claim Claims, err error)
		DecodeAccessCookie(req *http.Request) (claim Claims, err error)
	}

	RefreshDecoder interface {
		DecodeRefresh(tokenString string) (claim Claims, err error)
		DecodeRefreshCookie(req *http.Request) (claim Claims, err error)
	}

	Claims interface {
		Credentials() auth.StandardCredentials
		Valid() error
	}

	Tokenizer interface {
		Encoder
		Decoder
	}
)
