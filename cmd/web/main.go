package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Soheil7799/snippetbox/internal/models"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
	DB     *models.SnippetModel
}

func main() {
	flagAddress := flag.String("addr", ":4000", "HTTP Network Address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MariaDB/MySQL data source name (DSN)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
	}
	defer db.Close()

	app := &application{
		logger: logger,
		DB:     &models.SnippetModel{DB: db},
	}

	logger.Info(fmt.Sprintf("Starting server on %s", *flagAddress), slog.Any("address", *flagAddress))

	err = http.ListenAndServe(*flagAddress, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
