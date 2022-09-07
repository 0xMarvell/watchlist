package config

import (
	"fmt"
	"log"

	"github.com/0xMarvell/watchlist/pkg/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	utils.HandleErr("error reading environment variables:", err)
}

func Connect() {
	var err error

	host := viper.GetString("HOST")
	user := viper.GetString("USER")
	password := viper.GetString("PASSWORD")
	dbname := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		user,
		password,
		dbname,
		port)

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Database connection successful!")
}

func GetDB() *gorm.DB {
	return Db
}
