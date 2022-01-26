package sample

import (
	"context"
	"encoding/json"
	"fmt"
	"go-grpc/internal/rpc"
	"net/http"
)

type MovieServiceServer struct {
	rpc.UnimplementedMovieServiceServer
}

type SearchResult struct {
	Search       []Movie
	TotalResults string
}

// Movies ...
type Movie struct {
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

const (
	APIKey = "faf7e5bb"
)

func (s *MovieServiceServer) FetchArticle(ctx context.Context, in *rpc.FetchRequest) (*rpc.ListMovie, error) {
	var client = &http.Client{}
	var movies SearchResult

	searchword := in.GetSearchword()
	page := in.GetPage()
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&page=%s&s=%s", APIKey, page, searchword)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err

	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&movies)
	if err != nil {
		return nil, err

	}
	moviesRpc := make([]*rpc.Movie, 0)
	for _, movie := range movies.Search {
		moviesRpc = append(moviesRpc, &rpc.Movie{
			ImdbID:  movie.ID,
			Title:   movie.Title,
			Year:    movie.Year,
			Rated:   movie.Rated,
			Runtime: movie.Runtime,
			Genre:   movie.Genre,
		})
	}
	fmt.Println(moviesRpc)
	return &rpc.ListMovie{
		Movies: moviesRpc,
	}, nil
}

func (s *MovieServiceServer) GetArticle(ctx context.Context, in *rpc.SingleRequest) (*rpc.Movie, error) {
	var client = &http.Client{}
	var movie Movie

	id := in.GetId()
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s&plot=full", APIKey, id)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err

	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&movie)
	if err != nil {
		return nil, err

	}
	fmt.Println(movie)

	return &rpc.Movie{
		ImdbID:  movie.ID,
		Title:   movie.Title,
		Year:    movie.Year,
		Rated:   movie.Rated,
		Runtime: movie.Runtime,
	}, nil
}
