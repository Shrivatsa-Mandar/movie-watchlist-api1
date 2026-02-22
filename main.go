package main

import (
	"log"
	"net/http"

	"movie-watchlist-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/movies/search", handlers.SearchMovie).Methods("GET")
	r.HandleFunc("/watchlist", handlers.AddToWatchlist).Methods("POST")
	r.HandleFunc("/watchlist/{user_id}", handlers.GetUserWatchlist).Methods("GET")

	log.Println("Server running at http://localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", r))
}
