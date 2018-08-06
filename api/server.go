package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/api/recipes"
	"github.com/jjmschofield/GoCook/api/swagger"
		)

func Start(port string){
	router := gin.Default()

	router.LoadHTMLGlob("api/public/*")

	router.GET("", rootHandler)

	swagger.AddSwaggerRoutes(router)
	recipes.AddApiRoutes(router)

	router.Run(":" + port)
}
