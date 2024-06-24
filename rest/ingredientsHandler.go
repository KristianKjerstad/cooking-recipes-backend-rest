package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
		w.WriteHeader(http.StatusNotFound)
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

// AddIngredient godoc
// @Summary Add an ingredient
// @Description Add ingredient
// @ID addingredient
// @Produce json
// Accept json
// @Param  ingredient   body  model.IngredientWithoutID  true  "Ingredient"
// @Success 201 {object} model.Ingredient
// @Router /ingredients [post]
func AddIngredient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body model.IngredientWithoutID

	res := json.NewDecoder(r.Body).Decode(&body)
	if res != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getErrorResponse("Failed to decode ingredient"))
		return
	}
	existingIngredient := db.FindIngredientByName(body.Name)
	if existingIngredient != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write(getErrorResponse("An ingredient with this name already exists"))
		return
	}
	newIngredient := db.SaveIngredient(&body)
	data, err := loadDataAsJSON(newIngredient)
	if err != nil {
		w.Write(getErrorResponse("Failed to load ingredient"))
		fmt.Errorf("Failed to load ingredient as json")
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// GenerateIngredients godoc
// @Summary Generate ingredients
// @Description Generate ingredients
// @ID generateingredients
// @Produce json
// Accept json
// @Success 201 {object} []model.Ingredient
// @Router /ingredients/generate [post]
func GenerateIngredients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ingredientsCreated := make([]*model.Ingredient, 0)
	filename := "data/ingredienttestdata.json"

	fileContent, err := os.Open(filename)

	if err != nil {
		fmt.Errorf("failed to open file")
		w.Write(getErrorResponse("Failed to open testdata file"))
	}

	defer fileContent.Close()

	byteResult, _ := io.ReadAll(fileContent)

	var testData []model.Ingredient
	json.Unmarshal([]byte(byteResult), &testData)

	for _, ingredient := range testData {
		ingredientCopy := model.Ingredient{
			ID:   ingredient.ID,
			Name: ingredient.Name,
		}
		ingredientsCreated = append(ingredientsCreated, &ingredient)
		db.SaveIngredientWithID(&ingredientCopy)
	}

	w.WriteHeader(http.StatusCreated)
	res, _ := loadDataAsJSON(ingredientsCreated)
	w.Write(res)

}
