package errors

type (
	DetailedErrors struct {
		ErrorMessage string
		Details      interface{}
	}
)

func (d DetailedErrors) Error() string {
	return d.ErrorMessage
}

func NewDetailedErrors(errorMessage string, details ...interface{}) error {
	return DetailedErrors{errorMessage, details}
}

func InjectDetails(err error, details ...interface{}) error {
	if converted, ok := err.(DetailedErrors); ok {
		var arrInterface []interface{}
		if _, ok2 := converted.Details.([]interface{}); !ok2 {
			arrInterface = []interface{}{converted.Details}
		}
		arrInterface = append(arrInterface, details...)
		converted.Details = arrInterface
		return converted
	}
	return DetailedErrors{
		ErrorMessage: err.Error(),
		Details:      details,
	}
}

func Details(err error) []interface{} {
	if converted, ok := err.(DetailedErrors); ok {
		return converted.Details.([]interface{})
	}
	return []interface{}(nil)
}
