package todo

import (
	"github.com/ravielze/oculi/common/functions"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Delete(req request.Context) error {
	data := *req.Data()
	id := functions.Atoi(data["parameter.id"], 0)
	return h.domain.Todo.Delete(req, id)
}
