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
	return &SQLiteDB{db: db}, nil
}

// GetNutritionDetailByID retrieves a nutritional detail by its ID.
func (s *SQLiteDB) GetNutritionDetailByID(id string) (*models.NutritionDetail, error) {
	var detail models.NutritionDetail
	if err := s.db.First(&detail, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}

// SaveNutritionDetail saves a nutritional detail to the database.
func (s *SQLiteDB) SaveNutritionDetail(detail *models.NutritionDetail) error {
	return s.db.Save(detail).Error
}

// Additional methods can be implemented as needed...
