package api

import (
	"dietsense/internal/api/handlers"
	"dietsense/internal/services"

	"github.com/gin-gonic/gin"
)

// TODO: Make the service type part of POST body, rather than being Get param
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
			serviceType := c.Query("service")
			if serviceType == "" {
				serviceType = factory.DefaultService // Use default service type if not specified
			}
			service, err := factory.GetService(serviceType)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			handlers.AnalyzeFood(service)(c)
		})
	}
}
