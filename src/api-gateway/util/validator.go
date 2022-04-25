package util

import (
	"github.com/go-playground/validator/v10"
)

func GenerateValidator() *validator.Validate {
	validate := validator.New()

	return validate
}
