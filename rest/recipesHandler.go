package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

type ErrorResponse map[string]string

func getErrorResponse(errorText string) []byte {
	errorResponse := ErrorResponse{"error": errorText}
	res, _ := loadDataAsJSON(errorResponse)

	return res

}

// AllRecipes godoc
// @Summary Get all recipes
// @Description get all recipes
// @ID allrecipes
// @Produce json
// @Success 200 {object} []model.Recipe
// @Router /recipes [get]
func getAllRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var recipes = db.AllRecipes()
	if recipes == nil {
		fmt.Errorf("Failed to load recipe as json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorResponse("Failed to load recipe"))
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
		w.Write(getErrorResponse("Could not load recipe"))
	}
	w.Write(data)
}

// GetRecipe godoc
// @Summary Get recipe by ID
// @Description get a recipe by ID
// @ID getrecipe
// @Produce json
// @Success 200 {object} model.Recipe
// @Param        id   path      string  true  "Recipe ID"
// @Router /recipes/{id} [get]
func getRecipeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := chi.URLParam(r, "id")
	var recipes = db.FindRecipeByID(idParam)
	if recipes == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(getErrorResponse("Could not find recipe"))
		return
	}
	data, err := loadDataAsJSON(recipes)
	if err != nil {
		fmt.Errorf("Failed to load recipe as json")
		w.Write(getErrorResponse("Failed to load recipe"))
		return
	}
	w.Write(data)

}

// AddRecipe godoc
// @Summary Add a recipe
// @Description get a recipe by ID
// @ID addrecipe
// @Produce json
// Accept json
// @Param  recipe   body  model.Recipe  true  "Recipe"
// @Success 201 {object} model.RecipeWithoutID
// @Router /recipes [post]
func AddRecipe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body model.RecipeWithoutID

	res := json.NewDecoder(r.Body).Decode(&body)
	if res != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getErrorResponse("Failed to decode recipe"))
		return
	}
	db.SaveRecipe(&body)
	data, err := loadDataAsJSON(body)
	if err != nil {
		w.Write(getErrorResponse("Failed to load recipe"))
		fmt.Errorf("Failed to load recipe as json")
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// GenerateRecipes godoc
// @Summary Generate recipe
// @Description Generate recipes
// @ID generaterecipe
// @Produce json
// Accept json
// @Success 201 {object} []model.Recipe
// @Router /recipes/generate [post]
func GenerateRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// numRecipesToCreate := 10
	recipesCreated := make([]*model.Recipe, 0)
	filename := "data/testdata.json"

	fileContent, err := os.Open(filename)

	if err != nil {
		fmt.Errorf("failed to open file")
		w.Write(getErrorResponse("Failed to open testdata file"))
	}

	defer fileContent.Close()

	byteResult, _ := io.ReadAll(fileContent)

	var testData []model.Recipe
	json.Unmarshal([]byte(byteResult), &testData)

	for _, recipe := range testData {
		recipeCopy := model.RecipeWithoutID{
			Name:        recipe.Name,
			Description: recipe.Description,
			Steps:       recipe.Steps,
			Category:    recipe.Category,
			Ingredients: recipe.Ingredients,
		}
		recipesCreated = append(recipesCreated, &recipe)
		db.SaveRecipe(&recipeCopy)
	}

	w.WriteHeader(http.StatusCreated)
	res, _ := loadDataAsJSON(recipesCreated)
	w.Write(res)

}
