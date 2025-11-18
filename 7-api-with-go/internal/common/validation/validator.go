package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string]string {
	formatted := map[string]string{}

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			filed := strings.ToLower(e.Field())
			formatted[filed] = buildMessage(e)

		}
	}

	return formatted
}

func buildMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "max":
		return e.Field() + " exceeds maximum length"
	case "min":
		return e.Field() + " is too short"
	default:
		return "Invalid value for " + e.Field()
	}
}
