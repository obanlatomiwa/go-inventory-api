package models

import "github.com/go-playground/validator/v10"

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"

	case "gt":
		return "The value of " + err.Field() + " must be greater than " + err.Param()

	case "gte":
		return "The value of " + err.Field() + " must be greater than or equals" + err.Param()

	case "email":
		return "The value of " + err.Field() + " must be a valid email address"

	case "min":
		return "The minimum length of " + err.Field() + " is equals " + err.Param()

	default:
		return "Validation error in" + err.Field()
	}
}
