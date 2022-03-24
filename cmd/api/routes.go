package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/movies", app.getAllMovies)

	router.HandlerFunc(http.MethodGet, "/lounges", app.getAllLounges)
	router.HandlerFunc(http.MethodGet, "/lounge/:id_salon/tables", app.getTableFromLounge)

	return app.enableCORS(router)

}
