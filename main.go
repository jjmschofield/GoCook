package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/recipes"
	"github.com/jjmschofield/GoCook/ping"
)

func main() {
	router := gin.Default()

	ping.AddApiRoutes(router)
	recipes.AddApiRoutes(router)

	router.Run(":8080")
}
