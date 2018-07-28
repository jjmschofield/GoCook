package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/security/auth"
)

func AddApiRoutes(router *gin.Engine) *gin.Engine{
	routerGroup := router.Group("recipes")

	addMiddleware(routerGroup);
	addRoutes(routerGroup);

	return router;
}

func addMiddleware(group *gin.RouterGroup) *gin.RouterGroup{
	group.Use(auth.IsAuthenticatedMiddleware)
	return group;
}

func addRoutes(group *gin.RouterGroup) *gin.RouterGroup{

	group.GET("", getAllRequestHandler)
	group.POST("", saveRequestHandler)

	group.GET("/:id", getByIdRequestHandler)

	return group;
}



