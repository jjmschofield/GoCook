package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/ping/router"
	"github.com/jjmschofield/GoCook/recipes/router"
)

func main() {
	router := gin.Default()

	ping.AddRoutes(router)
	recipes.AddRoutes(router)

	router.Run(":8080")
}
