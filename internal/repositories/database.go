package repositories

import "dietsense/internal/models"

// Database defines the interface for database operations.
type Database interface {
	// Method to get a nutritional detail by its ID.
	GetNutritionDetailByID(id string) (*models.NutritionDetail, error)
	// Method to save a nutritional detail.
	SaveNutritionDetail(detail *models.NutritionDetail) error
	// Other methods as needed...
}
