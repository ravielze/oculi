package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Create(req request.Context, user dao.User) (dao.User, error) {
	if err := req.Transaction().Create(&user).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"User.Repository.Create",
			logs.KeyValue("User", user),
			logs.KeyValue("Error", err),
		))
		return dao.User{}, err
	}
	return user, nil
}
