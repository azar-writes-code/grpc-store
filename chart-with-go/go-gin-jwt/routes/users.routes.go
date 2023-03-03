package routes

import (
	"authserver/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())
	router.GET("/users", controllers.GetAllUsers())
}