package models

import (
	"time"
)

// APIKey represents an API key with associated metadata.
type APIKey struct {
	Key              string `gorm:"primaryKey"`
	Email            string `gorm:"index"`
	RateLimitPerHour int    `gorm:"default:20"`
	CreatedAt        time.Time
}
