package database

import (
	"context"
	"rest/model"

	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client           *mongo.Client
	recipeCollection *mongo.Collection
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// derefString returns defaultValue if str is nil, otherwise returns *str
func derefValue[T any](value *T, defaultValue T) T {
	if value != nil {
		return *value
	}
	return defaultValue
}

func Connect() *DB {
	DB_HOST := getEnv("DB_HOST", "localhost")
	DB_PASSWORD := getEnv("DB_PASSWORD", "password")
	DB_USER := getEnv("DB_USER", "admin")
	DB_PORT := getEnv("DB_PORT", "27017")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	recipeCollection := client.Database("data").Collection("recipes")

	return &DB{client: client, recipeCollection: recipeCollection}
}

func (db *DB) SaveRecipe(input *model.RecipeWithoutID) *model.Recipe {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := db.recipeCollection.InsertOne(ctx, input)
	if err != nil {
		log.Print(err)
		return nil
	}

	ingredients := make([]*model.Ingredient, len(input.Ingredients))
	for i, ing := range input.Ingredients {
		ingredients[i] = &model.Ingredient{
			Name:     ing.Name,
			Unit:     ing.Unit,
			Quantity: ing.Quantity,
		}
	}

	return &model.Recipe{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Steps:       input.Steps,
		Ingredients: ingredients,
	}
}

func (db *DB) FindRecipesByCategory(category model.Category) []*model.Recipe {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := db.recipeCollection.Find(ctx, bson.M{"category": category})
	if err != nil {
		return nil
	}
	var recipes []*model.Recipe
	for cur.Next(ctx) {
		var recipe *model.Recipe

		err := cur.Decode(&recipe)
		if err != nil {
			log.Print(err)
			return nil
		}
		recipes = append(recipes, recipe)
	}

	return recipes
}

func (db *DB) FindRecipeByID(ID string) *model.Recipe {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Print(err)
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := db.recipeCollection.FindOne(ctx, bson.M{"_id": ObjectID})
	if res.Err() != nil {
		return nil
	}
	recipe := model.Recipe{}
	res.Decode(&recipe)

	return &recipe
}

func (db *DB) FindRecipeByName(name string) *model.Recipe {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := db.recipeCollection.FindOne(ctx, bson.M{"name": name})
	if res.Err() != nil {
		return nil
	}
	recipe := model.Recipe{}

	res.Decode(&recipe)

	return &recipe
}

func (db *DB) AllRecipes() []*model.Recipe {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := db.recipeCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Print(err)
		return nil
	}
	var recipes []*model.Recipe
	for cur.Next(ctx) {
		var recipe *model.Recipe
		err := cur.Decode(&recipe)
		if err != nil {
			log.Print(err)
			return nil
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}

func (db *DB) UpdateRecipe(newRecipe *model.UpdateRecipeInput) (*model.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	originalRecipe := db.FindRecipeByID(newRecipe.ID)
	if originalRecipe == nil {
		fmt.Errorf("Could not get original recipe")
		return nil, nil
	}

	var ingredients []*model.Ingredient = make([]*model.Ingredient, len(newRecipe.Ingredients))
	for i, ing := range newRecipe.Ingredients {
		ingredient := model.Ingredient{
			ID:       ing.ID,
			Name:     derefValue(ing.Name, originalRecipe.Ingredients[i].Name),
			Quantity: derefValue(ing.Quantity, originalRecipe.Ingredients[i].Quantity),
			Unit:     derefValue(ing.Unit, originalRecipe.Ingredients[i].Unit),
		}
		ingredients[i] = &ingredient
	}

	recipe := model.Recipe{
		ID:          newRecipe.ID,
		Name:        derefValue(newRecipe.Name, originalRecipe.Name),
		Description: derefValue(newRecipe.Description, originalRecipe.Description),
		Category:    derefValue(newRecipe.Category, originalRecipe.Category),
		Steps:       newRecipe.Steps,
		Ingredients: ingredients,
	}
	ObjectID, err := primitive.ObjectIDFromHex(newRecipe.ID)
	res, err := db.recipeCollection.UpdateOne(ctx, bson.M{"_id": ObjectID}, bson.M{"$set": recipe})
	fmt.Println(res)
	if err != nil {
		fmt.Errorf("Failed to update recipe")
		return nil, nil
	}

	return &recipe, nil
}

func (db *DB) DeleteRecipe(ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := db.recipeCollection.DeleteOne(ctx, bson.M{"_id": ID})
	if err != nil {
		return false, nil
	}

	return true, nil
}
