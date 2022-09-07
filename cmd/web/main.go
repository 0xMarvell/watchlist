package main

import (
	"log"
	"net/http"

	"github.com/0xMarvell/watchlist/pkg/config"
	"github.com/0xMarvell/watchlist/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	_ "gorm.io/driver/postgres"
)

func main() {
	config.LoadEnv()

	r := mux.NewRouter()

	routes.RegisterRoutes(r)
	log.Println("Server is up and running...")
	log.Fatal(http.ListenAndServe(viper.GetString("PORT"), r))
}
