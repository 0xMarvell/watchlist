package routes

import (
	"github.com/0xMarvell/watchlist/pkg/controllers"
	"github.com/gorilla/mux"
)

// RegisterRoutes handles the routing for the API endpoints.
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1", controllers.Home).Methods("GET")
	r.HandleFunc("/api/v1/watchlist", controllers.ViewWatchlist).Methods("GET")
	r.HandleFunc("/api/v1/watchlist/add", controllers.AddShow).Methods("POST")
	r.HandleFunc("/api/v1/watchlist/{ShowId}", controllers.GetShowById).Methods("GET")
	r.HandleFunc("/api/v1/watchlist/{ShowId}", controllers.UpdateShow).Methods("PUT")
	r.HandleFunc("/api/v1/watchlist/{ShowId}", controllers.DeleteShow).Methods("DELETE")
}
