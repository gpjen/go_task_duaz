package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateInput(input interface{}) []string {
	var validationErrors []string

	err := validator.New().Struct(input)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("%s is %s", e.Field(), e.ActualTag())
			validationErrors = append(validationErrors, message)
		}
	}

	return validationErrors
}
