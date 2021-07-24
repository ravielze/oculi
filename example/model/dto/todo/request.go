package todo

import (
	"github.com/ravielze/oculi/common/model/dto"
	"github.com/ravielze/oculi/example/model/dao"
	"github.com/ravielze/oculi/request"
)

type (
	CreateTodoRequest struct {
		Title       string `json:"title" binding:"required,gte=8"`
		Description string `json:"description" binding:"max=1024"`
	}

	UpdateTodoRequest struct {
		ID          uint64 `json:"id" binding:"required,gt=0"`
		Title       string `json:"title" binding:"required,gte=8,max=256"`
		Description string `json:"description" binding:"max=1024"`
	}
)

func (i CreateTodoRequest) ToDAO(req request.Context) dao.Todo {
	return dao.Todo{
		OwnerID:     req.Identifier(),
		Title:       i.Title,
		Description: &i.Description,
		IsDone:      false,
	}
}

func (i UpdateTodoRequest) ToUpdateMapRequest() dto.Map {
	return dto.Map{
		"id":          i.ID,
		"title":       i.Title,
		"description": i.Description,
	}
}
