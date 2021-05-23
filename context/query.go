package context

import (
	"strings"
)

func boolString(x bool) string {
	if x {
		return "true"
	}
	return "false"
}

func stringBool(x string) bool {
	x = strings.ToLower(x)
	return x == "true"
}

// Get query with string value and set it to default if it's empty
func (ctx *Context) Query(query, def string) *Context {
	if !ctx.IsError() {

		q := ctx.ginCtx.DefaultQuery(query, def)
		if len(q) == 0 || len(strings.TrimSpace(q)) == 0 {
			q = def
		}

		ctx.query[query] = q

	}
	return ctx
}

// Get query with boolean value
func (ctx *Context) QueryBoolean(query string, def bool) *Context {
	if !ctx.IsError() {

		q := ctx.ginCtx.DefaultQuery(query, boolString(def))
		if (q != boolString(false) && q != boolString(true)) ||
			(len(q) == 0 || len(strings.TrimSpace(q)) == 0) {
			q = boolString(def)
		}

		ctx.query[query] = stringBool(q)

	}
	return ctx
}

// Get array queries
func (ctx *Context) QueryArray(query string) *Context {
	if !ctx.IsError() {

		ctx.query[query] = []string(ctx.ginCtx.QueryArray(query))

	}
	return ctx
}
