package main

import (
	"fmt"
	"go-crud/movie"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var pl = fmt.Println

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", movie.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", movie.GetMovie).Methods("GET")
	r.HandleFunc("/movies", movie.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", movie.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", movie.DeleteMovie).Methods("DELETE")

	pl("Starting server at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
