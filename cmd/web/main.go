package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	flagAddress := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info(fmt.Sprintf("Starting server on %s", *flagAddress), slog.Any("address", *flagAddress))

	err := http.ListenAndServe(*flagAddress, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}
