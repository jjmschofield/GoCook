// Package api GoCook API Reference.
//
// # Introduction
// GoCook's aim is to produce a simple way of building meal plans and shopping lists between multiple people.
//
// It all starts with the your repository of recipes, record the ingredients, steps and urls for each of your favorite dishes and kiss your bookmark list goodbye!
//
// Once you have a nice collection of meals, its easy to pick what you want to eat each day. The yield of each recipe lets you know where you will have leftovers that you can allocate to a cheeky free lunch.
//
// From a meal plan, the next logical step is the generation of a basic shopping list for you, no adds or anything like that - just a simple list to take on your weekly shop.
//
// Ultimately the most important part of this project is for the author to learn Golang, so all comments and PR's are welcome over at https://github.com/jjmschofield/GoCook
//
// # Accessing the API
// This API makes use of OAuth2.0 for authentication and authorization.
//
// Access to the version hosted on Heroku is not really for the public (this is a personal project) but if you get involved on GitHub and invite can surely be sent!
//
// # Terms Of Service
//
// There are no TOS at this moment, use at your own risk and we take no responsibility. Note that the api is currently running on http only so everyone can see what you are posting!
//
//     Schemes: http
//     Host: go-cook.herokuapp.com
//     BasePath: /
//     Version: 0.0.1
//     License: None Given https://github.com/jjmschofield/GoCook
//     Contact: Jack Schofield<jack@no-email-shared.none> https://github.com/jjmschofield/GoCook
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/api/ingredients"
	"github.com/jjmschofield/GoCook/api/recipes"
	"github.com/jjmschofield/GoCook/api/swagger"
	"github.com/jjmschofield/GoCook/common/middleware"
)

func Start(port string, profile string) {
	if profile != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	middleware.UseDefaults(router)

	router.LoadHTMLGlob("api/public/*")

	router.GET("", rootHandler)

	swagger.UseSwaggerRoutes(router)
	
	recipes.UseApiRoutes(router)
	ingredients.UseApiRoutes(router)

	router.Run(":" + port)
}
