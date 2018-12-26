package validate

import "gopkg.in/go-playground/validator.v9"

func Struct(obj interface{}) (isValid bool, validationMessage string) {
	validate := validator.New()

	structValidationErr := validate.Struct(obj)

	if structValidationErr != nil {
		return false, structValidationErr.Error()
	}

	return true, ""
}
