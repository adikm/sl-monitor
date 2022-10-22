package main

import (
	"net/http"
)

func (app *application) routes() {

	http.HandleFunc("/", app.notFound)
	http.HandleFunc("/stations", app.status)

}
