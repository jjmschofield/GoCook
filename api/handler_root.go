package api

import "github.com/gin-gonic/gin"

func rootHandler(context *gin.Context) {
	context.HTML(200, "index.html", gin.H{
		"title": "Main website",
	})
}
