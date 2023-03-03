package controllers

import (
	"github.com/gin-gonic/gin"
)
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		 c.JSON(200, gin.H{"data": "Hey Bro! It's working fine!"})
}
}