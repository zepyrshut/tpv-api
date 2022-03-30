package handlers

import (
	"github.com/zepyrshut/tpv/internal/config"
	"github.com/zepyrshut/tpv/internal/driver"
	"github.com/zepyrshut/tpv/internal/repository"
	"github.com/zepyrshut/tpv/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.Application
	DB  repository.DBRepo
}

func NewRepo(a *config.Application, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewMariaRepo(db.SQL, a),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
