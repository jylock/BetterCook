package main

import (
	"net/http"
)

func handleLoginGet(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login.html", nil)
}

func handleLoginPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	//password := r.FormValue("password")
	setLoginSession(username, w, r)

	// Redir to index
	http.Redirect(w, r, "/", http.StatusFound)
}

func handleLogoutGet(w http.ResponseWriter, r *http.Request) {
	if getUsername(r) != "" {
		clearLoginSession(w, r)
	}

	// Redir to index
	http.Redirect(w, r, "/", http.StatusFound)
}
