package routes

import (
	"github.com/0xMarvell/watchlist/pkg/controllers"
	"github.com/gorilla/mux"
)

// RegisterRoutes handles the routing for the API endpoints.
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/watchlist", controllers.ViewWatchlist).Methods("GET")
	r.HandleFunc("/watchlist/add", controllers.AddShow).Methods("POST")
	r.HandleFunc("/watchlist/{ShowId}", controllers.GetShowById).Methods("GET")
	r.HandleFunc("/watchlist/{ShowId}", controllers.UpdateShow).Methods("PUT")
	r.HandleFunc("/watchlist/{ShowId}", controllers.DeleteShow).Methods("DELETE")
}
