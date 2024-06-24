package model

import (
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	ID       string  `json:"_id" bson:"_id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type IngredientWithoutID struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type Recipe struct {
	ID          string        `json:"_id" bson:"_id"`
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Category    Category      `json:"category"`
	Steps       []string      `json:"steps"`
	Ingredients []*Ingredient `json:"ingredients"`
}

type RecipeWithoutID struct {
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Category    Category      `json:"category"`
	Steps       []string      `json:"steps"`
	Ingredients []*Ingredient `json:"ingredients"`
}

type RecipeTestData struct {
	ID          string        `json:"_id"`
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Category    Category      `json:"category"`
	Steps       []string      `json:"steps"`
	Ingredients []*Ingredient `json:"ingredients"`
}

type UpdateIngredientInput struct {
	ID       string   `json:"_id"`
	Name     *string  `json:"name,omitempty"`
	Quantity *float64 `json:"quantity,omitempty"`
	Unit     *string  `json:"unit,omitempty"`
}

type UpdateRecipeInput struct {
	ID          string                   `json:"_id"`
	Name        *string                  `json:"name,omitempty"`
	Description *string                  `json:"description,omitempty"`
	Category    *Category                `json:"category,omitempty"`
	Steps       []string                 `json:"steps,omitempty"`
	Ingredients []*UpdateIngredientInput `json:"ingredients,omitempty"`
}

type Category string

const (
	CategoryDrink      Category = "DRINK"
	CategoryMainCourse Category = "MAIN_COURSE"
	CategoryDessert    Category = "DESSERT"
	CategoryAppetizer  Category = "APPETIZER"
)

var AllCategory = []Category{
	CategoryDrink,
	CategoryMainCourse,
	CategoryDessert,
	CategoryAppetizer,
}
