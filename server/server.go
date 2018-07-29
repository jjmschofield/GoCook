package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/ping"
	"github.com/jjmschofield/GoCook/recipes"
	"github.com/spf13/viper"
)

func Start(){
	router := gin.Default()

	ping.AddApiRoutes(router)
	recipes.AddApiRoutes(router)

	router.Run(":" + viper.GetString("HTTP_PORT"))
}
