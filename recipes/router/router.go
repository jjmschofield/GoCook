package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/recipester/handlers"
)

func AddRoutes(router *gin.Engine) *gin.Engine{
	routerGroup := router.Group("recipes")

	addMiddleware(routerGroup);
	addRoutes(routerGroup);

	return router;
}

func addMiddleware(group *gin.RouterGroup) *gin.RouterGroup{
	return group;
}

func addRoutes(group *gin.RouterGroup) *gin.RouterGroup{

	group.GET("", recipes.GetAllRequestHandler)
	group.POST("", recipes.SaveRequestHandler)

	group.GET("/:id", recipes.GetByIdRequestHandler)

	return group;
}



