package service

import (
	"github.com/ravielze/oculi/example/model/dao"
	todoDto "github.com/ravielze/oculi/example/model/dto/todo"
	"github.com/ravielze/oculi/request"
)

func (s *service) Create(req request.Context, item todoDto.CreateTodoRequest) (dao.Todo, error) {
	return s.repository.Create(req, item.ToDAO(req))
}
