package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"rest/database"
	"rest/model"

	"github.com/go-chi/chi/v5"
)

var db = database.Connect()

func loadDataAsJSON[T any](dataToConvert T) ([]byte, error) {
	data, err := json.Marshal(dataToConvert)
	if err != nil {
		fmt.Println("Failed to marshal:", err)
		return nil, errors.New("Could not load recipe as json")
	}
	return data, nil

}

func getAllRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var recipes = db.AllRecipes()
	if recipes == nil {
		fmt.Errorf("Failed to load recipe as json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(recipes) == 0 {
		emptyList := make([]model.Recipe, 0)
		emptyData, _ := loadDataAsJSON(emptyList)
		w.Write(emptyData)
		return
	}
	data, err := loadDataAsJSON(recipes)
	if err != nil {
		fmt.Errorf("Failed to load recipe as json")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(data)
}

func getRecipeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	var recipes = db.FindRecipeByID(idParam)
	if recipes == nil {
		w.WriteHeader(http.StatusNotFound)
	}
	data, err := loadDataAsJSON(recipes)
	if err != nil {
		fmt.Errorf("Failed to load recipe as json")
	}
	w.Write(data)

}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func AddRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body model.NewRecipeInput

	res := json.NewDecoder(r.Body).Decode(&body)
	if res != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db.SaveRecipe(&body)
	data, err := loadDataAsJSON(body)
	if err != nil {
		fmt.Errorf("Failed to load recipe as json")
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
