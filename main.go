package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
