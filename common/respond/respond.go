package respond

import "github.com/gin-gonic/gin"

func Ok(context *gin.Context, payload interface{}) {
	context.JSON(200, payload)
}

func BadRequest(context *gin.Context, errorMessage string) {
	payload := ErrorPayload{
		Error: errorMessage,
	}
	context.JSON(400, payload)
}

func NotFound(context *gin.Context) {
	context.Status(404)
}

func InternalError(context *gin.Context, errorMessage string) {
	payload := ErrorPayload{
		Error: errorMessage,
	}
	context.JSON(500, payload)
}
