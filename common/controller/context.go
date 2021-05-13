package contutils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	uuid "github.com/gofrs/uuid"

	"github.com/gin-gonic/gin"
	code "github.com/ravielze/oculi/common/code"
	"github.com/ravielze/oculi/common/radix36"
	"github.com/ravielze/oculi/common/utils"
)

type Parameters map[string]string
type Queries map[string]string

type ControlChain struct {
	ctx      *gin.Context
	err      error
	params   Parameters
	code     string
	httpCode int
	query    Queries
	isError  bool
}

func NewControlChain(context *gin.Context) *ControlChain {
	return &ControlChain{
		ctx:      context,
		err:      errors.New("not error"),
		code:     "",
		httpCode: -1,
		params:   map[string]string{},
		query:    map[string]string{},
		isError:  false,
	}
}

func (cu *ControlChain) Bind(obj interface{}) *ControlChain {
	if cu.isError {
		return cu
	}
	if err := cu.ctx.ShouldBind(obj); err != nil {
		cu.err = err
		cu.httpCode = http.StatusUnprocessableEntity
		cu.code = code.UNCOMPATIBLE_ENTITY
		cu.isError = true
	}
	return cu
}

func (cu *ControlChain) ParamID(parameter string) *ControlChain {
	if cu.isError {
		return cu
	}
	p := cu.ctx.Param(parameter)
	if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
		cu.err = fmt.Errorf("parameter '%s' is missing", parameter)
		cu.httpCode = http.StatusUnprocessableEntity
		cu.code = code.PARAMETER_ERROR
		cu.isError = true
	} else {
		result := radix36.DecodeUUID(p)
		if result != uuid.Nil {
			cu.err = fmt.Errorf("parameter '%s' is not radix36", parameter)
			cu.httpCode = http.StatusUnprocessableEntity
			cu.code = code.PARAMETER_ERROR
			cu.isError = true
		} else {
			cu.params[parameter] = result.String()
		}
	}
	return cu
}

func (cu *ControlChain) Param(parameter string) *ControlChain {
	if cu.isError {
		return cu
	}
	p := cu.ctx.Param(parameter)
	if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
		cu.err = fmt.Errorf("parameter '%s' is missing", parameter)
		cu.httpCode = http.StatusUnprocessableEntity
		cu.code = code.PARAMETER_ERROR
		cu.isError = true
	} else {
		cu.params[parameter] = p
	}
	return cu
}

func (cu *ControlChain) Query(query, def string) *ControlChain {
	if cu.isError {
		return cu
	}
	q := cu.ctx.DefaultQuery(query, def)
	if len(q) == 0 || len(strings.TrimSpace(q)) == 0 {
		q = def
	}
	cu.query[query] = q
	return cu
}

func (cu *ControlChain) End() (bool, Parameters, Queries) {
	if cu.isError {
		utils.AbortAndResponseData(
			cu.ctx,
			cu.httpCode,
			cu.code,
			cu.err.Error(),
		)
		return false, Parameters{}, Queries{}
	}
	return true, cu.params, cu.query
}
