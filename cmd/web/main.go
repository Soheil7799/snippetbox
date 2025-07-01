package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	logErr  *log.Logger
	logInfo *log.Logger
}

func main() {
	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	logErr := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		logErr:  logErr,
		logInfo: logInfo,
	}

	addr := flag.String("addr", ":4000", "HTTP network listening address")

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: logErr,
		Handler:  app.routes(),
	}

	flag.Parse()
	logInfo.Printf("Starting server on %s", *addr)
	// err := http.ListenAndServe(*addr, mux)
	err := srv.ListenAndServe()
	logErr.Fatal(err)

}
