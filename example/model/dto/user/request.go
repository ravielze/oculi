package user

import (
	"github.com/ravielze/oculi/example/model/dao"
)

type (
	RegisterRequest struct {
		Username string `json:"username" binding:"required,min=8,alphanum"`
		Password string `json:"password" binding:"required,min=8"`
	}

	LoginRequest struct {
		Username string `json:"username" binding:"required,min=8,alphanum"`
		Password string `json:"password" binding:"required,min=8"`
	}
)

func (r RegisterRequest) ToDAO() dao.User {
	return dao.User{
		Username: r.Username,
		Password: r.Password,
	}
}
