package consts

import "errors"

var (
	ErrUnclaimedToken = errors.New("unclaimed token")
	ErrNoBearerToken  = errors.New("bearer token not found")
	ErrExpiredToken   = errors.New("token expired")
)
