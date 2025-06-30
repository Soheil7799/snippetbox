package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	template_set, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = template_set.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		w.Write([]byte("No id was found\nDefault snippet view"))
		return
	} else {
		// w.Write([]byte(fmt.Sprintf("Here is the snippet with the id: %d", id)))
		fmt.Fprintf(w, "Here is the snippet with the id: %d", id)
	}

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Need POST method to work !"))
		http.Error(w, "Need POST metod to work !", http.StatusMethodNotAllowed)
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"name":"soheil"}`))
	w.Write([]byte("Here I think user could create some snippets ?!..."))
}
