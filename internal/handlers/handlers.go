package handlers

import (
	"github.com/zepyrshut/tpv-api/internal/config"
	"github.com/zepyrshut/tpv-api/internal/database"
	"github.com/zepyrshut/tpv-api/internal/repository"
	"github.com/zepyrshut/tpv-api/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.Application
	DB  repository.DBRepo
}

func NewRepo(a *config.Application, db *database.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewMariaRepo(db.SQL, a),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
