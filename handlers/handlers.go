package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"movie-watchlist-api/config"
	"movie-watchlist-api/models"
	"movie-watchlist-api/services"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := config.DB.Exec("INSERT INTO users (username, email) VALUES (?, ?)", user.Username, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	user.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func SearchMovie(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	// Check cache first
	var m models.Movie
	err := config.DB.QueryRow("SELECT id, tmdb_id, title, year, genre, overview, poster_url, release_date FROM movies_cache WHERE title LIKE ?", "%"+title+"%").
		Scan(&m.ID, &m.TMDBID, &m.Title, &m.Year, &m.Genre, &m.Overview, &m.PosterURL, &m.ReleaseDate)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]models.Movie{m})
		return
	}

	// Not in cache, fetch from TMDB
	movies, err := services.SearchTMDB(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Cache the first result for simplicity
	if len(movies) > 0 {
		m = movies[0]
		res, _ := config.DB.Exec("INSERT OR IGNORE INTO movies_cache (tmdb_id, title, year, genre, overview, poster_url, release_date) VALUES (?, ?, ?, ?, ?, ?, ?)",
			m.TMDBID, m.Title, m.Year, m.Genre, m.Overview, m.PosterURL, m.ReleaseDate)
		id, _ := res.LastInsertId()
		m.ID = int(id)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func AddToWatchlist(w http.ResponseWriter, r *http.Request) {
	var watch models.Watchlist
	if err := json.NewDecoder(r.Body).Decode(&watch); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := config.DB.Exec("INSERT INTO watchlists (user_id, movie_id, status, rating) VALUES (?, ?, ?, ?)",
		watch.UserID, watch.MovieID, watch.Status, watch.Rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	watch.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(watch)
}

func GetUserWatchlist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["user_id"])

	rows, err := config.DB.Query(`
		SELECT w.id, w.user_id, w.movie_id, w.status, w.rating, w.created_at 
		FROM watchlists w 
		WHERE w.user_id = ?`, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var result []models.Watchlist
	for rows.Next() {
		var wl models.Watchlist
		rows.Scan(&wl.ID, &wl.UserID, &wl.MovieID, &wl.Status, &wl.Rating, &wl.CreatedAt)
		result = append(result, wl)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetRecommendations(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["user_id"])

	// Optimized query: Find movies in cache that have a genre present in movies the user rated highly (>=4),
	// but are not already in the user's watchlist.
	query := `
		SELECT DISTINCT m.id, m.tmdb_id, m.title, m.year, m.genre, m.overview, m.poster_url, m.release_date 
		FROM movies_cache m
		WHERE m.genre IN (
			SELECT DISTINCT mc.genre FROM movies_cache mc
			JOIN watchlists w ON mc.id = w.movie_id
			WHERE w.user_id = ? AND w.rating >= 4
		)
		AND m.id NOT IN (
			SELECT movie_id FROM watchlists WHERE user_id = ?
		)
		LIMIT 10`

	rows, err := config.DB.Query(query, userID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var recommended []models.Movie
	for rows.Next() {
		var m models.Movie
		err := rows.Scan(&m.ID, &m.TMDBID, &m.Title, &m.Year, &m.Genre, &m.Overview, &m.PosterURL, &m.ReleaseDate)
		if err != nil {
			continue
		}
		recommended = append(recommended, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommended)
}
