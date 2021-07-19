package external

import (
	"github.com/ravielze/oculi/validator"
	v10 "github.com/ravielze/oculi/validator/v10"
)

func NewValidator() (validator.Validator, error) {
	return v10.New()
}
