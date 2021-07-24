package service

import (
	"errors"

	"github.com/ravielze/oculi/example/model/dao"
	userDto "github.com/ravielze/oculi/example/model/dto/user"
	"github.com/ravielze/oculi/request"
)

var (
	emptyUser = dao.User{}
)

func (s *service) Register(req request.Context, item userDto.RegisterRequest) error {
	register := item.ToDAO()
	if user, _ := s.repository.
		GetByUsername(req, register.Username); user != emptyUser {
		return errors.New("account with that username already exist")
	}

	hashed, err := s.hash.Hash(register.Password)
	if err != nil {
		return err
	}
	register.Password = hashed

	_, err = s.repository.Create(req, register)
	return err
}