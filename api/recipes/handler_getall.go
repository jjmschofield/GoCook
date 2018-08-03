package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
)

func getAllRequestHandler(context *gin.Context) {

	recipes, err := GetAllFromStore()

	if err != nil{
		respond.InternalError(context,"Couldn't retrieve recipes")
		return
	}

	responsePayload := gin.H{
		"recipes": recipes,
	}

	respond.Ok(context, responsePayload)
}
