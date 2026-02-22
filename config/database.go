package config

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./movie_api_v2.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS movies_cache (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			tmdb_id INTEGER UNIQUE,
			title TEXT NOT NULL,
			year TEXT,
			genre TEXT,
			overview TEXT,
			poster_url TEXT,
			release_date TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS watchlists (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			movie_id INTEGER,
			status TEXT,
			rating INTEGER DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(movie_id) REFERENCES movies_cache(id)
		);`,
	}

	for _, q := range queries {
		_, err := DB.Exec(q)
		if err != nil {
			log.Fatalf("Error creating table: %v", err)
		}
	}
}
