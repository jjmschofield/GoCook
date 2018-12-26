package validate

import "gopkg.in/go-playground/validator.v9"

func UuidString(uuid string) (isValid bool, validationMessage string) {
	validate := validator.New()

	uuidValidationErr := validate.Var(uuid, "required,uuid4")

	if uuidValidationErr != nil {
		return false, uuidValidationErr.Error()
	}

	return true, ""
}
