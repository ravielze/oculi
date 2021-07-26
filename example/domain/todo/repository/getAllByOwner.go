package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) GetAllByOwner(req request.Context) ([]dao.Todo, error) {
	var todo []dao.Todo
	if err := req.Transaction().
		Where("owner_id = ?", req.Identifier()).
		Find(&todo).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Todo.Repository.GetAllByOwner",
			logs.KeyValue("RequestIdentifier", req.Identifier()),
			logs.KeyValue("Error", err),
		))
		return nil, err
	}
	return todo, nil
}
