package cache

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	consts "github.com/ravielze/oculi/constant/errors"
	"github.com/ravielze/oculi/encoding/jsoniter"
	errOculi "github.com/ravielze/oculi/errors"
	"github.com/ravielze/oculi/redis"
	"github.com/ravielze/oculi/request"
)

type (
	impl struct {
		cl *redis.Redis
	}
)

var json = jsoniter.New()

// NOTE untested
func New(client *redis.Redis) (redis.Cache, error) {
	if client == nil {
		return nil, errors.New("redis is not connected")
	}
	return &impl{client}, nil
}

func (i *impl) Exists(ctx request.ReqContext, key string) (bool, error) {
	r, err := i.cl.Exists(ctx.Context(), key).Result()
	if err != nil {
		return false, err
	}
	return (r > 0), err
}

func (i *impl) Expire(ctx request.ReqContext, key string, ttl time.Duration) error {
	if ttl == -1 {
		return nil
	}

	ok, err := i.cl.Expire(ctx.Context(), key, ttl).Result()
	if err != nil {
		return err
	}
	if !ok {
		return consts.ErrFailedToSetExpiryTime
	}
	return nil
}

func (i *impl) ExpireAt(ctx request.ReqContext, key string, tm time.Time) error {
	ok, err := i.cl.ExpireAt(ctx.Context(), key, tm).Result()
	if err != nil {
		return err
	}
	if !ok {
		return consts.ErrFailedToSetExpiryTime
	}
	return nil
}

func (i *impl) Rename(ctx request.ReqContext, key, newkey string) error {
	_, err := i.cl.Rename(ctx.Context(), key, newkey).Result()
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) MGet(ctx request.ReqContext, keys ...string) ([]interface{}, error) {
	return i.cl.MGet(ctx.Context(), keys...).Result()
}

func (i *impl) Get(ctx request.ReqContext, key string, obj interface{}) error {
	k := reflect.ValueOf(obj).Kind()
	if k != reflect.Ptr && k != reflect.Slice {
		return consts.ErrInvalidValue
	}

	data, err := i.cl.Get(ctx.Context(), key).Result()
	if err != nil {
		return err
	}

	if data == "" {
		return nil
	}
	return json.Unmarshal([]byte(data), obj)
}

func (i *impl) Set(ctx request.ReqContext, key string, value interface{}) error {
	return i.SetWithExpiration(ctx, key, value, 0)
}

func (i *impl) SetWithExpiration(ctx request.ReqContext, key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = i.cl.Set(ctx.Context(), key, data, ttl).Result()
	if err != nil {
		return err
	}

	return nil
}

func (i *impl) Del(ctx request.ReqContext, keys ...string) error {
	_, err := i.cl.Del(ctx.Context(), keys...).Result()
	if err != nil {
		return err
	}
	return err
}

func (i *impl) HSet(ctx request.ReqContext, key string, field string, value interface{}) error {
	return i.HSetWithExpiration(ctx, key, field, value, -1)
}

func (i *impl) HSetWithExpiration(ctx request.ReqContext, key string, field string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = i.cl.HSet(ctx.Context(), key, []string{field, string(data)}).Result()
	if err != nil {
		return err
	}

	if err = i.Expire(ctx, key, ttl); err != nil {
		return err
	}
	return nil
}

func (i *impl) HMSet(ctx request.ReqContext, key string, fieldValue map[string]interface{}) error {
	return i.HMSetWithExpiration(ctx, key, fieldValue, -1)
}

func (i *impl) HMSetWithExpiration(ctx request.ReqContext, key string, fieldValue map[string]interface{}, ttl time.Duration) error {
	translatedMap := make(map[string]interface{})
	for field, value := range fieldValue {
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		translatedMap[field] = data
	}

	_, err := i.cl.HSet(ctx.Context(), key, translatedMap).Result()
	if err != nil {
		return err
	}

	if err = i.Expire(ctx, key, ttl); err != nil {
		return err
	}
	return nil
}

func (i *impl) HMGet(ctx request.ReqContext, key string, fields ...string) ([]interface{}, error) {
	return i.cl.HMGet(ctx.Context(), key, fields...).Result()
}

func (i *impl) HGet(ctx request.ReqContext, key string, field string, obj interface{}) error {
	k := reflect.ValueOf(obj).Kind()
	if k != reflect.Ptr && k != reflect.Slice {
		return consts.ErrInvalidValue
	}

	data, err := i.cl.HGet(ctx.Context(), key, field).Result()
	if err != nil {
		return err
	}

	if data == "" {
		return nil
	}

	return json.Unmarshal([]byte(data), obj)
}

func (i *impl) HDel(ctx request.ReqContext, key string, fields ...string) error {
	_, err := i.cl.HDel(ctx.Context(), key, fields...).Result()
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) FlushDatabase(ctx request.ReqContext) error {
	result, err := i.cl.FlushDB(ctx.Context()).Result()
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

	if err := i.cl.Close(); err != nil {
		return errOculi.NewDetailedErrors(consts.ErrFailedToCloseRedis, err.Error())
	}
	return nil
}
