package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/api/ping"
	"github.com/jjmschofield/GoCook/api/recipes"
	)

func Start(port string){
	router := gin.Default()

	router.LoadHTMLGlob("api/public/*")

	router.GET("", rootHandler)
	ping.AddApiRoutes(router)
	recipes.AddApiRoutes(router)

	router.Run(":" + port)
}
