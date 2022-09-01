package models

import "gorm.io/gorm"

// User is the database model for storing user details.
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string
}
