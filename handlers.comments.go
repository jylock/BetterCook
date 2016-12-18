package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func handleCommentPost(w http.ResponseWriter, r *http.Request) {
	recipe := r.URL.Query().Get("recipe")
	username := getUsername(r)

	if username == "" {
		log.Println("User is not logged in")
		return
	}

	comment := r.FormValue("comment")
	rating, err := strconv.Atoi(r.FormValue("rating"))

	if err != nil {
		log.Println("Error converting rating to int", err)
		return
	}

	addComment(recipe, Comment{username, rating, comment})
	log.Println("Saved", recipe, comment, rating)
	io.WriteString(w, "Success")
}

// func handleCommentGet(w http.ResponseWriter, r *http.Request) {
// 	recipe := r.URL.Query().Get("recipe")

// 	if recipe == "" {
// 		log.Println("Cannot get comments: Recipe is not specified")
// 		return
// 	}

// 	comments := getComments(recipe)
// 	// Marshal and return
// 	commentsData, err := json.Marshal(comments)
// 	if err != nil {
// 		log.Println("Error marshalling comments", recipe)
// 		return
// 	}

// 	w.Write(commentsData)
// }
