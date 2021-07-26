package repository

import (
	"github.com/ravielze/oculi/common/model/dto"
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Update(req request.Context, todoId uint64, request dto.Map) error {
	if err := req.Transaction().Model(dao.Todo{}).
		Where("owner_id = ?", req.Identifier()).
		Where("id = ?", todoId).
		Updates(request.ToMap()).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Todo.Repository.Update",
			logs.KeyValue("ID", todoId),
			logs.KeyValue("RequestIdentifier", req.Identifier()),
			logs.KeyValue("RequestMap", request.ToMap()),
			logs.KeyValue("Error", err),
		))
		return err
	}
	return nil
}
