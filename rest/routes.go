package rest

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/swagger/*", httpSwagger.WrapHandler)
	router.Route("/recipes", a.loadRecipeRoutes)
	router.Route("/ingredients", a.loadIngredientRoutes)

	a.Router = router
}

func (a *App) loadRecipeRoutes(router chi.Router) {
	router.Get("/", getAllRecipes)
	router.Get("/{id}", getRecipeByID)
	router.Post("/", AddRecipe)
	router.Post("/generate", GenerateRecipes)
}

func (a *App) loadIngredientRoutes(router chi.Router) {
	router.Get("/", getAllIngredients)
	router.Get("/{name}", getIngredientByName)
}
