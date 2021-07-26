package todo

import (
	"github.com/ravielze/oculi/common/functions/typeutils"
	"github.com/ravielze/oculi/common/model/dto"
	"github.com/ravielze/oculi/example/model/dao"
)

type (
	TodoResponse struct {
		ID          uint64 `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		IsDone      bool   `json:"is_done"`
		dto.BaseModelSoftDelete
	}

	TodosResponse []TodoResponse
)

func NewTodoResponse(t dao.Todo) TodoResponse {
	return TodoResponse{
		BaseModelSoftDelete: dto.NewBaseModelSoftDelete(t.BaseModelSoftDelete),
		ID:                  t.ID,
		Title:               t.Title,
		IsDone:              t.IsDone,
		Description:         typeutils.String(t.Description, ""),
	}
}

func NewTodosResponse(t []dao.Todo) TodosResponse {
	if t == nil {
		return TodosResponse{}
	}
	result := make([]TodoResponse, len(t))
	for i, x := range t {
		result[i] = NewTodoResponse(x)
	}
	return TodosResponse(result)
}
