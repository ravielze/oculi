package todo

type (
	CreateTodoRequest struct {
		Title       string `json:"title" binding:"required,gte=8"`
		Description string `json:"description"`
	}

	UpdateTodoRequest struct {
		ID          uint64 `json:"id" binding:"required,gt=0"`
		Title       string `json:"title" binding:"required,gte=8"`
		Description string `json:"description,omitempty"`
		IsDone      bool   `json:"is_done"`
	}
)
