package service

import (
	"github.com/ravielze/oculi/request"
)

func (s *service) Delete(req request.ReqContext, todoId uint64) error {
	_, err := s.repository.GetByID(req, todoId)
	if err != nil {
		return err
	}
	return s.repository.Delete(req, todoId)
}
