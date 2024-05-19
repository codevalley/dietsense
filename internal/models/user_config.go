package models

// UserConfig represents a user's configuration preferences.
type UserConfig struct {
	UserID       string `gorm:"primaryKey"`
	DefaultModel string
	MaxTokens    int
	Temperature  float64
}
