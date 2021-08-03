package service

import (
	"github.com/ravielze/oculi/example/constants"
	"github.com/ravielze/oculi/example/model/dao"
	userDto "github.com/ravielze/oculi/example/model/dto/user"
	"github.com/ravielze/oculi/request"
)

var (
	emptyUser = dao.User{}
)

func (s *service) Register(req request.ReqContext, item userDto.RegisterRequest) error {
	register := item.ToDAO()
	if user, _ := s.repository.
		GetByUsername(req, register.Username); user != emptyUser {
		return constants.ErrUserRegistered
	}

	hashed, err := s.hash.Hash(register.Password)
	if err != nil {
		return err
	}
	register.Password = hashed

	_, err = s.repository.Create(req, register)
	return err
}
