package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zepyrshut/tpv/internal/config"
	"github.com/zepyrshut/tpv/internal/handlers"
)

var app *config.Application

func NewRoutes(a *config.Application) {
	app = a
}

func Routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", handlers.GetStatusHandler)

	router.HandlerFunc(http.MethodGet, "/movie/:id", handlers.GetOneMovie)
	router.HandlerFunc(http.MethodGet, "/movies", handlers.GetAllMovies)

	router.HandlerFunc(http.MethodGet, "/lounges", handlers.GetAllLounges)
	//router.HandlerFunc(http.MethodGet, "/lounge/:id_salon/tables", handlers.GetTableFromLounge)

	return router

}
