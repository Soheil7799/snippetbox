package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	files := []string{
		"./ui/html/pages/home.tmpl",
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
	}
	template_set, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = template_set.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		w.Write([]byte("No id was found\nDefault snippet view"))
		return
	} else {
		// w.Write([]byte(fmt.Sprintf("Here is the snippet with the id: %d", id)))
		fmt.Fprintf(w, "Here is the snippet with the id: %d", id)
	}

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Need POST method to work !"))
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"name":"soheil"}`))
	w.Write([]byte("Here I think user could create some snippets ?!..."))
}
