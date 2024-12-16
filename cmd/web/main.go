package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	port int
	dsn  string
}

type application struct {
	cfg    config
	logger *log.Logger
	db     *sql.DB
}

func main() {
	logger := log.New(os.Stdout, "[QUOTES]\t", log.LstdFlags)
	logger.Println("Kicking off...")

	var cfg config
	flag.IntVar(&cfg.port, "port", 5050, "HTTP server port")
	flag.StringVar(&cfg.dsn, "dsn", os.Getenv("QUOTES_DSN"), "DSN for the database in use")
	flag.Parse()

	db, err := open("mysql", cfg.dsn)
	if err != nil {
		logger.Printf("failed to open database connection: %s", err)
		os.Exit(1)
	}
	defer db.Close()

	logger.Println("Established database connection")

	app := &application{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}

	server := &http.Server{
		Addr:     fmt.Sprintf(":%d", cfg.port),
		Handler:  app.routes(),
		ErrorLog: logger,
	}

	logger.Printf("Listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		logger.Printf("failed to serve on %s: %s", server.Addr, err)
		os.Exit(1)
	}
}

func open(driver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
