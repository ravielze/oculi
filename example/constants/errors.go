package constants

import (
	"errors"
	"net/http"

	consts "github.com/ravielze/oculi/constant/errors"
	errorUtil "github.com/ravielze/oculi/errors"
)

var (
	ErrTodoNotFound           = errors.New("todo not found")
	ErrTodoAlreadyDoneState   = errors.New("todo is already at done state")
	ErrTodoAlreadyUndoneState = errors.New("todo is already at undone state")

	ErrUserRegistered = errors.New("account with that username already exist")
	ErrWrongPassword  = errors.New("wrong password")
	ErrNotLoggedIn    = errors.New("not logged in")

	ErrResetUnauthorized = errors.New("reset unauthorized")

	HealthMappers = errorUtil.Mappers{
		{Code: http.StatusUnauthorized, Err: ErrResetUnauthorized},
	}

	UserMappers = errorUtil.Mappers{
		{Code: http.StatusBadRequest, Err: ErrUserRegistered},
		{Code: http.StatusBadRequest, Err: ErrWrongPassword},
		{Code: http.StatusBadRequest, Err: ErrNotLoggedIn},
		{Code: http.StatusBadRequest, Err: consts.ErrRecordNotFound},
	}

	TodoMappers = errorUtil.Mappers{
		{Code: http.StatusBadRequest, Err: ErrTodoNotFound},
		{Code: http.StatusBadRequest, Err: ErrTodoAlreadyDoneState},
		{Code: http.StatusBadRequest, Err: ErrTodoAlreadyUndoneState},
		{Code: http.StatusBadRequest, Err: consts.ErrRecordNotFound},
	}
)
