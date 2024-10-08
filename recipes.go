package main

type Recipe struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Ingredients string `json:"ingredients"`
	Instructions string `json:"instructions"`
}

var recipes []Recipe

func init() {
	macaroniCheese := getMacaroniCheeseRecipe()
	addRecipe(macaroniCheese)
}

func getAllRecipes() []Recipe {
	return recipes
}

func addRecipe(newRecipe Recipe) {
	recipes = append(recipes, newRecipe)
}

func getMacaroniCheeseRecipe() Recipe {
	return Recipe{
		ID:          1,
		Title:       "Macaroni Cheese",
		Ingredients: "200g macaroni, 50g butter, 50g plain flour, 600ml milk, 250g grated cheddar cheese, Salt and pepper to taste",
		Instructions: "1. Cook the macaroni according to the package instructions, then drain and set aside. 2. In a saucepan, melt the butter over medium heat. Add the flour and stir continuously for 1-2 minutes to form a roux. 3. Gradually add the milk to the roux, stirring constantly until the sauce thickens and is smooth. 4. Remove the saucepan from the heat and stir in the grated cheddar cheese until melted and well combined. 5. Season the cheese sauce with salt and pepper to taste. 6. Add the cooked macaroni to the cheese sauce and mix well to coat the pasta evenly. 7. Serve the macaroni cheese hot, optionally garnished with extra grated cheese or breadcrumbs.",
	}
}
