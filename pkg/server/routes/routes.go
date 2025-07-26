package routes

import (
	"github.com/bandanascripts/tondru/pkg/server/controllers"
	"github.com/bandanascripts/tondru/pkg/service/middleware"
	"github.com/gin-gonic/gin"
)

func RegisteredRoutes(c *gin.Engine) {

	c.POST("/tondru/generatetoken", controllers.GenerateTokenHandler)
	c.GET("/tondru/inspect", middleware.TokenMiddleware(), controllers.InspectTokenHandler)
	c.GET("/tondru/userhistory", controllers.UserHistoryHandler)
}
