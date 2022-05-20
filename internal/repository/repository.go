package repository

import (
	"github.com/zepyrshut/tpv-api/internal/models"
)

type DBRepo interface {
	// Items
	OneItem(id int) (models.ItemRead, error)
	AllItems() ([]models.ItemEntity, error)
	AllEnabledItems() ([]models.ItemRead, error)

	// ItemCategories
	AllCategories() ([]models.ItemCategoryRead, error)
	OneCategory(category int) (models.ItemCategoryRead, error)

	// Lounges
	AllLounges() ([]models.Lounge, error)

	// Tables
	AllTablesFromSelectedLounge(id int) ([]models.Table, error)
}
