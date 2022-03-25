package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zepyrshut/tpv/internal/config"
	"github.com/zepyrshut/tpv/internal/driver"
	"github.com/zepyrshut/tpv/internal/handlers"
	"github.com/zepyrshut/tpv/internal/models"
	"github.com/zepyrshut/tpv/internal/routes"
)

const version = "0.1.0"

func main() {
	app := &config.Application{
		Config: config.Config{
			Port: 8080,
			Env:  "development",
		},
		Logger: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
	}

	st := &config.AppStatus{
		Environment: app.Config.Env,
		Version:     version,
	}

	flag.IntVar(&app.Config.Port, "port", 8081, "Port to listen")
	flag.StringVar(&app.Config.Env, "env", "development", "Environment")
	flag.StringVar(&app.Config.DB.DSN, "dsn", "root:infusorio@tcp(localhost:4306)/sysmehotel?parseTime=true", "Database DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "http: ", log.Ldate|log.Ltime)

	db, err := driver.OpenDB(app.Config.DB.DSN)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app.DB = models.DBModel{DB: db}

	routes.NewRoutes(app)
	handlers.NewStatusHandler(app, st)
	handlers.NewLoungeHandler(app, st)
	handlers.NewTableHandler(app, st)
	handlers.NewMovieHandler(app, st)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      routes.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Println("Starting server on port", app.Config.Port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
