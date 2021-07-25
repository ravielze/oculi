package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/request"
)

func (r *repository) GetByID(req request.Context, todoId uint64) (dao.Todo, error) {
	todo := dao.Todo{ID: todoId}
	if err := req.Transaction().
		Where("owner_id = ?", req.Identifier()).
		Take(&todo).Error(); err != nil {
		return dao.Todo{}, err
	}
	return todo, nil
}
