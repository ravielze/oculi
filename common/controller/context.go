package contutils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ravielze/fuzzy-broccoli/common/code"
	"github.com/ravielze/fuzzy-broccoli/common/radix36"
	"github.com/ravielze/fuzzy-broccoli/common/utils"
	uuid "github.com/satori/go.uuid"
)

type ControlChain struct {
	ctx      *gin.Context
	errors   []error
	params   map[string]string
	code     []string
	httpCode []int
	query    map[string]string
}

func NewControlChain(context *gin.Context) *ControlChain {
	return &ControlChain{
		ctx:      context,
		errors:   []error{},
		code:     []string{},
		httpCode: []int{},
		params:   map[string]string{},
		query:    map[string]string{},
	}
}

func (cu *ControlChain) Bind(obj interface{}) *ControlChain {
	if err := cu.ctx.ShouldBind(obj); err != nil {
		cu.errors = append(cu.errors, err)
		cu.httpCode = append(cu.httpCode, http.StatusUnprocessableEntity)
		cu.code = append(cu.code, code.UNCOMPATIBLE_ENTITY)
	}
	return cu
}

func (cu *ControlChain) ParamID(parameter string) *ControlChain {
	p := cu.ctx.Param(parameter)
	if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
		cu.errors = append(cu.errors, fmt.Errorf("parameter '%s' is missing", parameter))
		cu.httpCode = append(cu.httpCode, http.StatusUnprocessableEntity)
		cu.code = append(cu.code, code.PARAMETER_ERROR)
	} else {
		result := radix36.DecodeUUID(p)
		if result != uuid.Nil {
			cu.errors = append(cu.errors, fmt.Errorf("parameter '%s' is not radix36", parameter))
			cu.httpCode = append(cu.httpCode, http.StatusUnprocessableEntity)
			cu.code = append(cu.code, code.PARAMETER_ERROR)
		} else {
			cu.params[parameter] = result.String()
		}
	}
	return cu
}

func (cu *ControlChain) Param(parameter string) *ControlChain {
	p := cu.ctx.Param(parameter)
	if len(p) == 0 || len(strings.TrimSpace(p)) == 0 {
		cu.errors = append(cu.errors, fmt.Errorf("parameter '%s' is missing", parameter))
		cu.httpCode = append(cu.httpCode, http.StatusUnprocessableEntity)
		cu.code = append(cu.code, code.PARAMETER_ERROR)
	} else {
		cu.params[parameter] = p
	}
	return cu
}

func (cu *ControlChain) Query(query, def string) *ControlChain {
	q := cu.ctx.DefaultQuery(query, def)
	if len(q) == 0 || len(strings.TrimSpace(q)) == 0 {
		q = def
	}
	cu.query[query] = q
	return cu
}

func (cu *ControlChain) End() (bool, map[string]string, map[string]string) {
	if len(cu.errors) > 0 && len(cu.httpCode) > 0 && len(cu.code) > 0 {
		err := cu.errors[0]
		httpCode := cu.httpCode[0]
		code := cu.code[0]
		utils.AbortAndResponseData(
			cu.ctx,
			httpCode,
			code,
			err.Error(),
		)
		return false, map[string]string{}, map[string]string{}
	}
	return true, cu.params, cu.query
}
