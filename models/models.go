package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Movie struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdb_id"`
	Genre  string `json:"genre"`
}

type Watchlist struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	MovieID int    `json:"movie_id"`
	Status  string `json:"status"`
	Rating  int    `json:"rating"`
}
