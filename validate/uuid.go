package validate

import "gopkg.in/go-playground/validator.v9"

func UuidString(uuid string) (isValid bool, validationMessage string){
	validate := validator.New()

	idValidationErr := validate.Var(uuid, "required,uuid4")

	if idValidationErr != nil {
		return false, idValidationErr.Error()
	}

	return true, ""
}