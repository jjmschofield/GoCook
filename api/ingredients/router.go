package ingredients

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/auth"
)

func UseApiRoutes(router *gin.Engine) *gin.Engine {
	routerGroup := router.Group("ingredients")

	addMiddleware(routerGroup)
	addRoutes(routerGroup)

	return router
}

func addMiddleware(group *gin.RouterGroup) *gin.RouterGroup {
	group.Use(auth.AuthenticationMiddleware)
	return group
}

func addRoutes(group *gin.RouterGroup) *gin.RouterGroup {

	group.GET("", getAllRequestHandler)
	group.POST("", saveRequestHandler)

	group.GET("/:id", getByIdRequestHandler)

	return group
}
