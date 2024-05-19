package sqlite

import (
	"dietsense/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLiteDB implements the Database interface for SQLite.
type SQLiteDB struct {
	db *gorm.DB
}

// NewSQLiteDB initializes a new SQLite database connection.
func NewSQLiteDB(dsn string) (*SQLiteDB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Migrate the schema
	db.AutoMigrate(&models.UserConfig{}, &models.UsageStats{})
	return &SQLiteDB{db: db}, nil
}

// GetLLMKey retrieves the LLM key for a user.
func (s *SQLiteDB) GetLLMKey(userID string) (string, error) {
	var userConfig models.UserConfig
	if err := s.db.First(&userConfig, "user_id = ?", userID).Error; err != nil {
		return "", err
	}
	return userConfig.DefaultModel, nil
}

// SaveLLMKey saves the LLM key for a user.
func (s *SQLiteDB) SaveLLMKey(userID, key string) error {
	return s.db.Model(&models.UserConfig{}).Where("user_id = ?", userID).Update("default_model", key).Error
}

// GetUserConfig retrieves the configuration preferences for a user.
func (s *SQLiteDB) GetUserConfig(userID string) (*models.UserConfig, error) {
	var config models.UserConfig
	if err := s.db.First(&config, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

// SaveUserConfig saves the configuration preferences for a user.
func (s *SQLiteDB) SaveUserConfig(userID string, config *models.UserConfig) error {
	return s.db.Save(config).Error
}

// GetUserUsageStats retrieves the usage statistics for a user.
func (s *SQLiteDB) GetUserUsageStats(userID string) (*models.UsageStats, error) {
	var stats models.UsageStats
	if err := s.db.First(&stats, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &stats, nil
}

// SaveUserUsageStats saves the usage statistics for a user.
func (s *SQLiteDB) SaveUserUsageStats(userID string, stats *models.UsageStats) error {
	return s.db.Save(stats).Error
}
