package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/respond"
	"github.com/jjmschofield/GoCook/validate"
)

func getByIdRequestHandler(context *gin.Context) {
	id := context.Param("id")

	validRequest, validationError := isValidRequest(id)

	if !validRequest {
		respond.BadRequest(context, validationError)
	}

	recipe, found := GetFromStoreById(id)

	if found {
		respond.Ok(context, createResponsePayload(recipe))
	} else {
		respond.NotFound(context)
	}
}

func isValidRequest(id string) (isValid bool, validationMessage string) {
	return validate.UuidString(id)
}

func createResponsePayload(recipe Recipe) gin.H {
	responsePayload := gin.H{
		"recipe": recipe,
	}
	return responsePayload
}
