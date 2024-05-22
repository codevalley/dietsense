package api

import (
	"dietsense/internal/api/handlers"
	"dietsense/internal/middleware"
	"dietsense/internal/repositories"
	"dietsense/internal/services"
	"dietsense/pkg/config"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine, factory *services.ServiceFactory, db repositories.Database) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	// Define allowed IP addresses
	//allowedIPs := []string{"127.0.0.1", "::1", "your_allowed_ip1", "your_allowed_ip2"}
	allowedIPs := config.Config.AllowedIPs

	api := router.Group("/api/v1")
	{
		api.POST("/analyze", handlers.AnalyzeFood(factory, db))
		api.POST("/generate-api-key", middleware.RestrictToIPs(allowedIPs), handlers.GenerateAPIKey(db))
	}
}
