package validate

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CustomMessages map[string]string

type ValidationError struct {
	Field           string
	Tag             string
	RawMessage      string
	FriendlyMessage string
}

type ValidationResult struct {
	Errors []ValidationError
	Ok     bool
}

func Struct(data any, customMessages CustomMessages) ValidationResult {
	result := ValidationResult{
		Ok: true,
	}

	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		var errs []ValidationError
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, err := range validationErrors {
				fieldTag := fmt.Sprintf("%s.%s", err.StructField(), err.Tag())
				errs = append(errs, ValidationError{
					Field:           err.StructField(),
					Tag:             err.Tag(),
					RawMessage:      err.Error(),
					FriendlyMessage: customMessages[fieldTag],
				})
			}
		}
		result.Errors = errs
		result.Ok = false
	}
	return result
}
