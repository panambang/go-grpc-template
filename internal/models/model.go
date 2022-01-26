package models

type SearchResult struct {
	Search       []Movies
	TotalResults string
}

// Movies ...
type Movies struct {
	ID         string   `json:"imdbID"`
	Title      string   `json:"title,omitempty"`
	Year       string   `json:"Year,omitempty"`
	Rated      string   `json:"Rated,omitempty"`
	Runtime    string   `json:"Runtime,omitempty"`
	Genre      string   `json:"Genre,omitempty"`
	Director   string   `json:"Director,omitempty"`
	Writer     string   `json:"Writer,omitempty"`
	Language   string   `json:"Language,omitempty"`
	Actors     string   `json:"Actors,omitempty"`
	Country    string   `json:"Country,omitempty"`
	Awards     string   `json:"Awards,omitempty"`
	Poster     string   `json:"Poster,omitempty"`
	Ratings    []Rating `json:"Ratings,omitempty"`
	Type       string   `json:"Type,omitempty"`
	ImdbRating string   `json:"imdbRating,omitempty"`
	DVD        string   `json:"updated_at,omitempty"`
	Released   string   `json:"created_at,omitempty"`
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
