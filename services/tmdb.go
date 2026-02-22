package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"movie-watchlist-api/models"
)

const TMDB_BASE_URL = "https://api.themoviedb.org/3"

var genreMap = map[int]string{
	28:    "Action",
	12:    "Adventure",
	16:    "Animation",
	35:    "Comedy",
	80:    "Crime",
	99:    "Documentary",
	18:    "Drama",
	10751: "Family",
	14:    "Fantasy",
	36:    "History",
	27:    "Horror",
	10402: "Music",
	9648:  "Mystery",
	10749: "Romance",
	878:   "Science Fiction",
	10770: "TV Movie",
	53:    "Thriller",
	10752: "War",
	37:    "Western",
}

type TMDBResponse struct {
	Results []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		ReleaseDate string `json:"release_date"`
		Overview    string `json:"overview"`
		PosterPath  string `json:"poster_path"`
		GenreIDs    []int  `json:"genre_ids"`
	} `json:"results"`
}

func SearchTMDB(query string) ([]models.Movie, error) {
	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		// Dummy data for testing if no API key is provided
		return []models.Movie{
			{
				TMDBID:      1,
				Title:       query,
				Year:        "2026",
				Genre:       "Action",
				Overview:    "A dummy overview for " + query,
				PosterURL:   "https://via.placeholder.com/500x750?text=" + query,
				ReleaseDate: "2026-01-01",
			},
		}, nil
	}

	url := fmt.Sprintf("%s/search/movie?api_key=%s&query=%s", TMDB_BASE_URL, apiKey, query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tmdbResp TMDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&tmdbResp); err != nil {
		return nil, err
	}

	var movies []models.Movie
	for _, res := range tmdbResp.Results {
		genre := "Unknown"
		if len(res.GenreIDs) > 0 {
			if g, ok := genreMap[res.GenreIDs[0]]; ok {
				genre = g
			}
		}

		year := ""
		if len(res.ReleaseDate) >= 4 {
			year = res.ReleaseDate[:4]
		}

		movies = append(movies, models.Movie{
			TMDBID:      res.ID,
			Title:       res.Title,
			Year:        year,
			Genre:       genre,
			Overview:    res.Overview,
			PosterURL:   "https://image.tmdb.org/t/p/w500" + res.PosterPath,
			ReleaseDate: res.ReleaseDate,
		})
	}

	return movies, nil
}
