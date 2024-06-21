package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"rest/database"

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
	var recipes = db.AllRecipes()

	data, err := loadDataAsJSON(recipes)
	if err != nil {
		fmt.Errorf("Failed to load recipe as json")
	}
	w.Write(data)
}

func getRecipeByID(w http.ResponseWriter, r *http.Request) {
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
