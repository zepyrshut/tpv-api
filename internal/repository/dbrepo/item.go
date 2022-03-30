package dbrepo

import (
	"context"
	"time"

	"github.com/zepyrshut/tpv/internal/models"
)

func (m *mariaDBRepo) AllItems() ([]*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT 
				id_complementog, complementog, precio 
			  FROM 
			  	complementog
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(
			&item.ItemId,
			&item.Name,
			&item.Price,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, &item)

	}

	return items, nil
}

func (m *mariaDBRepo) OneItem(id int) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT 
				id_complementog, complementog, precio 
			  FROM 
			   	complementog 
			  WHERE 
			   	id_complementog = ?
	`
	row := m.DB.QueryRowContext(ctx, query, id)

	var item models.Item

	err := row.Scan(
		&item.ItemId,
		&item.Name,
		&item.Price,
	)

	if err = row.Err(); err != nil {
		return nil, err
	}

	return &item, nil
}
