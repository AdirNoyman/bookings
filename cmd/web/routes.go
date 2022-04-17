package main

import (
	"github.com/AdirNoyman/bookings/pkg/config"
	"github.com/AdirNoyman/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// Create CSRFToken
	mux.Use(NoSurf)
	// Save the session data so our server will act in a stateful way
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}