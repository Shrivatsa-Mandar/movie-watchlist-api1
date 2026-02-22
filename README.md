# Capstone Project 4: Movie Watchlist & Recommendation API

A REST API built with **Go (Golang)** that allows users to manage their movie watchlists and get recommendations. The project includes a simple front-end to interact with the API.

---

## Features

- Create users  
- Search for movies (dummy data for simplicity)  
- Add movies to watchlists  
- View a user's watchlist  
- Simple front-end using HTML, CSS, and JavaScript  

---

## Project Structure


movie-watchlist-api/
├── main.go # Starts server and routes
├── config/
│ └── memory.go # In-memory storage for users, movies, watchlists
├── handlers/
│ └── handlers.go # All API endpoints logic
├── models/
│ └── models.go # User, Movie, Watchlist structs
└── frontend/
├── index.html # Front-end UI
├── style.css # Styling for UI
└── app.js # JS for API calls


---

## Installation & Running

1. **Clone the repository**

```bash
git clone https://github.com/Shrivatsa-Mandar/movie-watchlist-api.git
cd movie-watchlist-api

Initialize Go module & install dependencies

go mod init movie-watchlist-api   # if not done
go get github.com/gorilla/mux

Run the Go server

go run .

Server runs at: http://localhost:9090

API Endpoints
Endpoint	Method	Description
/users	POST	Create a new user. JSON body: { "username": "Alice", "email": "alice@example.com" }
/movies/search	GET	Search a movie by title. Example: /movies/search?title=Inception
/watchlist	POST	Add a movie to user's watchlist. JSON: { "user_id":1, "movie_id":1, "status":"WATCHLIST", "rating":0 }
/watchlist/{user_id}	GET	Get the watchlist of a specific user
Front-end Usage

Open frontend/index.html in your browser (or serve with a local server).

Use the UI to:

Create a user

Search movies

Add movies to watchlist

View user watchlist

Note: The front-end expects the Go server to run on port 9090.

Data Storage

Uses in-memory storage (config/memory.go)

No external database needed

For production, can be easily upgraded to SQLite or PostgreSQL

Dependencies

Go 1.25+

Gorilla Mux

Submission Notes

Backend-only implementation works with simple front-end

API tested with Postman and browser

Server port: 9090