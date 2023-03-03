package routes

import (
	"authserver/controllers"
	"github.com/gin-gonic/gin"
)

func PingRoute(router *gin.Engine)  {
	router.GET("/", controllers.Ping())
}