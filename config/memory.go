package config

import "movie-watchlist-api/models"

var Users []models.User
var Movies = make(map[string]models.Movie)
var Watchlists []models.Watchlist

var UserIDCounter = 1
var MovieIDCounter = 1
var WatchlistIDCounter = 1
