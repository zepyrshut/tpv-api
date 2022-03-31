package repository

import (
	"github.com/zepyrshut/tpv/internal/models"
)

type DBRepo interface {
	// Items
	OneItem(id int) (*models.Item, error)
	AllItems() ([]*models.Item, error)
	AllEnabledItems() ([]*models.Item, error)

	// ItemTypes
	AllTypes() ([]*models.ItemType, error)
	OneType(id int) (*models.ItemType, error)

	// Lounges
	AllLounges() ([]*models.Lounge, error)

	// Tables
	AllTablesFromSelectedLounge(id int) ([]*models.Table, error)

	// Movies
	AllMovies() ([]*models.Movie, error)
	OneMovie(id int) (*models.Movie, error)
}
