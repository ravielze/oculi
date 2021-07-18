package response

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	oculiContext "github.com/ravielze/oculi/context"
	oculiValidator "github.com/ravielze/oculi/validator"
)

type (
	Response interface{}

	standardResponse struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}

	errorResponse struct {
		Code   int         `json:"Code"`
		Errors interface{} `json:"error"`
	}

	errorField struct {
		Field   string `json:"field"`
		Reason  string `json:"reason"`
		Message string `json:"message"`
	}

	Responder struct {
		validator oculiValidator.Validator
	}
)

func New(validator oculiValidator.Validator) *Responder {
	return &Responder{
		validator: validator,
	}
}

func (r *Responder) NewJSON(ctx *oculiContext.Context, data interface{}) error {
	var resp Response
	if ctx.ResponseCode() >= 400 || ctx.HasError() {
		resp = r.handleError(ctx.ResponseCode(), ctx.Errors())
	} else if data == nil {
		return ctx.JSON(ctx.ResponseCode(), nil)
	} else {
		resp = standardResponse{
			Code: ctx.ResponseCode(),
			Data: data,
		}
	}
	return ctx.JSON(ctx.ResponseCode(), resp)
}

func (r *Responder) handleError(responseCode int, data []error) errorResponse {
	if l := len(data); l == 1 {
		return errorResponse{
			Code:   responseCode,
			Errors: data[0].Error(),
		}
	} else if l > 1 {
		msg, errfields := r.buildErrors(responseCode, data)
		if errfields == nil {
			return errorResponse{
				Code:   responseCode,
				Errors: msg,
			}
		}
		return errorResponse{
			Code:   responseCode,
			Errors: errfields,
		}
	}
	return errorResponse{
		Code:   responseCode,
		Errors: "unknown error has occured",
	}
}

func (r *Responder) buildErrors(responseCode int, data []error) (string, []errorField) {
	if responseCode != http.StatusBadRequest {
		return data[0].Error(), nil
	}
	err, ok := data[0].(validator.ValidationErrors)
	if ok {
		errors := make([]errorField, len(err))
		for i := range errors {
			errors[i] = errorField{
				Field:   err[i].StructNamespace(),
				Reason:  err[i].Tag(),
				Message: err[i].Translate(*r.validator.Translator()),
			}
		}
	}
	return data[0].Error(), nil
}
