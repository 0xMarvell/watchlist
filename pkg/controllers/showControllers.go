package controllers

import "gorm.io/gorm"

type Show struct {
	gorm.Model
	Name      string `json:"name"`
	Category  string `json:"category"`
	Completed bool   `json:"completed"`
}
