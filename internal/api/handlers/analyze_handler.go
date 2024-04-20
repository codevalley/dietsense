package handlers

import (
	"dietsense/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AnalyzeFood creates a handler function that depends on an ImageAnalysisService
func AnalyzeFood(service services.FoodAnalysisService) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		context := c.PostForm("context")

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
