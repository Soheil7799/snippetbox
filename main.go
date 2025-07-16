package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is home page"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is snippet view page"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is snippet create page"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
