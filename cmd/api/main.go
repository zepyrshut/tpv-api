package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zepyrshut/tpv/driver"
	"github.com/zepyrshut/tpv/internal/models"
)

const version = "0.1.0"

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	logger *log.Logger
	DB     models.DBModel
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8081, "Port to listen")
	flag.StringVar(&cfg.env, "env", "development", "Environment")
	flag.StringVar(&cfg.db.dsn, "dsn", "root:infusorio@tcp(localhost:4306)/sysmehotel?parseTime=true", "Database DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "http: ", log.Ldate|log.Ltime)

	db, err := driver.OpenDB(cfg.db.dsn)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		DB:     models.DBModel{DB: db},
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Println("Starting server on port", cfg.port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
