package main

import (
	"log"
	"net/http"
	"os"

	"github.com/supersingh05/quarantstreamz/config"
)

// gnna add templates and sql to this eventually
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	cfg := config.ParseConfig()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(cfg),
	}

	infoLog.Printf("Starting server on %s", cfg.Addr)
	errorLog.Fatal(srv.ListenAndServe())
}
