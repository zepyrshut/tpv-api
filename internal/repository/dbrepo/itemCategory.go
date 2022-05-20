package dbrepo

import (
	"context"
	"time"

	"github.com/zepyrshut/tpv-api/internal/models"
)

func (m *mariaDBRepo) AllCategories() ([]models.ItemCategoryRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT 
				id_tipo_comg, tipo_comg 
			  FROM 
			  	tipo_comg
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.ItemCategoryRead
	for rows.Next() {
		var category models.ItemCategoryRead
		err := rows.Scan(
			&category.Id,
			&category.CategoryName,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)

	}

	return categories, nil
}

func (m *mariaDBRepo) OneCategory(id int) (models.ItemCategoryRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT 
				id_tipo_comg, tipo_comg 
			  FROM 
			  	tipo_comg
			  WHERE
			  	id_tipo_comg = ?
	`
	rows, _ := m.DB.QueryContext(ctx, query, id)
	// TODO check if rows is empty
	defer rows.Close()

	var oneCategory models.ItemCategoryRead
	for rows.Next() {
		err := rows.Scan(
			&oneCategory.Id,
			&oneCategory.CategoryName,
		)
		if err != nil {
			return oneCategory, err
		}

	}

	m.appendItems(oneCategory)

	return oneCategory, nil
}

func (m *mariaDBRepo) appendItems(oneType models.ItemCategoryRead) (models.ItemCategoryRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT
				id_complementog, complementog
			  FROM
			  	complementog
			  WHERE
			  	id_tipo_comg = ?
	`
	rows, _ := m.DB.QueryContext(ctx, query, oneType.Id)
	defer rows.Close()

	itemsAppend := make(map[string]models.ItemRead)
	for rows.Next() {
		var item models.ItemRead
		err := rows.Scan(
			&item.Id,
			&item.Name,
		)
		if err != nil {
			return oneType, err
		}
		itemsAppend[item.Id] = item

	}

	oneType.Items = itemsAppend

	return oneType, nil
}
