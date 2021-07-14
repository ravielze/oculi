package request

import (
	"strconv"
	"strings"

	"github.com/ravielze/oculi/request"
)

// Get query with string value and set it to default if it's empty
func (ctx *reqCtx) Query(query, def string) request.Context {
	if !ctx.HasError() {

		q := ctx.ginCtx.DefaultQuery(query, def)
		if len(q) == 0 || len(strings.TrimSpace(q)) == 0 {
			q = def
		}

		ctx.query[query] = q

	}
	return ctx
}

// Get query with boolean value
func (ctx *reqCtx) QueryBoolean(query string, def bool) request.Context {
	if !ctx.HasError() {

		q := ctx.ginCtx.DefaultQuery(query, strconv.FormatBool(def))
		if (q != strconv.FormatBool(false) && q != strconv.FormatBool(true)) ||
			(len(q) == 0 || len(strings.TrimSpace(q)) == 0) {
			q = strconv.FormatBool(def)
		}

		ctx.query[query], _ = strconv.ParseBool(q)

	}
	return ctx
}

// Get array queries
func (ctx *reqCtx) QueryArray(query string) request.Context {
	if !ctx.HasError() {

		ctx.query[query] = []string(ctx.ginCtx.QueryArray(query))

	}
	return ctx
}
