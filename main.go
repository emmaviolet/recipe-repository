package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/recipes", listRecipesHandler)
	http.HandleFunc("/recipes/add", addRecipeHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	recipes := getAllRecipes()
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to the Go Recipe Tutorial!</h1>")
	fmt.Fprintf(w, "<h2>Recipes:</h2>")
	fmt.Fprintf(w, "<ul>")
	for _, recipe := range recipes {
		fmt.Fprintf(w, "<li>%s</li>", recipe.Title)
	}
	fmt.Fprintf(w, "</ul>")
	fmt.Fprintf(w, `<a href="/recipes">View All Recipes</a>`)
	fmt.Fprintf(w, `<a href="/recipes/add" style="position: absolute; top: 10px; right: 10px;">Add Recipe</a>`)
}
