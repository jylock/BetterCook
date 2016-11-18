package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RecipeInfo struct {
	Title        string
	ID           string
	ImgURL       string
	Time         int
	Calories     int
	Budget       int
	Cuisine      string
	MealType     string
	Serves       int
	Ingredients  []string
	Instructions []string
}

func handleRecipeGet(w http.ResponseWriter, r *http.Request) {
	recipe := r.URL.Query().Get("recipe")

	// Read recipe json file from /recipe
	data, err := ioutil.ReadFile("./recipes/" + recipe + ".json")
	if err != nil {
		log.Println("Error reading comments.json", err)
		return
	}

	var recipeInfo RecipeInfo
	json.Unmarshal(data, &recipeInfo)

	tmplData := map[string]interface{}{}
	tmplData["R"] = recipeInfo
	tmplData["rating"] = getRating(recipe)
	tmplData["username"] = getUsername(r)

	renderTemplateFromDir(w, "templates", "recipe_template.html", tmplData)
}

func handleRecipeThumbnailGet(w http.ResponseWriter, r *http.Request) {
	recipe := r.URL.Query().Get("recipe")

	// Read recipe json file from /recipe
	data, err := ioutil.ReadFile("./recipes/" + recipe + ".json")
	if err != nil {
		log.Println("Error reading comments.json", err)
		return
	}

	var recipeInfo RecipeInfo
	json.Unmarshal(data, &recipeInfo)

	tmplData := map[string]interface{}{}
	tmplData["R"] = recipeInfo
	tmplData["rating"] = getRating(recipe)
	tmplData["username"] = getUsername(r)

	renderTemplateFromDir(w, "templates", "thumbnail_template.html", tmplData)
}

func handleRecipeBudgetGet(w http.ResponseWriter, r *http.Request) {
	recipe := r.URL.Query().Get("recipe")

	// Read recipe json file from /recipe
	data, err := ioutil.ReadFile("./recipes/" + recipe + ".json")
	if err != nil {
		log.Println("Error reading comments.json", err)
		return
	}

	var recipeInfo RecipeInfo
	json.Unmarshal(data, &recipeInfo)

	tmplData := map[string]interface{}{}
	tmplData["R"] = recipeInfo
	tmplData["rating"] = getRating(recipe)
	tmplData["username"] = getUsername(r)

	renderTemplateFromDir(w, "templates", "budget_template.html", tmplData)
}

func handleRecipeCommentsGet(w http.ResponseWriter, r *http.Request) {
	recipe := r.URL.Query().Get("recipe")

	comments := getComments(recipe)
	tmplData := map[string]interface{}{}
	tmplData["Comments"] = comments
	tmplData["recipe"] = recipe
	tmplData["username"] = getUsername(r)

	renderTemplateFromDir(w, "templates", "comments_template.html", tmplData)
}

func handleAllRecipesGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	w.Header().Set("Expires", "0")                                         // Proxies
	files, _ := ioutil.ReadDir("./recipes/")

	recipes := []RecipeInfo{}
	for _, v := range files {
		data, _ := ioutil.ReadFile("./recipes/" + v.Name())
		var recipeInfo RecipeInfo
		json.Unmarshal(data, &recipeInfo)
		recipes = append(recipes, recipeInfo)
	}

	recipeListJSON, _ := json.Marshal(recipes)
	w.Write(recipeListJSON)
}
