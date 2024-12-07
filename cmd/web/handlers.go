package main

import (
	"net/http"
	"text/template"
)

func (app *application) home(response http.ResponseWriter, request *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Printf("failed to parse the file: %s", err)
		response.Write([]byte("something terrible happened. sorry!\n"))

		return
	}

	if err := ts.ExecuteTemplate(response, "base", nil); err != nil {
		app.logger.Printf("failed to execute the template: %s", err)
		response.Write([]byte("something terrible happened. sorry!\n"))

		return
	}
}

func (app *application) about(response http.ResponseWriter, request *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/about.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Printf("failed to parse the file: %s", err)
		response.Write([]byte("something terrible happened. sorry!\n"))

		return
	}

	if err := ts.ExecuteTemplate(response, "base", nil); err != nil {
		app.logger.Printf("failed to execute the template: %s", err)
		response.Write([]byte("something terrible happened. sorry!\n"))

		return
	}
}
