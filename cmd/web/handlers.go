package main

import (
	"html/template"
	"net/http"
)

type data struct {
	ID  int
	Day string
}

func (app *application) home(response http.ResponseWriter, request *http.Request) {
	ts, err := template.ParseFiles("./ui/html/home.html")
	if err != nil {
		app.logger.Printf("failed to parse the file: %s", err)
		response.Write([]byte("something terrible happened. sorry!\n"))

		return
	}

	data := &data{
		ID:  123,
		Day: "Friday",
	}
	if err := ts.Execute(response, data); err != nil {
		app.logger.Printf("failed to execute the template: %s", err)
		response.Write([]byte("something terrible happened. sorry!\n"))

		return
	}
}
