package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
	"github.com/jjmschofield/GoCook/common/validate"
)

func getByIdRequestHandler(context *gin.Context) {
	id := context.Param("id")

	validRequest, validationError := isValidGetByIdRequest(id)

	if !validRequest {
		respond.BadRequest(context, validationError)
		return
	}

	recipe, storeErr := GetFromStoreById(id)

	if storeErr != nil {
		respond.NotFound(context)
		return
	}

	respond.Ok(context, createGetByIdResponsePayload(recipe))

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
