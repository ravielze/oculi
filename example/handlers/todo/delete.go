package todo

import (
	"github.com/ravielze/oculi/common/functions"
	consts "github.com/ravielze/oculi/constant/key"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Delete(req request.ReqContext) error {
	data := *req.Data()
	id := functions.Atoi(data[consts.ParameterPrefix("id")].(string), 0)
	return h.domain.Todo.Delete(req, id)
}
