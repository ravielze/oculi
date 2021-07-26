package user

import (
	userDto "github.com/ravielze/oculi/example/model/dto/user"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Register(req request.Context, item userDto.RegisterRequest) error {
	return h.domain.User.Register(req, item)
}
