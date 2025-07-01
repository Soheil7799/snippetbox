package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	logErr := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()

	fileserver := http.FileServer((http.Dir("./ui/static")))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	addr := flag.String("addr", ":4000", "HTTP network listening address")

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: logErr,
		Handler:  mux,
	}

	flag.Parse()
	logInfo.Printf("Starting server on %s", *addr)
	// err := http.ListenAndServe(*addr, mux)
	err := srv.ListenAndServe()
	logErr.Fatal(err)

}
