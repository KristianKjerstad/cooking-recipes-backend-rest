package model

type Ingredient struct {
	ID       string  `json:"_id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type NewIngredientInput struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type NewRecipeInput struct {
	Name        string                `json:"name"`
	Description *string               `json:"description,omitempty"`
	Category    Category              `json:"category"`
	Steps       []*string             `json:"steps"`
	Ingredients []*NewIngredientInput `json:"ingredients"`
}

type Recipe struct {
	ID          string        `json:"_id"`
	Name        string        `json:"name"`
	Description *string       `json:"description,omitempty"`
	Category    Category      `json:"category"`
	Steps       []*string     `json:"steps"`
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
	Steps       []*string                `json:"steps,omitempty"`
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
