package models

// UserKey represents a key-value pair for a user.
type UserKey struct {
	UserID   string `gorm:"primaryKey"`
	KeyName  string `gorm:"primaryKey"`
	KeyValue string
}
