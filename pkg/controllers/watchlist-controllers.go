package controllers

import (
	"fmt"
	"net/http"
)

// Home displays simple message on the home page.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Home Page</h1>")
}

// AddShow adds a new show to a user's watchlist
func AddShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Add New Show</h1>")
}

// ViewWatchlist retrieves all shows in a user's watchlist.
func ViewWatchlist(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Get All Shows</h1>")
}

// GetShowById retrieves a particular show from a user's watchlist using its Id.
func GetShowById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Get Show By Id</h1>")
}

// GetShowById updates the details of a particular show using its Id.
func UpdateShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Update Show</h1>")
}

// DeleteShow deletes a show from a user's watchlist.
func DeleteShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Delete Show</h1>")
}
