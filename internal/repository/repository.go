package repository

import (
	"github.com/zepyrshut/tpv/internal/models"
)

type DBRepo interface {
	// Items
	OneItem(id int) (*models.ItemEntity, error)
	AllItems() ([]*models.ItemEntity, error)
	AllEnabledItems() ([]*models.ItemRead, error)

	// ItemTypes
	AllTypes() ([]*models.ItemType, error)
	OneType(id int) (*models.ItemType, error)

	// Lounges
	AllLounges() ([]*models.Lounge, error)

	// Tables
	AllTablesFromSelectedLounge(id int) ([]*models.Table, error)
}
