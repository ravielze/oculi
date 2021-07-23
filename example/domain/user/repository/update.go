package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Update(req request.Context, userId uint64, request map[string]interface{}) error {
	if err := req.Transaction().Model(dao.User{}).
		Where("id", userId).
		Updates(request).Error(); err != nil {
		return err
	}
	return nil
}
