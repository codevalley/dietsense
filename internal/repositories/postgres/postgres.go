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
	return &PostgresDB{db: db}, nil
}

// GetNutritionDetailByID retrieves a nutritional detail by its ID.
func (p *PostgresDB) GetNutritionDetailByID(id string) (*models.NutritionDetail, error) {
	var detail models.NutritionDetail
	if err := p.db.First(&detail, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &detail, nil
}

// SaveNutritionDetail saves a nutritional detail to the database.
func (p *PostgresDB) SaveNutritionDetail(detail *models.NutritionDetail) error {
	return p.db.Save(detail).Error
}

// Additional methods can be implemented as needed...
