package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Delete(req request.Context, todoId uint64) error {
	if err := req.Transaction().
		Where("owner_id = ?", req.Identifier()).
		Delete(&dao.Todo{ID: todoId}).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Todo.Repository.Delete",
			logs.KeyValue("ID", todoId),
			logs.KeyValue("RequestIdentifier", req.Identifier()),
			logs.KeyValue("Error", err),
		))
		return err
	}
	return nil
}
