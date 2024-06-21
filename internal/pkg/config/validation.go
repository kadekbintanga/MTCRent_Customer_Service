package config

import (
	"github.com/globalxtreme/gobaseconf/middleware"
	"github.com/go-playground/validator/v10"
)

func InitValidation() {
	v := middleware.Validator{}
	v.RegisterValidation(func(validate *validator.Validate) {
		// Enter your custom validation rules
	})
}
