package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Comment struct {
	User   string
	Rating int
	Text   string
}

func initComments() {
	// Read comments from comments.json
	data, err := ioutil.ReadFile("./comments/comments.json")
	if err != nil {
		log.Println("Error reading comments.json", err)
		return
	}
	// Unmarshal into map[string][]Comment, where key is the recipeID
	commentMap := map[string][]Comment{}
	err = json.Unmarshal(data, &commentMap)
	if err != nil {
		log.Println("Error unmarshalling comments")
		return
	}

	// Put comments into thread-safe store
	for recipeID, comments := range commentMap {
		serverCache.SetCache("comments", recipeID, comments)
	}
}

func getComments(recipeID string) []Comment {
	read := serverCache.GetCache("comments", recipeID)
	if read == nil {
		return []Comment{}
	}
	return serverCache.GetCache("comments", recipeID).([]Comment)
}

func addComment(recipeID string, comment Comment) {
	// First get comments
	comments := getComments(recipeID)
	// Append the new comment
	comments = append(comments, comment)
	// Save
	serverCache.SetCache("comments", recipeID, comments)
}

func getRating(recipe string) float64 {
	comments := getComments(recipe)
	if len(comments) == 0 {
		return 0
	}

	totalStars := 0
	for _, v := range comments {
		totalStars += v.Rating
	}
	average := float64(totalStars) / float64(len(comments))
	return average
}
