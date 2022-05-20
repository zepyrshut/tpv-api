package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zepyrshut/tpv-api/internal/config"
	"github.com/zepyrshut/tpv-api/internal/handlers"
	"github.com/zepyrshut/tpv-api/internal/middleware"
)

var app *config.Application

func NewRoutes(a *config.Application) {
	app = a
}

func Routes() *gin.Engine {

	router := gin.Default()

	// CORS and CSRF protection
	router.Use(cors.Default())
	router.Use(middleware.Sessions("session"))
	router.Use(middleware.CORSMiddleware())

	// Status
	router.GET("/status", handlers.Repo.GetStatusHandler)

	// Lounges
	router.GET("/lounges", handlers.Repo.GetAllLounges)

	// Tables
	router.GET("/tables/:id", handlers.Repo.GetTableFromLounge)

	// Items
	router.GET("/items/all", handlers.Repo.GetAllItems)
	router.GET("/item/:id", handlers.Repo.GetOneItem)
	router.GET("/items/enabled", handlers.Repo.GetAllEnabledItems)

	// ItemsType
	router.GET("/categories/all", handlers.Repo.GetAllCategories)
	router.GET("/category/:id", handlers.Repo.GetOneCategory)

	return router

}
