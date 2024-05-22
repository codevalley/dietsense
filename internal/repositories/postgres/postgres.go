package postgres

import (
	"dietsense/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDB implements the Database interface for PostgreSQL.
type PostgresDB struct {
	db *gorm.DB
}

// NewPostgresDB initializes a new PostgreSQL database connection.
func NewPostgresDB(dsn string) (*PostgresDB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Migrate the schema
	db.AutoMigrate(&models.UserConfig{}, &models.UsageStats{})
	return &PostgresDB{db: db}, nil
}

// GetLLMKey retrieves the LLM key for a user.
func (p *PostgresDB) GetLLMKey(userID string) (string, error) {
	var userConfig models.UserConfig
	if err := p.db.First(&userConfig, "user_id = ?", userID).Error; err != nil {
		return "", err
	}
	return userConfig.DefaultModel, nil
}

// GetUserKey retrieves a key-value pair for a user.
func (p *PostgresDB) GetUserKey(userID, keyName string) (string, error) {
	var userKey models.UserKey
	if err := p.db.First(&userKey, "user_id = ? AND key_name = ?", userID, keyName).Error; err != nil {
		return "", err
	}
	return userKey.KeyValue, nil
}

// GetUserConfig retrieves the configuration preferences for a user.
func (p *PostgresDB) GetUserConfig(userID string) (*models.UserConfig, error) {
	var config models.UserConfig
	if err := p.db.First(&config, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

// SaveUserKey saves a key-value pair for a user.
func (p *PostgresDB) SaveUserKey(userID, keyName, keyValue string) error {
	userKey := models.UserKey{
		UserID:   userID,
		KeyName:  keyName,
		KeyValue: keyValue,
	}
	return p.db.Save(&userKey).Error
}

// SaveUserConfig saves the configuration preferences for a user.
func (p *PostgresDB) SaveUserConfig(userID string, config *models.UserConfig) error {
	return p.db.Save(config).Error
}

// GetUserUsageStats retrieves the usage statistics for a user.
func (p *PostgresDB) GetUserUsageStats(userID string) (*models.UsageStats, error) {
	var stats models.UsageStats
	if err := p.db.First(&stats, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &stats, nil
}

// SaveUserUsageStats saves the usage statistics for a user.
func (p *PostgresDB) SaveUserUsageStats(userID string, stats *models.UsageStats) error {
	return p.db.Save(stats).Error
}

// GetAPIKey retrieves an API key.
func (p *PostgresDB) GetAPIKey(key string) (*models.APIKey, error) {
	var apiKey models.APIKey
	if err := p.db.First(&apiKey, "key = ?", key).Error; err != nil {
		return nil, err
	}
	return &apiKey, nil
}

// SaveAPIKey saves an API key.
func (p *PostgresDB) SaveAPIKey(apiKey *models.APIKey) error {
	return p.db.Save(apiKey).Error
}
