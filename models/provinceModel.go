package models

import "gorm.io/gorm"

// Provinces model
type Provinces struct {
	gorm.Model
	Name string
}
