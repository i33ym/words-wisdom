package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type Quote struct {
	ID        int
	CreatedAt time.Time
	Quote     string
	Author    string
	Version   int
}

func (app *application) form(response http.ResponseWriter, request *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/create.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Printf("failed to parse the file: %s", err)
		response.Write([]byte("something terrible happened. sorry!"))

		return
	}

	if err := ts.ExecuteTemplate(response, "base", nil); err != nil {
		app.logger.Printf("failed to execute the template: %s", err)
		response.Write([]byte("something terrible happened. sorry!"))

		return
	}
}

func (app *application) create(response http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		response.Write([]byte("bad request!"))
		return
	}

	quote := request.PostForm.Get("quote")
	author := request.PostForm.Get("author")

	query := fmt.Sprintf("insert into quotes (quote, author) values ('%s', '%s');", quote, author)
	if _, err := app.db.Exec(query); err != nil {
		app.logger.Printf("failed to insert a new record: %s", err)
		response.Write([]byte("something terrible happened. sorry!"))

		return
	}

	response.Write([]byte("congratulations!"))
}

func (app *application) home(response http.ResponseWriter, request *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Printf("failed to parse the file: %s", err)
		response.Write([]byte("something terrible happened. sorry!"))

		return
	}

	qs := []Quote{}
	q := Quote{}

	query := "select * from quotes;"

	rows, err := app.db.Query(query)
	if err != nil {
		app.logger.Printf("failed to fetch the quotes: %s", err)
		response.Write([]byte("something terrible happened. sorry!"))

		return
	}

	for rows.Next() {
		if err := rows.Scan(&q.ID, &q.CreatedAt, &q.Quote, &q.Author, &q.Version); err != nil {
			app.logger.Printf("failed to scan the quote: %s", err)
			response.Write([]byte("something terrible happened. sorry!"))

			return
		}

		qs = append(qs, q)
	}

	if err := ts.ExecuteTemplate(response, "base", qs); err != nil {
		app.logger.Printf("failed to execute the template: %s", err)
		response.Write([]byte("something terrible happened. sorry!"))

		return
	}
}

func (app *application) view(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil || id < 1 {
		response.Write([]byte("page not found!"))
		return
	}

	// query := "select * from quotes where id = ?;"
	query := fmt.Sprintf("select * from quotes where id = %d;", id)

	row := app.db.QueryRow(query)

	quote := Quote{}

	if err := row.Scan(&quote.ID, &quote.CreatedAt, &quote.Quote, &quote.Author, &quote.Version); err != nil {
		if err == sql.ErrNoRows {
			response.Write([]byte("page not found!"))
			return
		}

		app.logger.Printf("failed to fetch the quote: %s", err)
		response.Write([]byte("something terrible happened. sorry!"))

		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/view.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Printf("failed to parse the file: %s", err)
		response.Write([]byte("something terrible happened. sorry!"))

		return
	}

	if err := ts.ExecuteTemplate(response, "base", quote); err != nil {
		app.logger.Printf("failed to execute the template: %s", err)
		response.Write([]byte("something terrible happened. sorry!"))

		return
	}
}
