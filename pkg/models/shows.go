package models

import (
	"github.com/0xMarvell/watchlist/pkg/config"
	"github.com/0xMarvell/watchlist/pkg/utils"
	"gorm.io/gorm"
)

var Db *gorm.DB

// Show is the database model for storing tv shows or movies.
type Show struct {
	Title       string `json:"name"`
	Type        string `json:"type"`
	ReleaseDate string `json:"release_date"`
	Seen        bool   `json:"seen"`
}

func init() {
	config.Connect()
	Db = config.GetDB()
	err := Db.AutoMigrate(&Show{})
	utils.HandleErr("Migration failed:", err)

}
