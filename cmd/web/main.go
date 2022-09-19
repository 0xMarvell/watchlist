package main

import (
	"log"

	"github.com/0xMarvell/watchlist/pkg/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.Connect()
	config.RunMigrations()
}

func main() {
	r := gin.Default()
	// routes.RegisterUserRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Fatal(r.Run())
}
