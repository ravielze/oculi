package request

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ravielze/oculi/persistent/sql"
	"github.com/ravielze/oculi/request"
)

type (
	reqCtx struct {
		ctx          context.Context
		ec           echo.Context
		db           sql.API
		tx           sql.API
		errors       []error
		responseCode int
	}
)

func New(ec echo.Context, db sql.API) request.Context {
	return &reqCtx{
		ctx:          context.Background(),
		ec:           ec,
		tx:           nil,
		db:           db,
		errors:       make([]error, 0),
		responseCode: 200,
	}
}

func NewWithCtx(ec echo.Context, ctx context.Context, db sql.API) request.Context {
	return &reqCtx{
		ctx:          ctx,
		ec:           ec,
		tx:           nil,
		db:           db,
		errors:       make([]error, 0),
		responseCode: 200,
	}
}

func (r *reqCtx) Echo() echo.Context {
	return r.ec
}

func (r *reqCtx) Context() context.Context {
	return r.ctx
}

func (r *reqCtx) HasError() bool {
	return len(r.errors) > 0
}

func (r *reqCtx) AddError(responseCode int, err ...error) {
	if r.responseCode < 400 {
		r.responseCode = responseCode
	}
	r.errors = append(r.errors, err...)
}

func (r *reqCtx) SetResponseCode(code int) {
	r.responseCode = code
}

func (r *reqCtx) ResponseCode() int {
	return r.responseCode
}
func (r *reqCtx) Error() error {
	if len(r.errors) > 0 {
		return r.errors[0]
	}
	return nil
}

func (r *reqCtx) HasTransaction() bool {
	return r.tx != nil
}

func (r *reqCtx) Transaction() sql.API {
	return r.tx
}

func (r *reqCtx) NewTransaction() sql.API {
	r.tx = r.db.Begin()
	return r.tx
}

func (r *reqCtx) CommitTransaction() sql.API {
	if r.tx == nil {
		return nil
	}

	r.tx.Commit()
	return r.tx
}

func (r *reqCtx) RollbackTransaction() sql.API {
	if r.tx == nil {
		return nil
	}

	r.tx.Rollback()
	return r.tx
}
