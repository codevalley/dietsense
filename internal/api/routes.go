package api

import (
	"dietsense/internal/api/handlers"
	"dietsense/internal/repositories"
	"dietsense/internal/services"

	"github.com/gin-gonic/gin"
)

// TODO: Make the service type part of POST body, rather than being Get param
// SetupRoutes configures the API routes
func SetupRoutes(router *gin.Engine, factory *services.ServiceFactory, db repositories.Database) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := router.Group("/api/v1")
	{
		api.POST("/analyze", handlers.AnalyzeFood(factory, db))
	}
}
