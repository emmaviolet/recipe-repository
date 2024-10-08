package main

type Recipe struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	Instructions string `json:"instructions"`
}

var recipes []Recipe

func getAllRecipes() []Recipe {
	return recipes
}

func addRecipe(newRecipe Recipe) {
	recipes = append(recipes, newRecipe)
}
