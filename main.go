package main

import (
	"log"
	"net/http"

	"movie-watchlist-api/config"
	"movie-watchlist-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/movies/search", handlers.SearchMovie).Methods("GET")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(handlers.AuthMiddleware)
	api.HandleFunc("/watchlist", handlers.AddToWatchlist).Methods("POST")
	api.HandleFunc("/watchlist/{user_id}", handlers.GetUserWatchlist).Methods("GET")
	api.HandleFunc("/recommendations/{user_id}", handlers.GetRecommendations).Methods("GET")

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	log.Println("Server running at http://localhost:9097")
	log.Fatal(http.ListenAndServe(":9097", r))
}
