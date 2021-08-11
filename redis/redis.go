package redis

import (
	ctx "context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/ravielze/oculi/context"
)

type (
	ConnectionInfo struct {
		Address  string
		Password string
		Database int

		// Maximum number of socket connections.
		// Recommended Value: 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
		PoolSize int

		// Minimum number of idle connections which is useful when establishing
		// new connection is slow.
		MinIdleConnections int

		// Dial timeout for establishing new connections.
		// Recommended Value: 5 seconds
		DialTimeout time.Duration

		// Timeout for socket reads. If reached, commands will fail
		// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
		// Recommended Value: 3 seconds.
		ReadTimeout time.Duration

		// Timeout for socket writes. If reached, commands will fail
		// with a timeout instead of blocking.
		// Recommended Value: ReadTimeout.
		WriteTimeout time.Duration

		// Amount of time client waits for connection if all connections
		// are busy before returning an error.
		// Recommended Value is ReadTimeout + 1 second.
		PoolTimeout time.Duration

		// Amount of time after which client closes idle connections.
		// Should be less than server's timeout.
		// Recommended Value: 5 minutes. -1 disables idle timeout check.
		IdleTimeout time.Duration

		// Frequency of idle checks made by idle connections reaper.
		// Recommended Value: 1 minute. -1 disables idle connections reaper,
		// but idle connections are still discarded by the client
		// if IdleTimeout is set.
		IdleCheckFrequency time.Duration
	}

	Redis struct {
		rds      *redis.Client
		Lock     sync.Mutex
		Channels map[string]PubSub
	}

	// NOTE untested
	PubSub interface {
		Publish(msg interface{}) error
		Subscribe(channelName string) error
		Channel() <-chan *redis.Message
		Close() error
	}

	// NOTE untested
	Cache interface {
		// Check if key is exists
		Exists(ctx context.Context, key string) (bool, error)
		// Set an expired time for key within a duration
		Expire(ctx context.Context, key string, ttl time.Duration) error
		// Set an expired time for key to a specific time
		ExpireAt(ctx context.Context, key string, tm time.Time) error
		// Rename a key
		Rename(ctx context.Context, key, newkey string) error

		// Multi Get
		MGet(ctx context.Context, keys ...string) ([]interface{}, error)
		Get(ctx context.Context, key string, obj interface{}) error
		Set(ctx context.Context, key string, value interface{}) error
		SetWithExpiration(ctx context.Context, key string, value interface{}, ttl time.Duration) error
		Del(ctx context.Context, keys ...string) error

		HSet(ctx context.Context, key string, field string, value interface{}) error
		HSetWithExpiration(ctx context.Context, key string, field string, value interface{}, ttl time.Duration) error
		HMSet(ctx context.Context, key string, fieldValue map[string]interface{}) error
		HMSetWithExpiration(ctx context.Context, key string, fieldValue map[string]interface{}, ttl time.Duration) error

		HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error)
		HGet(ctx context.Context, key string, field string, obj interface{}) error
		HDel(ctx context.Context, key string, fields ...string) error

		FlushDatabase(ctx context.Context) error
		Close() error
	}
)

func New(connInfo ConnectionInfo) (*Redis, error) {
	result := redis.NewClient(&redis.Options{
		Addr:               connInfo.Address,
		DB:                 connInfo.Database,
		Password:           connInfo.Password,
		PoolSize:           connInfo.PoolSize,
		MinIdleConns:       connInfo.MinIdleConnections,
		DialTimeout:        connInfo.DialTimeout,
		ReadTimeout:        connInfo.ReadTimeout,
		WriteTimeout:       connInfo.WriteTimeout,
		PoolTimeout:        connInfo.PoolTimeout,
		IdleTimeout:        connInfo.IdleTimeout,
		IdleCheckFrequency: connInfo.IdleCheckFrequency,
	})

	if _, err := result.Ping(ctx.Background()).Result(); err != nil {
		return nil, errors.Wrap(err, "Failed to connect redis!")
	}

	return &Redis{
		rds: result,
	}, nil
}

func (c *Redis) Client() *redis.Client {
	return c.rds
}
