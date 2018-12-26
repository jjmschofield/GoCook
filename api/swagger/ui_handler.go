package swagger

import "github.com/gin-gonic/gin"

func getSwaggerGui(context *gin.Context) {
	context.HTML(200, "swagger.html", gin.H{
		"title": "GoCook API Reference",
	})
}
