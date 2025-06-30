package main

import (
	"log"
	"net/http"
	"strings"
)

func home(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		w.Write([]byte("Hello " + strings.ToUpper(name) + "\nWelcome to my snippetBox Experiment"))
	} else {
		w.Write([]byte("Hello from SnippetBox Home"))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on : 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}

}
