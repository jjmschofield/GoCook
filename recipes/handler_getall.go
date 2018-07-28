package recipes

import (
	"github.com/gin-gonic/gin"
)

func getAllRequestHandler(context *gin.Context) {

	allRecipes := GetAllFromStore()

	responsePayload := gin.H{
		"recipes": allRecipes,
	};
	context.JSON(200, responsePayload)
}
