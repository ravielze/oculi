package request

import (
	"github.com/labstack/echo/v4"
	uDto "github.com/ravielze/oculi/common/model/dto/user"
	consts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/persistent/sql"
	"github.com/ravielze/oculi/request"
)

type (
	reqCtx struct {
		request.Context
		ec echo.Context
	}
)

func New(ec echo.Context, db sql.API) request.EchoContext {
	r := &reqCtx{
		Context: request.NewBase(db),
		ec:      ec,
	}
	if item := ec.Get(consts.KeyCredentials); item != nil {
		if c, ok := item.(uDto.CredentialsDTO); ok {
			r.SetIdentifier(c.ID)
		}
	}
	return r
}

func NewWithIdentifier(ec echo.Context, db sql.API, identifier uint64) request.EchoContext {
	return &reqCtx{
		Context: request.NewBaseWithIdentifier(identifier, db),
		ec:      ec,
	}
}

func (r *reqCtx) Echo() echo.Context {
	return r.ec
}
