package response

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	oculiContext "github.com/ravielze/oculi/context"
	"github.com/ravielze/oculi/request"
	oculiValidator "github.com/ravielze/oculi/validator"
)

type (
	standardResponse struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}

	errorResponse struct {
		Code   int         `json:"code"`
		Errors interface{} `json:"error"`
	}

	errorField struct {
		Field   string `json:"field"`
		Reason  string `json:"reason"`
		Message string `json:"message"`
	}

	responder struct {
		validator     oculiValidator.Validator
		isDevelopment bool
	}

	Responder interface {
		NewJSONResponse(ctx *oculiContext.Context, req request.Context, data interface{}) error
	}

	Response interface{}
)

func New(validator oculiValidator.Validator, isDevelopment bool) Responder {
	return &responder{
		validator:     validator,
		isDevelopment: isDevelopment,
	}
}
func (r *responder) NewJSONResponse(ctx *oculiContext.Context, req request.Context, data interface{}) error {
	if req != nil {
		ctx.Merge(req)
	}
	return r.newJSON(ctx, data)
}

func (r *responder) newJSON(ctx *oculiContext.Context, data interface{}) error {
	var resp Response
	if ctx.ResponseCode() >= 400 || ctx.HasError() {
		resp = r.handleError(ctx.ResponseCode(), ctx.Errors())
	} else if data == nil {
		resp = standardResponse{
			Code: ctx.ResponseCode(),
			Data: "success",
		}
	} else {
		resp = standardResponse{
			Code: ctx.ResponseCode(),
			Data: data,
		}
	}
	if r.isDevelopment {
		return ctx.JSONPretty(ctx.ResponseCode(), resp, " ")
	}
	return ctx.JSON(ctx.ResponseCode(), resp)
}

func (r *responder) handleError(responseCode int, data []error) errorResponse {
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

func (r *responder) buildErrors(responseCode int, data []error) (string, []errorField) {
	if responseCode != http.StatusBadRequest {
		return data[0].Error(), nil
	}
	err, ok := data[0].(validator.ValidationErrors)
	if ok {
		errors := make([]errorField, len(err))
		for i := range errors {
			errors[i] = errorField{
				Field:   err[i].Namespace(),
				Reason:  err[i].Tag(),
				Message: err[i].Translate(*r.validator.Translator()),
			}
		}
		return "", errors
	}
	return data[0].Error(), nil
}
