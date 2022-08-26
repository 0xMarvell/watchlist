package main

import (
	"log"
	"net/http"

	"github.com/0xMarvell/watchlist/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterRoutes(r)
	log.Println("Server is up and running...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
