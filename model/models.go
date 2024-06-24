package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	ID   string `json:"_id" bson:"_id"`
	Name string `json:"name"`
}

type IngredientMeta struct {
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type IngredientWithoutID struct {
	Name string `json:"name"`
}

type Recipe struct {
	ID              string               `json:"_id" bson:"_id"`
	Name            string               `json:"name"`
	Description     string               `json:"description,omitempty"`
	Category        Category             `json:"category"`
	Steps           []string             `json:"steps"`
	Ingredients     []primitive.ObjectID `json:"ingredients"`
	IngredientsMeta []IngredientMeta     `json:"ingredients_meta"`
}

type RecipeWithoutID struct {
	Name            string               `json:"name"`
	Description     string               `json:"description,omitempty"`
	Category        Category             `json:"category"`
	Steps           []string             `json:"steps"`
	Ingredients     []primitive.ObjectID `json:"ingredients"`
	IngredientsMeta []IngredientMeta     `json:"ingredients_meta"`
}

type RecipeTestData struct {
	ID              string               `json:"_id"`
	Name            string               `json:"name"`
	Description     string               `json:"description,omitempty"`
	Category        Category             `json:"category"`
	Steps           []string             `json:"steps"`
	Ingredients     []primitive.ObjectID `json:"ingredients"`
	IngredientsMeta []IngredientMeta     `json:"ingredients_meta"`
}

// type UpdateRecipeInput struct {
// 	ID              string               `json:"_id"`
// 	Name            *string              `json:"name,omitempty"`
// 	Description     *string              `json:"description,omitempty"`
// 	Category        *Category            `json:"category,omitempty"`
// 	Steps           []string             `json:"steps,omitempty"`
// 	Ingredients     []primitive.ObjectID `json:"ingredients"`
// 	IngredientsMeta []IngredientMeta     `json:"ingredients_meta"`
// }

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
