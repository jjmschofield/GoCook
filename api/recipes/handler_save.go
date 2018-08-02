package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
	"github.com/jjmschofield/GoCook/common/validate"
)

type saveRequestBody struct {
	Recipe Recipe `json:"recipe" binding:"required"`
}

func saveRequestHandler(context *gin.Context) {
	var requestBody saveRequestBody

	bindError := context.Bind(&requestBody)

	if bindError != nil {
		respond.BadRequest(context, bindError.Error())
		return
	}

	validRequest, validationMessage := isValidSaveRequest(requestBody.Recipe)

	if !validRequest {
		respond.BadRequest(context, validationMessage)
		return
	}

	savedRecipe := SaveToStore(requestBody.Recipe)
	respond.Ok(context, createSaveResponsePayload(savedRecipe))
}

func isValidSaveRequest(recipe Recipe) (validRequest bool, validationMessage string) {
	return validate.Struct(recipe)
}

func createSaveResponsePayload(recipe Recipe) gin.H {
	responsePayload := gin.H{
		"recipe": recipe,
	}
	return responsePayload
}
