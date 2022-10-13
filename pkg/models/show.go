package models

import "gorm.io/gorm"

// Show is the database models for shows= objects
type Show struct {
	gorm.Model
	Name      string `json:"name"`
	Category  string `json:"category"` // e.g movie, TV series, anime, etc
	Completed bool   `json:"completed"`
	UserID    uint
}
