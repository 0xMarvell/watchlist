package routes

import (
	"github.com/0xMarvell/watchlist/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.Home).Methods("GET")
}
