package pkg_validation

import "github.com/go-playground/validator/v10"

type (
	XValidator struct {
		validator *validator.Validate
	}

	ErrorResponse struct {
		FailedField string
		Tag         string
	}
)

//Initiate Validator
func NewXValidator() XValidator {
	return XValidator{
		validator: validator.New(),
	}
}
