package consts

import "errors"

var (
	ErrFailedToSetExpiryTime    = errors.New("failed to set expiry time")
	ErrFailedToFlushRedis       = errors.New("failed to flush redis")
	ErrFailedToCloseRedis       = "failed to close redis"
	ErrFailedToCloseRedisPubsub = "failed to close redis pubsub"
)
