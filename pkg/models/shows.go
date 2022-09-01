package models

// Show is the database model for storing tv shows or movies.
type Show struct {
	Title       string `json:"name"`
	Type        string `json:"type"`
	ReleaseDate string `json:"release_date"`
	Seen        bool   `json:"seen"`
}
