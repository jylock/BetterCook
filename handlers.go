package main

import (
	"net/http"
	"strings"
)

func handleIndexGet(w http.ResponseWriter, r *http.Request) {
	tmplData := map[string]string{"username": getUsername(r)}

	renderTemplate(w, "index.html", tmplData)
}

func handleTemplateGet(w http.ResponseWriter, r *http.Request) {
	pageName := strings.ToLower(r.URL.Query().Get("page"))
	tmplData := map[string]string{"username": getUsername(r)}

	renderTemplateFromDir(w, "templates", pageName+".html", tmplData)
}
