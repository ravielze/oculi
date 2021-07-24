package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Delete(req request.Context, todoId uint64) error {
	if err := req.Transaction().
		Where("owner_id = ?", req.Identifier()).
		Delete(dao.Todo{ID: todoId}).Error(); err != nil {
		return err
	}
	return nil
}
