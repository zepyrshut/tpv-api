package config

import (
	"log"

	"github.com/zepyrshut/tpv/internal/models"
)

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type Config struct {
	Port int
	Env  string
	DB   struct {
		DSN string
	}
}

type Application struct {
	Config Config
	Logger *log.Logger
	DB     models.DBModel
}
