package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := "9898"

	http.Handle("/", getHandlers())
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getHandlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handleIndexGet).Methods("GET")
	router.HandleFunc("/page", handlePageGet).Methods("GET")
	router.HandleFunc("/login", handleLoginGet).Methods("GET")
	router.HandleFunc("/login", handleLoginPost).Methods("POST")
	router.HandleFunc("/logout", handleLogoutGet).Methods("GET") // XXX: Logout is a GET

	return router
}
