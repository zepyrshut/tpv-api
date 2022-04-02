package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zepyrshut/tpv/internal/config"
	"github.com/zepyrshut/tpv/internal/handlers"
	midd "github.com/zepyrshut/tpv/internal/middleware"
)

var app *config.Application

func NewRoutes(a *config.Application) {
	app = a
}

func Routes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(midd.NoSurf)
	mux.Use(midd.EnableCORS)

	// Status
	mux.Get("/status", handlers.Repo.GetStatusHandler)

	// Lounges
	mux.Get("/lounges", handlers.Repo.GetAllLounges)

	// Tables
	mux.Get("/tables", handlers.Repo.GetTableFromLounge)

	// Items
	mux.Get("/items/all", handlers.Repo.GetAllItems)
	mux.Get("/item/{id}", handlers.Repo.GetOneItem)
	mux.Get("/items/enabled", handlers.Repo.GetAllEnabledItems)

	// ItemsType
	mux.Get("/categories/all", handlers.Repo.GetAllCategories)
	mux.Get("/category/{id}", handlers.Repo.GetOneCategory)

	return mux

}
