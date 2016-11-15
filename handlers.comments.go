package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handleCommentPost(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r)

	if username == "" {
		log.Println("User is not logged in")
		return
	}

	recipe := r.FormValue("recipe")
	comment := r.FormValue("comment")
	rating, err := strconv.Atoi(r.FormValue("rating"))

	if err != nil {
		log.Println("Error converting rating to int", err)
		return
	}

	// TODO: Save comment
	log.Println(recipe, comment, rating)
}

func handleCommentGet(w http.ResponseWriter, r *http.Request) {
	recipe := strings.ToLower(r.URL.Query().Get("recipe"))

	if recipe == "" {
		log.Println("Cannot get comments: Recipe is not specified")
	}

	// TODO: Get comments
	log.Println(recipe)
}
