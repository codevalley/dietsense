package api

import "github.com/gin-gonic/gin"

// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
