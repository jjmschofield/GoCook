package recipes

import (
	"github.com/gin-gonic/gin"
)

func getByIdRequestHandler(context *gin.Context){
	id := context.Param("id")

	recipe, found := GetFromStoreById(id)

	if(found){
		responsePayload := gin.H{
			"recipe": recipe,
		};
		context.JSON(200, responsePayload)
	} else{
		context.Status(404)
	}
}
