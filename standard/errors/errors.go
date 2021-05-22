package stderr

import (
	"fmt"
)

type (
	BasicError struct {
		detail string
	}

	SpecifiedError struct {
		thing  string
		detail string
	}

	MultiError struct {
		errors []error
	}

	A interface{}
)

func (e BasicError) Error() string {
	return fmt.Sprintf("detail: '%s'", e.detail)
}

func (e SpecifiedError) Error() string {
	return fmt.Sprintf("thing: '%s', detail: '%s'", e.thing, e.detail)
}

func (e MultiError) Error() string {
	return fmt.Sprint(e.errors)
}

// Create new basic standard error
func New(detail string) error {
	return BasicError{
		detail: detail,
	}
}

// Create new specific standard error
func NewSpecific(thing string, detail string) error {
	return SpecifiedError{
		thing:  thing,
		detail: detail,
	}
}

// Wrap multiple errors into an error
func Wrap(errors ...error) error {
	var errFinal []error
	for _, e := range errors {
		if me, ok := e.(MultiError); ok {
			errFinal = append(errFinal, me.errors...)
		} else {
			errFinal = append(errFinal, e)
		}
	}
	return MultiError{
		errors: errFinal,
	}
}
