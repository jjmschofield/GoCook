package respond

import "github.com/gin-gonic/gin"

func Ok(context *gin.Context, jsonPayload interface{}){
	context.JSON(200, jsonPayload)
}

func BadRequest (context *gin.Context, message string){
	jsonPayload := gin.H{
		"error": message,
	}
	context.JSON(400, jsonPayload)
}

func NotFound(context *gin.Context){
	context.Status(404)
}