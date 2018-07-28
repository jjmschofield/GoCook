package ping

import "github.com/gin-gonic/gin"

func pongHandler(context *gin.Context){
	responsePayload := gin.H{
		"message": "pong",
	};

	context.JSON(200, responsePayload)
}