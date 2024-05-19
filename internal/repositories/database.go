package repositories

import "dietsense/internal/models"

// Database defines the interface for database operations.
type Database interface {
	// Save/Retrieve users' LLM Keys
	GetLLMKey(userID string) (string, error)
	SaveLLMKey(userID, key string) error

	// Save/Retrieve default config preferences
	GetUserConfig(userID string) (*models.UserConfig, error)
	SaveUserConfig(userID string, config *models.UserConfig) error

	// Store/Retrieve aggregate info of users LLM calls
	GetUserUsageStats(userID string) (*models.UsageStats, error)
	SaveUserUsageStats(userID string, stats *models.UsageStats) error
}
