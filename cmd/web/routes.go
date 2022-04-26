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
	// Create CSRFToken. This direct our app to ignore any request that doesn't include CSRFToken
	mux.Use(NoSurf)
	// Save the session data so our server will act in a stateful way
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/president", handlers.Repo.President)
	mux.Get("/royal-sweet", handlers.Repo.RoyalSweet)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/contact", handlers.Repo.Contact)

	// Serve static files
	// fileServer returns a handler
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
