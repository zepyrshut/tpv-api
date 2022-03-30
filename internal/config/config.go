package config

import (
	"log"
)

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type Config struct {
	Port string
	Env  string
	DB   struct {
		DSN string
	}
}

type Application struct {
	Status       AppStatus
	Config       Config
	InProduction bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
}
