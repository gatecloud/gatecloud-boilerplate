package validations

import validator "gopkg.in/go-playground/validator.v8"

// InitValidation inits a validation handler
func InitValidation() *validator.Validate {
	config := &validator.Config{
		TagName: "validate",
	}

	return validator.New(config)
}
