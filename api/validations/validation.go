package validations

import (
	"gatecloud-boilerplate/api/models"

	validator "gopkg.in/go-playground/validator.v8"
)

// InitValidation inits a validation handler
func InitValidation() *validator.Validate {
	config := &validator.Config{
		TagName: "validate",
	}

	validate := validator.New(config)
	validate.RegisterStructValidation(CrossFieldValidation, models.TestAPI{})
	validate.RegisterStructValidation(ReadOnlyValidation, models.TestAPI{})
	return validate
}
