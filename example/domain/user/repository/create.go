package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Create(req request.Context, user dao.User) (dao.User, error) {
	if err := req.Transaction().Create(&user).Error(); err != nil {
		return dao.User{}, err
	}
	return user, nil
}
