package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/logger"
)

func UseDefaults(router *gin.Engine) *gin.Engine{
	router.Use(logger.ContextLoggerMiddleware("monolith"))

	router.Use(logger.RequestLoggerMiddleware)

	router.Use(gin.Recovery())

	router.Use(CorsMiddleware())

	return router
}

