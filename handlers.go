package main

import (
	"encoding/json"
	"net/http"
)

func listRecipesHandler(w http.ResponseWriter, r *http.Request) {
	recipes := getAllRecipes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func addRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var newRecipe Recipe
	if err := json.NewDecoder(r.Body).Decode(&newRecipe); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	addRecipe(newRecipe)
	w.WriteHeader(http.StatusCreated)
}
