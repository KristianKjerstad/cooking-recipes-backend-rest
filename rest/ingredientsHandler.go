package rest

import (
	"fmt"
	"net/http"
	"rest/model"

	"github.com/go-chi/chi/v5"
)

// AllIngredients godoc
// @Summary Get all ingredients
// @Description get all ingredients
// @ID allingredients
// @Produce json
// @Success 200 {object} []model.Ingredient
// @Router /ingredients [get]
func getAllIngredients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ingredients = db.AllIngredients()
	if ingredients == nil {
		fmt.Errorf("Failed to load ingredient as json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorResponse("Failed to load ingredient"))
		return
	}
	if len(ingredients) == 0 {
		emptyList := make([]model.Ingredient, 0)
		emptyData, _ := loadDataAsJSON(emptyList)
		w.Write(emptyData)
		return
	}
	data, err := loadDataAsJSON(ingredients)
	if err != nil {
		fmt.Errorf("Failed to load ingredient as json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorResponse("Could not load ingredient"))
	}
	w.Write(data)
}

// GetIngredientByName godoc
// @Summary Get ingredient by name
// @Description get ingredient by name
// @ID ingredientbyname
// @Produce json
// @Success 200 {object} model.Ingredient
// @Param        name   path      string  true  "Name"
// @Router /ingredients/{name} [get]
func getIngredientByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	nameParam := chi.URLParam(r, "name")
	var ingredient = db.FindIngredientByName(nameParam)
	if ingredient == nil {
		fmt.Errorf("Failed to get ingredient")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorResponse("Failed to get ingredient"))
		return
	}

	data, err := loadDataAsJSON(ingredient)
	if err != nil {
		fmt.Errorf("Failed to load ingredient as json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorResponse("Could not load ingredient"))
	}
	w.Write(data)
}
