package stderr

import (
	"fmt"
)

type (
	BasicError struct {
		Detail string `json:"detail"`
	}

	SpecifiedError struct {
		Thing  string `json:"thing"`
		Detail string `json:"detail"`
	}

	MultiError struct {
		Errors []error `json:"errors"`
	}
)

func (e BasicError) Error() string {
	return fmt.Sprintf("detail: '%s'", e.Detail)
}

func (e SpecifiedError) Error() string {
	return fmt.Sprintf("thing: '%s', detail: '%s'", e.Thing, e.Detail)
}

func (e MultiError) Error() string {
	return fmt.Sprint(e.Errors)
}

// Create new basic standard error
func New(detail string) error {
	return BasicError{
		Detail: detail,
	}
}

// Create new basic standard error from error
func NewFromError(err error) error {
	return BasicError{
		Detail: err.Error(),
	}
}

// Create new specific standard error
func NewSpecific(thing string, detail string) error {
	return SpecifiedError{
		Thing:  thing,
		Detail: detail,
	}
}

// Wrap multiple errors into an error
func NewMultipleError(errors ...error) error {
	var errFinal []error
	for _, e := range errors {
		if me, ok := e.(MultiError); ok {
			errFinal = append(errFinal, me.Errors...)
		} else {
			errFinal = append(errFinal, e)
		}
	}
	return MultiError{
		Errors: errFinal,
	}
}
