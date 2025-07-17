package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	template_set, err := template.ParseFiles("ui/html/pages/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	err = template_set.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "this is the view page for note id: %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is snippet create page"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("here we create a new snippet hopefully"))
}
