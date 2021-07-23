package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/request"
)

func (r *repository) GetByID(req request.Context, userId uint64) (dao.User, error) {
	user := dao.User{ID: userId}
	if err := req.Transaction().
		Take(&user).Error(); err != nil {
		return dao.User{}, err
	}
	return user, nil
}
