package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Create(req request.Context, item dao.Todo) (dao.Todo, error) {
	if err := req.Transaction().
		Create(&item).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Todo.Repository.Create",
			logs.KeyValue("Todo", item),
			logs.KeyValue("Error", err),
		))
		return dao.Todo{}, err
	}
	return item, nil
}
