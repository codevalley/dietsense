package models

import (
	"gorm.io/gorm"
)

// NutritionDetail represents the nutritional information of a food item.
type NutritionDetail struct {
	gorm.Model
	Component  string  `json:"component"`
	Value      string  `json:"value"`
	Unit       string  `json:"unit"`
	Confidence float64 `json:"confidence"`
}
