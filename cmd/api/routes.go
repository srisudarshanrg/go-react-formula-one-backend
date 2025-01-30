package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) Routes() http.Handler {
	mux := chi.NewRouter()

	// middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// routes
	mux.Get("/home", app.Home)
	mux.Get("/drivers", app.Drivers)
	mux.Get("/teams", app.Teams)

	mux.Post("/search", app.search)
	mux.Post("/compare", app.compare)

	return mux
}
