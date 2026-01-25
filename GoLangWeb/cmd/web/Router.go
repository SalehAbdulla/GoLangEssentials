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
	mux.HandleFunc("GET /", handler.Repo.Home) // Pattern with HTTP method
	// mux.HandleFunc("GET /post/{id}", postsHandler) // Pattern with path parameter

	return mux
}
