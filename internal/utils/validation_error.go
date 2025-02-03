package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationError(err validator.ValidationErrors) map[string]string {
	errorMap := make(map[string]string)
	for _, fieldError := range err {
		field := fieldError.Field()
		tag := fieldError.Tag()
		param := fieldError.Param()

		switch tag {
		case "required":
			errorMap[field] = fmt.Sprintf("%s is required", field)
		case "min":
			errorMap[field] = fmt.Sprintf("%s must be at least %s characters long", field, param)
		case "max":
			errorMap[field] = fmt.Sprintf("%s must be at most %s characters long", field, param)
		case "email":
			errorMap[field] = fmt.Sprintf("%s must be a valid email address", field)
		case "numeric":
			errorMap[field] = fmt.Sprintf("%s must be a numeric value", field)
		default:
			errorMap[field] = fmt.Sprintf("%s is invalid", field)
		}
	}

	for field, message := range errorMap {
		fmt.Printf("%s: %s\n", field, message)
	}

	return errorMap
}
