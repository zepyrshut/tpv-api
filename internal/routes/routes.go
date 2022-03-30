package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/zepyrshut/tpv/internal/config"
	"github.com/zepyrshut/tpv/internal/handlers"
	"github.com/zepyrshut/tpv/internal/middleware"
)

var app *config.Application

func NewRoutes(a *config.Application) {
	app = a
}

func Routes() http.Handler {

	router := httprouter.New()

	// Status
	router.HandlerFunc(http.MethodGet, "/status", handlers.Repo.GetStatusHandler)

	// Lounges
	router.HandlerFunc(http.MethodGet, "/lounges", handlers.Repo.GetAllLounges)

	// Tables
	router.HandlerFunc(http.MethodGet, "/tables/:id", handlers.Repo.GetTableFromLounge)

	// Items
	router.HandlerFunc(http.MethodGet, "/items", handlers.Repo.GetAllItems)
	router.HandlerFunc(http.MethodGet, "/item/:id", handlers.Repo.GetOneItem)

	return middleware.EnableCORS(router)

}
