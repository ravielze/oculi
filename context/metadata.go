package context

import (
	"encoding/json"
	"net/http"
	"time"

	stdcode "github.com/ravielze/oculi/standard/code"
)

func (ctx *Context) Set(key string, value interface{}) *Context {
	ctx.ginCtx.Set(key, value)
	return ctx
}

func (ctx *Context) SetObject(key string, value interface{}) *Context {
	buff, err := json.Marshal(value)
	if err != nil {
		ctx.Error(
			err,
			http.StatusInternalServerError,
			stdcode.SERVER_ERROR,
		)
	} else {
		ctx.Set(key, string(buff))
	}
	return ctx
}

func (ctx *Context) GetObject(key string, object interface{}) *Context {
	strBuff := ctx.GetString(key)
	err := json.Unmarshal([]byte(strBuff), object)
	if err != nil {
		ctx.Error(
			err,
			http.StatusInternalServerError,
			stdcode.SERVER_ERROR,
		)
	}
	return ctx
}

func (ctx *Context) Get(key string) (interface{}, bool) {
	return ctx.ginCtx.Get(key)
}

func (ctx *Context) MustGet(key string) interface{} {
	return ctx.ginCtx.MustGet(key)
}

func (ctx *Context) GetString(key string) string {
	return ctx.ginCtx.GetString(key)
}

func (ctx *Context) GetBool(key string) bool {
	return ctx.ginCtx.GetBool(key)
}

func (ctx *Context) GetInt(key string) int {
	return ctx.ginCtx.GetInt(key)
}

func (ctx *Context) GetInt64(key string) int64 {
	return ctx.ginCtx.GetInt64(key)
}

func (ctx *Context) GetUint(key string) uint {
	return ctx.ginCtx.GetUint(key)
}

func (ctx *Context) GetUint64(key string) uint64 {
	return ctx.ginCtx.GetUint64(key)
}

func (ctx *Context) GetFloat64(key string) float64 {
	return ctx.ginCtx.GetFloat64(key)
}

func (ctx *Context) GetTime(key string) time.Time {
	return ctx.ginCtx.GetTime(key)
}

func (ctx *Context) GetDuration(key string) time.Duration {
	return ctx.ginCtx.GetDuration(key)
}

func (ctx *Context) GetStringSlice(key string) []string {
	return ctx.ginCtx.GetStringSlice(key)
}

func (ctx *Context) GetStringMap(key string) map[string]interface{} {
	return ctx.ginCtx.GetStringMap(key)
}

func (ctx *Context) GetStringMapString(key string) map[string]string {
	return ctx.ginCtx.GetStringMapString(key)
}

func (ctx *Context) GetStringMapStringSlice(key string) map[string][]string {
	return ctx.ginCtx.GetStringMapStringSlice(key)
}
