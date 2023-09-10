package movie

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies = []Movie{
	{ID: "1", Isbn: "32145", Title: "First", Director: &Director{FirstName: "N", LastName: "N"}},
	{ID: "2", Isbn: "32167", Title: "Second", Director: &Director{FirstName: "N", LastName: "N"}},
	{ID: "3", Isbn: "32198", Title: "Third", Director: &Director{FirstName: "N", LastName: "N"}},
	{ID: "3", Isbn: "98765", Title: "Third", Director: &Director{FirstName: "N", LastName: "N"}},
	{ID: "4", Isbn: "45678", Title: "Fourth", Director: &Director{FirstName: "N", LastName: "N"}},
}

// var movies []Movie

func GetMovies(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(movies)
}

func DeleteMovie(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(rw).Encode(movies)
}

func GetMovie(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(rw).Encode(item)
			return
		}
	}
}

func CreateMovie(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(rw).Encode(movie)
}

func UpdateMovie(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(rw).Encode(movie)
		}
	}
}

func GetMoviesList() []Movie {
	return movies
}
