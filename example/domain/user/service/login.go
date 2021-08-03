package service

import (
	commonUserDto "github.com/ravielze/oculi/common/model/dto/user"
	"github.com/ravielze/oculi/example/constants"
	"github.com/ravielze/oculi/example/model/dao"
	userDto "github.com/ravielze/oculi/example/model/dto/user"
	"github.com/ravielze/oculi/request"
)

func (s *service) Login(req request.ReqContext, item userDto.LoginRequest) (dao.User, string, error) {
	user, err := s.repository.GetByUsername(req, item.Username)
	if err != nil {
		return dao.User{}, "", err
	}

	if errPassword := s.hash.Verify(item.Password, user.Password); errPassword != nil {
		return dao.User{}, "", constants.ErrWrongPassword
	}
	token, err := s.resource.Tokenizer.
		CreateAndEncode(
			commonUserDto.CredentialsDTO{
				ID:       user.ID,
				Metadata: user,
			},
			s.resource.Config.JWTExp,
		)

	if err != nil {
		return dao.User{}, "", err
	}
	return user, token, nil
}
