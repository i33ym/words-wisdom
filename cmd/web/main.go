package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port int
	dsn string
}

type application struct {
	cfg config
	logger *log.Logger
}

func main() {
	logger := log.New(os.Stdout, "[WORDS-WISDOM]\t", log.LstdFlags)
	logger.Println("Kicking off...")

	var cfg config
	flag.IntVar(&cfg.port, "port", 5050, "HTTP server port")
	flag.StringVar(&cfg.dsn, "dsn", os.Getenv("WORDS-WISDOM-DSN"), "DSN for the database in use")
	flag.Parse()

	app := &application{
		cfg: cfg,
		logger: logger,
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		ErrorLog: logger,
	}

	logger.Printf("Listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		logger.Printf("failed to serve on %s: %s", server.Addr, err)
		os.Exit(1)
	}
}