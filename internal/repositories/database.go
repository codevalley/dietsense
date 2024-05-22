package repositories

import "dietsense/internal/models"

// Database defines the interface for database operations.
type Database interface {
	// Save/Retrieve users' key-value pairs
	GetUserKey(userID, keyName string) (string, error)
	SaveUserKey(userID, keyName, keyValue string) error

	// Save/Retrieve default config preferences
	GetUserConfig(userID string) (*models.UserConfig, error)
	SaveUserConfig(userID string, config *models.UserConfig) error

	// Store/Retrieve aggregate info of users LLM calls
	GetUserUsageStats(userID string) (*models.UsageStats, error)
	SaveUserUsageStats(userID string, stats *models.UsageStats) error

	// Save/Retrieve API keys
	GetAPIKey(key string) (*models.APIKey, error)
	SaveAPIKey(apiKey *models.APIKey) error
}
