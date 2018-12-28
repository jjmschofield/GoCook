package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{
		"http://localhost:3000",
		"https://go-cook-web.herokuapp.com/",
	}

	config.AllowHeaders = []string{
		"authorization",
	}

	return cors.New(config)
}
