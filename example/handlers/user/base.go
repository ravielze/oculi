package user

import (
	"github.com/ravielze/oculi/example/domain"
	userDto "github.com/ravielze/oculi/example/model/dto/user"
	"github.com/ravielze/oculi/example/resources"
	"github.com/ravielze/oculi/request"
)

type (
	handler struct {
		domain   domain.Domain
		resource resources.Resource
	}

	Handler interface {
		Login(req request.Context, item userDto.LoginRequest) (userDto.CredentialResponse, error)
		Register(req request.Context, item userDto.RegisterRequest) error
		Check(req request.EchoContext) (userDto.UserResponse, error)
	}
)

func NewHandler(domain domain.Domain, resource resources.Resource) (Handler, error) {
	return &handler{
		domain:   domain,
		resource: resource,
	}, nil
}
