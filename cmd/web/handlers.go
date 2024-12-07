package main

import "net/http"

func (app *application) home(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Welcome Home...\n"))
}