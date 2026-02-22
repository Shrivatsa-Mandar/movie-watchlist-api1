# Capstone Project 4: Movie Watchlist & Recommendation API

A REST API built with **Go (Golang)** that allows users to manage their movie watchlists and get recommendations. The project integrates with **TMDB API** and uses **SQLite** for persistence.

---

## Features

- **User Management**: Create and manage users.
- **External Search**: Search for movies using the TMDB API.
- **Caching**: Searched movies are cached in SQLite for performance.
- **Watchlist**: Add movies to watchlists, update status (PLAN_TO_WATCH, WATCHED), and rate them (1-5).
- **Recommendation Engine**: Personalized movie recommendations based on user's high-rated genres.
- **Web UI**: A simple front-end using HTML, CSS, and JavaScript.

---

## Project Structure

```
movie-watchlist-api/
├── main.go               # Entry point and routing
├── config/
│   └── database.go       # SQLite connection and migrations
├── handlers/
│   └── handlers.go       # API endpoint logic
├── models/
│   └── models.go         # Data structures (User, Movie, Watchlist)
├── services/
│   └── tmdb.go           # TMDB API client & genre mapping
└── frontend/
    ├── index.html        # UI interface
    ├── style.css         # UI styles
    └── app.js            # Frontend logic
```

---

## Installation & Running

1. **Clone the repository**
   ```bash
   git clone https://github.com/Shrivatsa-Mandar/movie-watchlist-api.git
   cd movie-watchlist-api
   ```

2. **Setup Dependencies**
   ```bash
   go mod init movie-watchlist-api
   go get github.com/gorilla/mux
   go get github.com/mattn/go-sqlite3
   ```

3. **Set Environment Variables**
   Set your TMDB API key:
   ```bash
   # Windows (PowerShell)
   $env:TMDB_API_KEY = "your_tmdb_api_key"
   ```

4. **Run the Server**
   ```bash
   go run .
   ```
   Server runs at: `http://localhost:9096`

---

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/users` | `POST` | Create a new user. Body: `{"username": "Alice", "email": "alice@email.com"}` |
| `/movies/search` | `GET` | Search TMDB by title. Example: `/movies/search?title=Inception` |
| `/watchlist` | `POST` | Add to watchlist. Body: `{"user_id":1, "movie_id":1, "status":"WATCHLIST", "rating":0}` |
| `/watchlist/{user_id}` | `GET` | Get user's watchlist. |
| `/recommendations/{user_id}` | `GET` | Get personalized movie recommendations. |

---

## Front-end Usage

Open `frontend/index.html` in your browser.
Ensure the Go server is running on port `9096`.

---

## Tech Stack

- **Go 1.25+**
- **SQLite3**
- **Gorilla Mux**
- **TMDB API**
- **Vanilla JS/HTML/CSS**