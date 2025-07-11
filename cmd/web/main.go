package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Soheil7799/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logErr        *log.Logger
	logInfo       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	logErr := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	addr := flag.String("addr", ":4000", "HTTP network listening address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	db, err := openDB(*dsn)
	if err != nil {
		logErr.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logErr.Fatal(err)
	}

	app := &application{
		logErr:        logErr,
		logInfo:       logInfo,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: logErr,
		Handler:  app.routes(),
	}

	flag.Parse()
	logInfo.Printf("Starting server on %s", *addr)
	// err := http.ListenAndServe(*addr, mux)
	err = srv.ListenAndServe()
	logErr.Fatal(err)

}
