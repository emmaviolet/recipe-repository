package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func listRecipesHandler(w http.ResponseWriter, r *http.Request) {
    recipes := getAllRecipes()
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(recipes)
}

func addRecipeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, `
            <h1>Add a New Recipe</h1>
            <form action="/recipes/add" method="post">
                <label for="title">Title:</label><br>
                <input type="text" id="title" name="title"><br>
                <label for="ingredients">Ingredients:</label><br>
                <textarea id="ingredients" name="ingredients"></textarea><br>
                <label for="instructions">Instructions:</label><br>
                <textarea id="instructions" name="instructions"></textarea><br>
                <input type="submit" value="Add Recipe">
            </form>
        `)
        return
    }

    if r.Method == http.MethodPost {
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Unable to parse form", http.StatusBadRequest)
            return
        }

        newRecipe := Recipe{
            Title:        r.FormValue("title"),
            Ingredients:  r.FormValue("ingredients"),
            Instructions: r.FormValue("instructions"),
        }

        addRecipe(newRecipe)
        w.WriteHeader(http.StatusCreated)
    }
}
