package main

import (
	"log"
	"net/http"
	"strings"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	name := r.URL.Query().Get("name")
	if name != "" {
		w.Write([]byte("Hello " + strings.ToUpper(name) + "\nWelcome to my snippetBox Experiment"))
	} else {
		w.Write([]byte("Hello from SnippetBox Home"))
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here We should show some specific snippet i think !..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here I think user could create some snippets ?!..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on : 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}

}
