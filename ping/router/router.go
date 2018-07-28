package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/ping/router/handlers"
)

func AddApiRoutes(router *gin.Engine) *gin.Engine{
	routerGroup := router.Group("ping")

	addMiddleware(routerGroup);
	addRoutes(routerGroup);

	return router;
}

func addMiddleware(group *gin.RouterGroup) *gin.RouterGroup{
	return group;
}

func addRoutes(group *gin.RouterGroup) *gin.RouterGroup{
	group.GET("", ping.GetPingHandler)
	return group;
}



