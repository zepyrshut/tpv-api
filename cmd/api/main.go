package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zepyrshut/tpv/internal/config"
	"github.com/zepyrshut/tpv/internal/driver"
	"github.com/zepyrshut/tpv/internal/handlers"
	"github.com/zepyrshut/tpv/internal/middleware"
	"github.com/zepyrshut/tpv/internal/routes"
)

// Environment variables
const version = "0.1.0"
const environment = "development"
const port = "127.0.0.1:8081"
const inProduction = false
const dsn = "root:infusorio@tcp(localhost:4306)/sysmehotel?parseTime=true"

// Initalize variables
var app config.Application
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	flag.StringVar(&app.Config.Port, "port", port, "Port to listen")
	flag.StringVar(&app.Status.Environment, "env", environment, "Environment")
	flag.StringVar(&app.Status.Version, "version", version, "Version")
	flag.BoolVar(&app.InProduction, "production", inProduction, "Production")
	flag.StringVar(&app.Config.DB.DSN, "dsn", dsn, "Database DSN")
	flag.Parse()

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	db, err := driver.ConnectSQL(app.Config.DB.DSN)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.SQL.Close()

	routes.NewRoutes(&app)
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	middleware.NewMiddleware(&app)

	srv := &http.Server{
		Addr:         app.Config.Port,
		Handler:      routes.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	infoLog.Println("Starting server on port", app.Config.Port)

	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Println(err)
	}

}
