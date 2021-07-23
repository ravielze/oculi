package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Update(req request.Context, todoId uint64, request map[string]interface{}) error {
	if err := req.Transaction().Model(dao.User{}).
		Where("owner_id = ?", req.Identifier()).
		Where("id = ?", todoId).
		Updates(request).Error(); err != nil {
		return err
	}
	return nil
}
