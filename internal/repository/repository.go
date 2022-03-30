package repository

import (
	"github.com/zepyrshut/tpv/internal/models"
)

type DBRepo interface {
	// Items
	OneItem(id int) (*models.Item, error)
	AllItems() ([]*models.Item, error)

	// Lounges
	AllLounges() ([]*models.Lounge, error)

	// Tables
	AllTablesFromSelectedLounge(id int) ([]*models.Table, error)
}
