package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) GetByUsername(req request.ReqContext, username string) (dao.User, error) {
	var result dao.User
	if err := req.Transaction().
		Model(dao.User{}).
		Where("username = ?", username).
		Take(&result).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"User.Repository.GetByUsername",
			logs.KeyValue("Username", username),
			logs.KeyValue("Error", err),
		))
		return dao.User{}, err
	}
	return result, nil
}
