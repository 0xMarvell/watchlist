package models

// Show is the model for objects to be stored in the wathclist database.
type Show struct {
	Title       string `json:"name"`
	Type        string `json:"type"`
	ReleaseDate string `json:"release_date"`
	Seen        bool   `json:"seen"`
}
