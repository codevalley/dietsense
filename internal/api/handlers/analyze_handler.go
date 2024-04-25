package handlers

import (
	"dietsense/internal/services"
	"dietsense/pkg/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AnalyzeFood creates a handler function that dynamically chooses the food analysis service based on the request.
func AnalyzeFood(factory *services.ServiceFactory) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract 'service_type' from query parameters, fallback to default if not provided
		serviceType := c.Query("service")
		if serviceType == "" {
			serviceType = factory.DefaultService // Use default service type if not specified
		}

		// Get the service implementation based on the provided or default 'service_type'
		service, err := factory.GetService(serviceType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Extracting image and context from the form-data
		fileHeader, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload"})
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot open uploaded file"})
			return
		}
		defer file.Close()

		context := fmt.Sprintf("Here is some info about the picture:*%s*\n%s", c.PostForm("context"), config.Config.ContextString)

		// Calling the service to process the image
		result, err := service.AnalyzeFood(file, context)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to analyze image", "details": err.Error()})
			return
		}

		// Sending the processed result back as JSON
		c.JSON(http.StatusOK, result)
	}
}
