package models

// UsageStats represents the usage statistics for a user.
type UsageStats struct {
	UserID       string `gorm:"primaryKey"`
	APICallCount int
	TokenUsage   int
}
