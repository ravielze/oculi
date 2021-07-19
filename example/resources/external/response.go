package external

import (
	"github.com/ravielze/oculi/response"
	"github.com/ravielze/oculi/validator"
)

func NewResponder(v validator.Validator) *response.Responder {
	return response.New(v)
}
