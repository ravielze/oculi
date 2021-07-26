package errors

import (
	"errors"
	"net/http"
)

type (
	Mappers []MapOption

	MapOption struct {
		Code int
		Err  error
	}
)

func (m MapOption) Error() string {
	return m.Err.Error()
}

func Transform(err error, m Mappers) MapOption {
	for i := 0; i < len(m); i++ {
		if errors.Is(err, m[i].Err) {
			return m[i]
		}
	}
	return MapOption{
		Code: http.StatusInternalServerError,
		Err:  err,
	}
}
