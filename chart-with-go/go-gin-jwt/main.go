package main

import (
	"github.com/gin-gonic/gin"

	"authserver/configs"
	"authserver/routes"

)

func main() {
        router := gin.Default()
		configs.ConnectDB()
        routes.PingRoute(router)
		routes.UserRoute(router)
		

        router.Run("localhost:6000") 
}