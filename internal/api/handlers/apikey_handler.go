package handlers

import (
	"dietsense/internal/models"
	"dietsense/internal/repositories"
	"dietsense/pkg/logging"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type APIKeyRequest struct {
	Email            string `json:"email" binding:"required"`
	RateLimitPerHour int    `json:"rate_limit_per_hour" binding:"required"`
}

// GenerateAPIKey generates a new API key and saves it to the database.
func GenerateAPIKey(db repositories.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req APIKeyRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		key := uuid.New().String()
		apiKey := &models.APIKey{
			Key:              key,
			Email:            req.Email,
			RateLimitPerHour: req.RateLimitPerHour,
			CreatedAt:        time.Now(),
		}

		if err := db.SaveAPIKey(apiKey); err != nil {
			logging.Log.Error("Failed to save API key: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save API key"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"api_key": key})
	}
}
