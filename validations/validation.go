package validations

import validator "gopkg.in/go-playground/validator.v8"

func InitValidation() *validator.Validate {
	config := &validator.Config{TagName: "validate"}
	Validator := validator.New(config)
	return Validator
}
