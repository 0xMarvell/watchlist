package models

import "gorm.io/gorm"

// SHow is the database models for shows= objects
type Show struct {
	gorm.Model
	Name       string `json:"name"`
	Category   string `json:"category"`
	StarRating uint   `json:"star_rating"`
	Completed  bool   `json:"completed"`
	UserID     uint
}
