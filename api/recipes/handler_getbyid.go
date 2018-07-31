package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/respond"
	"github.com/jjmschofield/GoCook/validate"
)

func getByIdRequestHandler(context *gin.Context) {
	id := context.Param("id")

	validRequest, validationError := isValidGetByIdRequest(id)

	if !validRequest {
		respond.BadRequest(context, validationError)
	}

	recipe, found := GetFromStoreById(id)

	if found {
		respond.Ok(context, createGetByIdResponsePayload(recipe))
	} else {
		respond.NotFound(context)
	}
}

func isValidGetByIdRequest(id string) (isValid bool, validationMessage string) {
	return validate.UuidString(id)
}

func createGetByIdResponsePayload(recipe Recipe) gin.H {
	responsePayload := gin.H{
		"recipe": recipe,
	}
	return responsePayload
}
