package rest

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", Hello)
	router.Route("/recipes", a.loadRecipeRoutes)

	a.Router = router
}

func (a *App) loadRecipeRoutes(router chi.Router) {
	router.Get("/", getAllRecipes)
	router.Get("/{id}", getRecipeByID)
}
