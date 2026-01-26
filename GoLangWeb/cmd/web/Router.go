package main

import (
	"myApp/pkg/config"
	"myApp/pkg/http/handler"
	"net/http"
)

func Router(app *config.AppConfig) http.Handler {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register handlers with patterns
	// This handler become an argument to the middleware
	mux.HandleFunc("GET /", FirstMiddlware(handler.Repo.Home)) // Pattern with HTTP method
	// mux.HandleFunc("GET /post/{id}", postsHandler) // Pattern with path parameter

	return mux
}
