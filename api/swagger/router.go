package swagger

import (
	"github.com/gin-gonic/gin"
)

func AddSwaggerRoutes(router *gin.Engine) *gin.Engine {
	routerGroup := router.Group("swagger")
	addMiddleware(routerGroup)
	addRoutes(routerGroup)

	return router
}

func addMiddleware(group *gin.RouterGroup) *gin.RouterGroup {
	return group
}

func addRoutes(group *gin.RouterGroup) *gin.RouterGroup {
	group.GET("", getSwaggerGui)
	group.StaticFile("/swagger.json", "docs/swagger/swagger.json")
	group.StaticFile("/swagger.yaml", "docs/swagger/swagger.yaml")

	return group
}