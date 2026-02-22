package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"movie-watchlist-api/config"
	"movie-watchlist-api/models"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	user.ID = config.UserIDCounter
	config.UserIDCounter++

	config.Users = append(config.Users, user)
	json.NewEncoder(w).Encode(user)
}

func AddMovieToCache(movie models.Movie) {
	config.Movies[movie.Title] = movie
}

func SearchMovie(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	if movie, exists := config.Movies[title]; exists {
		json.NewEncoder(w).Encode(movie)
		return
	}

	movie := models.Movie{
		ID:     config.MovieIDCounter,
		Title:  title,
		Year:   "2026",
		ImdbID: "tt1234567",
		Genre:  "Action",
	}

	config.MovieIDCounter++
	AddMovieToCache(movie)

	json.NewEncoder(w).Encode(movie)
}

func AddToWatchlist(w http.ResponseWriter, r *http.Request) {
	var watch models.Watchlist
	json.NewDecoder(r.Body).Decode(&watch)

	watch.ID = config.WatchlistIDCounter
	config.WatchlistIDCounter++

	config.Watchlists = append(config.Watchlists, watch)
	json.NewEncoder(w).Encode(watch)
}

func GetUserWatchlist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["user_id"])

	var result []models.Watchlist
	for _, wlist := range config.Watchlists {
		if wlist.UserID == userID {
			result = append(result, wlist)
		}
	}

	json.NewEncoder(w).Encode(result)
}
