package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/recipes/store"
)

func GetByIdRequestHandler(context *gin.Context){
	id := context.Param("id")

	recipe := recipies.GetFromStoreById(id)

	responsePayload := gin.H{
		"recipe": recipe,
	};

	context.JSON(200, responsePayload)
}
