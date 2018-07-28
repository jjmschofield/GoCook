package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/recipesre"
	"github.com/jjmschofield/GoCook/recipesels"
)

func SaveRequestHandler(context *gin.Context){
	var jsonBody struct {
		Recipe recipes.Recipe `json:"recipe" binding:"required"`
	}

	bindError := context.Bind(&jsonBody)

	if(bindError == nil){

		isValid, validationMessage := jsonBody.Recipe.IsValid()

		if(isValid){
			savedRecipe := recipies.SaveToStore(jsonBody.Recipe)

			responsePayload := gin.H{
				"recipe": savedRecipe,
			};

			context.JSON(200, responsePayload)
		} else {
			responsePayload := gin.H{
				"error": validationMessage,
			};

			context.JSON(400, responsePayload)
		}


	}
}