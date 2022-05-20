package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zepyrshut/tpv-api/internal/config"
	"github.com/zepyrshut/tpv-api/internal/database"
	"github.com/zepyrshut/tpv-api/internal/handlers"
	"github.com/zepyrshut/tpv-api/internal/middleware"
	"github.com/zepyrshut/tpv-api/internal/routes"
	"github.com/zepyrshut/tpv-api/internal/util"
)

// Application properties
const version = "0.2.3-beta.9"
const environment = "development"
const inProduction = false

// Initalize application
var app config.Application

func main() {
	server, err := run()
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	err = server.Run("localhost:" + app.Config.Port)
	if err != nil {
		app.ErrorLog.Println(err)
	}

}

func run() (*gin.Engine, error) {
	// Environment variables
	dsn := util.GoDotEnvVariable("DATA_SOURCE_NAME")
	apiPort := util.GoDotEnvVariable("API_PORT")
	csrfToken := util.GoDotEnvVariable("CSRF_TOKEN")

	// Application flags
	// Port
	flag.StringVar(&app.Config.Port, "port", apiPort, "Port to listen")
	// Version and environment
	flag.StringVar(&app.Status.Version, "version", version, "Version")
	flag.StringVar(&app.Status.Environment, "env", environment, "Environment")
	flag.BoolVar(&app.InProduction, "production", inProduction, "Production")
	flag.StringVar(&app.Config.Session.Secret, "secret", csrfToken, "Secret")
	// Database
	flag.StringVar(&app.Config.DB.DSN, "dsn", dsn, "Database DSN")
	flag.Parse()

	// Logging format
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize database
	db, err := database.ConnectSQL(app.Config.DB.DSN)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	// Initialize handlers and routes
	routes.NewRoutes(&app)
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	middleware.NewMiddleware(&app)
	server := routes.Routes()

	return server, err
}
