package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/ping/router"
	"github.com/jjmschofield/GoCook/recipes"
)

func main() {
	router := gin.Default()

	ping.AddApiRoutes(router)
	recipes.AddApiRoutes(router)

	router.Run(":8080")
}