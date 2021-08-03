package service

import (
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/request"
)

func (s *service) GetAllByOwner(req request.ReqContext) ([]dao.Todo, error) {
	return s.repository.GetAllByOwner(req)
}
