package user

import (
	"github.com/ravielze/oculi/common/model/dto"
	"github.com/ravielze/oculi/example/model/dao"
)

type (
	UserResponse struct {
		ID       uint64 `json:"id"`
		Username string `json:"username"`
		dto.BaseModel
	}

	CredentialResponse struct {
		User  UserResponse `json:"user"`
		Token string       `json:"token"`
	}
)

func NewUserResponse(user dao.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		BaseModel: dto.NewBaseModel(user.BaseModel),
	}
}

func NewCredentialResponse(user dao.User, token string) CredentialResponse {
	return CredentialResponse{
		User:  NewUserResponse(user),
		Token: token,
	}
}
