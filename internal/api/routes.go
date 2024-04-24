package api

import (
	"dietsense/internal/api/handlers"
	"dietsense/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine, factory *services.ServiceFactory) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api := router.Group("/api/v1")
	{
		api.POST("/analyze", func(c *gin.Context) {
			serviceType := c.Query("service_type") // Or extract from POST body or headers
			service, err := factory.GetService(serviceType)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			handlers.AnalyzeFood(service)(c)
		})
	}
}
