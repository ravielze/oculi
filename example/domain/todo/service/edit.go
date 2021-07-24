package service

import (
	todoDto "github.com/ravielze/oculi/example/model/dto/todo"
	"github.com/ravielze/oculi/request"
)

func (s *service) Edit(req request.Context, item todoDto.UpdateTodoRequest) error {
	_, err := s.repository.GetByID(req, item.ID)
	if err != nil {
		return err
	}

	return s.repository.Update(req, item.ID, item.ToUpdateMapRequest())
}
