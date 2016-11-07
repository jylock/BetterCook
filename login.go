package main

import (
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(64))

func getUsername(r *http.Request) string {
	session, err := store.Get(r, "login-session")
	if err != nil {
		return ""
	}
	val := session.Values["username"]
	username, ok := val.(string)
	if !ok {
		return ""
	}
	return username
}

func setLoginSession(username string, w http.ResponseWriter, r *http.Request) {
	log.Println("Setting login session: " + username)
	session, err := store.Get(r, "login-session")
	if err != nil {
	}
	session.Values["username"] = username
	session.Save(r, w)
}

func clearLoginSession(w http.ResponseWriter, r *http.Request) {
	log.Println("Clearing login session: " + getUsername(r))
	session, err := store.Get(r, "login-session")
	if err != nil {
	}
	session.Values["username"] = ""
	session.Save(r, w)
}
