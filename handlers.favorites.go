package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleFavoritesReadGet(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r)
	if username == "" {
		log.Println("User is not logged in")
		return
	}

	favIDs := getUserFavorites(username)
	// Json serialize the favIDs array
	favData, err := json.Marshal(favIDs)
	if err != nil {
		log.Println("Error marshalling favorites", favIDs)
		return
	}
	w.Write(favData)
}

func handleFavoritesAddGet(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r)
	recipe := r.URL.Query().Get("recipe")
	if username == "" {
		log.Println("User is not logged in")
		return
	}
	addToFavorites(username, recipe)
}

func handleFavoritesRemoveGet(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r)
	recipe := r.URL.Query().Get("recipe")
	if username == "" {
		log.Println("User is not logged in")
		return
	}
	removeFromFavorites(username, recipe)
}

func handleFavoritesExists(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r)
	recipe := r.URL.Query().Get("recipe")
	if username == "" {
		log.Println("User is not logged in")
		return
	}
	exists := favoriteExists(username, recipe)
	// Json serialize the boolean
	existsData, err := json.Marshal(exists)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(existsData)
}
