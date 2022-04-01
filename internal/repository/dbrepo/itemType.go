package dbrepo

import (
	"context"
	"time"

	"github.com/zepyrshut/tpv/internal/models"
)

func (m *mariaDBRepo) AllTypes() ([]*models.ItemType, error) {
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

	var types []*models.ItemType
	for rows.Next() {
		var typex models.ItemType
		err := rows.Scan(
			&typex.Id,
			&typex.ItemTypeName,
		)
		if err != nil {
			return nil, err
		}

		types = append(types, &typex)

	}

	return types, nil
}

func (m *mariaDBRepo) OneType(id int) (*models.ItemType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT 
				id_tipo_comg, tipo_comg 
			  FROM 
			  	tipo_comg
			  WHERE
			  	id_tipo_comg = ?
	`
	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var typex models.ItemType
	for rows.Next() {
		err := rows.Scan(
			&typex.Id,
			&typex.ItemTypeName,
		)
		if err != nil {
			return nil, err
		}

	}

	return &typex, nil
}
