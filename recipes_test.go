package main

import (
	"testing"
)

func TestGetAllRecipes(t *testing.T) {
	recipes := getAllRecipes()
	if len(recipes) != 1 {
		t.Errorf("Expected 1 recipe, got %d", len(recipes))
	}

	macaroniCheese := getMacaroniCheeseRecipe()
	if recipes[0].ID != macaroniCheese.ID || recipes[0].Title != macaroniCheese.Title || recipes[0].Ingredients != macaroniCheese.Ingredients || recipes[0].Instructions != macaroniCheese.Instructions {
		t.Errorf("Expected recipe %+v, got %+v", macaroniCheese, recipes[0])
	}

	newRecipe := Recipe{ID: 2, Title: "Test Recipe", Ingredients: "Test Ingredients", Instructions: "Test Instructions"}
	addRecipe(newRecipe)

	recipes = getAllRecipes()
	if len(recipes) != 2 {
		t.Errorf("Expected 2 recipes, got %d", len(recipes))
	}

	if recipes[1].ID != newRecipe.ID || recipes[1].Title != newRecipe.Title || recipes[1].Ingredients != newRecipe.Ingredients || recipes[1].Instructions != newRecipe.Instructions {
		t.Errorf("Expected recipe %+v, got %+v", newRecipe, recipes[1])
	}
}

func TestAddRecipe(t *testing.T) {
	initialRecipes := getAllRecipes()
	newRecipe := Recipe{ID: 3, Title: "Another Test Recipe", Ingredients: "More Test Ingredients", Instructions: "More Test Instructions"}
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

func TestGetMacaroniCheeseRecipe(t *testing.T) {
	expectedRecipe := Recipe{
		ID:          1,
		Title:       "Macaroni Cheese",
		Ingredients: "200g macaroni, 50g butter, 50g plain flour, 600ml milk, 250g grated cheddar cheese, Salt and pepper to taste",
		Instructions: "1. Cook the macaroni according to the package instructions, then drain and set aside. 2. In a saucepan, melt the butter over medium heat. Add the flour and stir continuously for 1-2 minutes to form a roux. 3. Gradually add the milk to the roux, stirring constantly until the sauce thickens and is smooth. 4. Remove the saucepan from the heat and stir in the grated cheddar cheese until melted and well combined. 5. Season the cheese sauce with salt and pepper to taste. 6. Add the cooked macaroni to the cheese sauce and mix well to coat the pasta evenly. 7. Serve the macaroni cheese hot, optionally garnished with extra grated cheese or breadcrumbs.",
	}

	recipe := getMacaroniCheeseRecipe()
	if recipe.ID != expectedRecipe.ID || recipe.Title != expectedRecipe.Title || recipe.Ingredients != expectedRecipe.Ingredients || recipe.Instructions != expectedRecipe.Instructions {
		t.Errorf("Expected recipe %+v, got %+v", expectedRecipe, recipe)
	}
}
