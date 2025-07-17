package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"ui/html/pages/home.gohtml",
		"ui/html/base.gohtml",
		"ui/html/partials/nav.gohtml",
	}
	template_set, err := template.ParseFiles(files...)
	if err != nil {
		//app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		//http.Error(w, "internal server error", http.StatusInternalServerError)
		app.serverError(w, r, err)
		return
	}
	err = template_set.ExecuteTemplate(w, "base", nil)
	if err != nil {
		//app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		//http.Error(w, "internal server error", http.StatusInternalServerError)
		app.serverError(w, r, err)
		
		return
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "this is the view page for note id: %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is snippet create page"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("here we create a new snippet hopefully"))
}
