package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AddCorsMiddleware(router *gin.Engine) *gin.Engine {
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{
		"http://localhost:3000",
		"https://go-cook-web.herokuapp.com/",
	}

	router.Use(cors.New(config))

	return router
}
