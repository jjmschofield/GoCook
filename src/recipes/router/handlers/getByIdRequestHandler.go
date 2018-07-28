package recipes

import (
	"github.com/jjmschofield/GoCook/src/recipes/store"
	"github.com/gin-gonic/gin"
)

func GetByIdRequestHandler(context *gin.Context){
	id := context.Param("id")

	recipe := recipies.GetFromStoreById(id)

	responsePayload := gin.H{
		"recipe": recipe,
	};

	context.JSON(200, responsePayload)
}
