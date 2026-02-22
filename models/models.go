package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"` // Store hashed password
}

type Movie struct {
	ID          int    `json:"id"`
	TMDBID      int    `json:"tmdb_id"`
	Title       string `json:"title"`
	Year        string `json:"year"`
	Genre       string `json:"genre"`
	Overview    string `json:"overview"`
	PosterURL   string `json:"poster_url"`
	ReleaseDate string `json:"release_date"`
}

type Watchlist struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	MovieID   int    `json:"movie_id"`
	Status    string `json:"status"` // e.g., "PLAN_TO_WATCH", "WATCHED"
	Rating    int    `json:"rating"` // 1-5
	CreatedAt string `json:"created_at"`
}
