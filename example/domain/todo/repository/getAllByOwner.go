package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/request"
)

func (r *repository) GetAllByOwner(req request.Context) ([]dao.Todo, error) {
	var todo []dao.Todo
	if err := req.Transaction().
		Where("owner_id = ?", req.Identifier()).
		Find(&todo).Error(); err != nil {
		return nil, err
	}
	return todo, nil
}
