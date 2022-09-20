package main

import (
	"log"

	"github.com/0xMarvell/watchlist/pkg/config"
	"github.com/0xMarvell/watchlist/pkg/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.Connect()
	config.RunMigrations()
}

func main() {
	r := gin.Default()
	routes.RegisterUserRoutes(r)
	// routes.RegisterShowRoutes(r)
	log.Fatal(r.Run())
}
