package main

import (
	"net/http"

	"github.com/supersingh05/quarantstreamz/config"
)

func (app *application) routes(cfg config.Config) http.Handler {
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/", app.home)

	// file server
	fileServer := http.FileServer(http.Dir(cfg.StaticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
