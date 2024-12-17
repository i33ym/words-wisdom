package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/view", app.view)
	mux.HandleFunc("/form", app.form)
	mux.HandleFunc("/create", app.create)
	return mux
}