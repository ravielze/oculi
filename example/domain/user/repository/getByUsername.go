package repository

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/request"
)

func (r *repository) GetByUsername(req request.Context, username string) (dao.User, error) {
	var result dao.User
	if err := req.Transaction().
		Model(dao.User{}).
		Where("username = ?", username).
		Take(&result).Error(); err != nil {
		return dao.User{}, err
	}
	return result, nil
}
