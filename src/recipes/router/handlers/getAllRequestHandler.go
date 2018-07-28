package recipes

import (
"github.com/jjmschofield/GoCook/src/recipes/store"
"github.com/gin-gonic/gin"
)

func GetAllRequestHandler(context *gin.Context){

	allRecipes := recipies.GetAllFromStore()

	responsePayload := gin.H{
		"recipes": allRecipes,
	};

	context.JSON(200, responsePayload)
}

