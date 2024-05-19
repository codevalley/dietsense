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

// SaveLLMKey saves the LLM key for a user.
func (p *PostgresDB) SaveLLMKey(userID, key string) error {
	return p.db.Model(&models.UserConfig{}).Where("user_id = ?", userID).Update("default_model", key).Error
}

// GetUserConfig retrieves the configuration preferences for a user.
func (p *PostgresDB) GetUserConfig(userID string) (*models.UserConfig, error) {
	var config models.UserConfig
	if err := p.db.First(&config, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &config, nil
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
