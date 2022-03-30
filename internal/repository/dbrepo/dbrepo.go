package dbrepo

import (
	"database/sql"

	"github.com/zepyrshut/tpv/internal/config"
	"github.com/zepyrshut/tpv/internal/repository"
)

type mariaDBRepo struct {
	App *config.Application
	DB  *sql.DB
}

func NewMariaRepo(conn *sql.DB, app *config.Application) repository.DBRepo {
	return &mariaDBRepo{
		App: app,
		DB:  conn,
	}
}
