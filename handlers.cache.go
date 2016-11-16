package main

import (
	"io"
	"log"
	"net/http"

	"github.com/LibertyLocked/cachestore"
)

var genStore = cachestore.NewCacheStore()

func handleSetCachePost(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r)
	if username == "" {
		log.Println("User is not logged in")
		return
	}

	key := r.FormValue("key")
	val := r.FormValue("value")
	genStore.SetCache(username, key, val)
	log.Println("Set cache store", username, key, val)
}

func handleGetCacheGet(w http.ResponseWriter, r *http.Request) {
	username := getUsername(r)
	if username == "" {
		log.Println("User is not logged in")
		return
	}

	key := r.URL.Query().Get("key")
	val := genStore.GetCache(username, key)

	// data, err := json.Marshal(val)
	// if err != nil {
	// 	log.Println("Error marshalling storage")
	// }
	io.WriteString(w, val.(string))
}
