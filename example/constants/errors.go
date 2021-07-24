package constants

import "errors"

var (
	ErrTodoNotFound           = errors.New("todo not found")
	ErrTodoAlreadyDoneState   = errors.New("todo is already at done state")
	ErrTodoAlreadyUndoneState = errors.New("todo is already at undone state")
)
