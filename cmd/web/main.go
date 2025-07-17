package main

import (
	"database/sql"
	"flag"
	"fmt"

	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
}

func main() {
	flagAddress := flag.String("addr", ":4000", "HTTP Network Address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MariaDB/MySQL data source name (DSN)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
	}

	db, err := openDB(*dsn)
	if err != nil {
		app.logger.Error(err.Error())
	}
	defer db.Close()

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
