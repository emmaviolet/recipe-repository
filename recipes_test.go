package main

import (
	"testing"
)

func TestGetAllRecipes(t *testing.T) {
	recipes := getAllRecipes()
	if len(recipes) != 0 {
		t.Errorf("Expected 0 recipes, got %d", len(recipes))
	}

	newRecipe := Recipe{ID: 1, Title: "Test Recipe", Ingredients: "Test Ingredients", Instructions: "Test Instructions"}
	addRecipe(newRecipe)

	recipes = getAllRecipes()
	if len(recipes) != 1 {
		t.Errorf("Expected 1 recipe, got %d", len(recipes))
	}

	if recipes[0].ID != newRecipe.ID || recipes[0].Title != newRecipe.Title || recipes[0].Ingredients != newRecipe.Ingredients || recipes[0].Instructions != newRecipe.Instructions {
		t.Errorf("Expected recipe %+v, got %+v", newRecipe, recipes[0])
	}
}

func TestAddRecipe(t *testing.T) {
	initialRecipes := getAllRecipes()
	newRecipe := Recipe{ID: 2, Title: "Another Test Recipe", Ingredients: "More Test Ingredients", Instructions: "More Test Instructions"}
	addRecipe(newRecipe)

	recipes := getAllRecipes()
	if len(recipes) != len(initialRecipes)+1 {
		t.Errorf("Expected %d recipes, got %d", len(initialRecipes)+1, len(recipes))
	}

	found := false
	for _, recipe := range recipes {
		if recipe.ID == newRecipe.ID && recipe.Title == newRecipe.Title && recipe.Ingredients == newRecipe.Ingredients && recipe.Instructions == newRecipe.Instructions {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("New recipe %+v not found in recipes", newRecipe)
	}
}
