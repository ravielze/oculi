package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	consts "github.com/ravielze/oculi/constant/errors"
	"github.com/ravielze/oculi/context"
	errOculi "github.com/ravielze/oculi/errors"
	"github.com/ravielze/oculi/redis"
)

type (
	impl struct {
		cl *redis.Redis
	}
)

// NOTE untested
func New(client *redis.Redis) (redis.Cache, error) {
	if client == nil {
		return nil, errors.New("redis is not connected")
	}
	return &impl{client}, nil
}

func (i *impl) Exists(ctx context.Context, key string) (bool, error) {
	r, err := i.cl.Client().Exists(ctx.Context(), key).Result()
	if err != nil {
		return false, err
	}
	return (r > 0), err
}

func (i *impl) Expire(ctx context.Context, key string, ttl time.Duration) error {
	if ttl == -1 {
		return nil
	}

	ok, err := i.cl.Client().Expire(ctx.Context(), key, ttl).Result()
	if err != nil {
		return err
	}
	if !ok {
		return consts.ErrFailedToSetExpiryTime
	}
	return nil
}

func (i *impl) ExpireAt(ctx context.Context, key string, tm time.Time) error {
	ok, err := i.cl.Client().ExpireAt(ctx.Context(), key, tm).Result()
	if err != nil {
		return err
	}
	if !ok {
		return consts.ErrFailedToSetExpiryTime
	}
	return nil
}

func (i *impl) Rename(ctx context.Context, key, newkey string) error {
	_, err := i.cl.Client().Rename(ctx.Context(), key, newkey).Result()
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	return i.cl.Client().MGet(ctx.Context(), keys...).Result()
}

func (i *impl) Get(ctx context.Context, key string, obj interface{}) error {
	k := reflect.ValueOf(obj).Kind()
	if k != reflect.Ptr && k != reflect.Slice {
		return consts.ErrInvalidValue
	}

	data, err := i.cl.Client().Get(ctx.Context(), key).Result()
	if err != nil {
		return err
	}

	if data == "" {
		return nil
	}
	return json.Unmarshal([]byte(data), obj)
}

func (i *impl) Set(ctx context.Context, key string, value interface{}) error {
	return i.SetWithExpiration(ctx, key, value, 0)
}

func (i *impl) SetWithExpiration(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = i.cl.Client().Set(ctx.Context(), key, data, ttl).Result()
	if err != nil {
		return err
	}

	return nil
}

func (i *impl) Del(ctx context.Context, keys ...string) error {
	_, err := i.cl.Client().Del(ctx.Context(), keys...).Result()
	if err != nil {
		return err
	}
	return err
}

func (i *impl) HSet(ctx context.Context, key string, field string, value interface{}) error {
	return i.HSetWithExpiration(ctx, key, field, value, -1)
}

func (i *impl) HSetWithExpiration(ctx context.Context, key string, field string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = i.cl.Client().HSet(ctx.Context(), key, []string{field, string(data)}).Result()
	if err != nil {
		return err
	}

	if err = i.Expire(ctx, key, ttl); err != nil {
		return err
	}
	return nil
}

func (i *impl) HMSet(ctx context.Context, key string, fieldValue map[string]interface{}) error {
	return i.HMSetWithExpiration(ctx, key, fieldValue, -1)
}

func (i *impl) HMSetWithExpiration(ctx context.Context, key string, fieldValue map[string]interface{}, ttl time.Duration) error {
	translatedMap := make(map[string]interface{})
	for field, value := range fieldValue {
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		translatedMap[field] = data
	}

	_, err := i.cl.Client().HSet(ctx.Context(), key, translatedMap).Result()
	if err != nil {
		return err
	}

	if err = i.Expire(ctx, key, ttl); err != nil {
		return err
	}
	return nil
}

func (i *impl) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return i.cl.Client().HMGet(ctx.Context(), key, fields...).Result()
}

func (i *impl) HGet(ctx context.Context, key string, field string, obj interface{}) error {
	k := reflect.ValueOf(obj).Kind()
	if k != reflect.Ptr && k != reflect.Slice {
		return consts.ErrInvalidValue
	}

	data, err := i.cl.Client().HGet(ctx.Context(), key, field).Result()
	if err != nil {
		return err
	}

	if data == "" {
		return nil
	}

	return json.Unmarshal([]byte(data), obj)
}

func (i *impl) HDel(ctx context.Context, key string, fields ...string) error {
	_, err := i.cl.Client().HDel(ctx.Context(), key, fields...).Result()
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) FlushDatabase(ctx context.Context) error {
	result, err := i.cl.Client().FlushDB(ctx.Context()).Result()
	if err != nil {
		return err
	}
	if strings.EqualFold(result, "OK") {
		return consts.ErrFailedToFlushRedis
	}
	return nil
}

func (i *impl) Close() error {
	i.cl.Lock.Lock()
	defer i.cl.Lock.Unlock()

	var errDetails []interface{}
	for chName, ch := range i.cl.Channels {
		if err := ch.Close(); err != nil {
			errDetails = append(errDetails, fmt.Sprintf("%s: %s", chName, err.Error()))
		}
	}
	if len(errDetails) != 0 {
		return errOculi.NewDetailedErrors(consts.ErrFailedToCloseRedisPubsub, errDetails...)
	}

	if err := i.cl.Client().Close(); err != nil {
		return errOculi.NewDetailedErrors(consts.ErrFailedToCloseRedis, err.Error())
	}
	return nil
}
