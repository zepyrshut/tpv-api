package dbrepo

import (
	"context"
	"time"

	"github.com/zepyrshut/tpv/internal/models"
)

func (m *mariaDBRepo) AllItems() ([]*models.ItemEntity, error) {
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

	var items []*models.ItemEntity
	for rows.Next() {
		var item models.ItemEntity
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

func (m *mariaDBRepo) AllEnabledItems() ([]*models.ItemRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT
				it.id_complementog, it.complementog, it.PVP, it.favorito, it.impresora, ty.id_tipo_comg,
				ty.padre, ty.tipo_comg, it.date_mod, it.date_sinc
			  FROM
	  			tipo_comg ty    
			  INNER JOIN
	 			complementog it
			  ON
	  			it.id_tipo_comg = ty.id_tipo_comg
	 		  WHERE 
				it.Venta = 'S' 
			  AND 
	  			it.cafeteria = 'S'
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.ItemRead
	for rows.Next() {
		var item models.ItemRead
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.PublicPrice,
			&item.Fav,
			&item.Printer,
			&item.CategoryId,
			&item.ParentCategoryId,
			&item.CategoryName,
			&item.UpdatedAt,
			&item.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		m.appendParentCategoryName(&item)

		items = append(items, &item)

	}

	return items, nil
}

func (m *mariaDBRepo) OneItem(id int) (*models.ItemRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT
				it.id_complementog, it.complementog, it.PVP, it.favorito, it.impresora, 
				ty.id_tipo_comg, ty.padre, ty.tipo_comg, it.date_mod, it.date_sinc
			  FROM
				tipo_comg ty    
			  INNER JOIN
				complementog it
			  ON
				it.id_tipo_comg = ty.id_tipo_comg
			  WHERE 
				it.id_complementog = ?
	`
	row := m.DB.QueryRowContext(ctx, query, id)

	var item models.ItemRead
	err := row.Scan(
		&item.Id,
		&item.Name,
		&item.PublicPrice,
		&item.Fav,
		&item.Printer,
		&item.CategoryId,
		&item.ParentCategoryId,
		&item.CategoryName,
		&item.UpdatedAt,
		&item.CreatedAt,
	)

	if err = row.Err(); err != nil {
		return nil, err
	}

	m.appendParentCategoryName(&item)

	return &item, nil
}

func (m *mariaDBRepo) appendParentCategoryName(item *models.ItemRead) (*models.ItemRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT 
				child.id_tipo_comg,	child.tipo_comg, parent.id_tipo_comg, parent.tipo_comg
			  FROM
				tipo_comg 
			  AS 
			  	child
			  JOIN 
			    tipo_comg 
			  AS 
			  	parent 
			  ON 
			  	parent.id_tipo_comg = child.padre
			  WHERE 
			  	child.id_tipo_comg = ?
	`
	rows, _ := m.DB.QueryContext(ctx, query, item.CategoryId)
	defer rows.Close()

	parentCategoryName := make(map[string]string)
	for rows.Next() {
		var pcn models.ItemRead
		err := rows.Scan(
			&pcn.Id,
			&pcn.Name,
			&pcn.ParentCategoryId,
			&pcn.ParentCategoryName,
		)
		if err != nil {
			return nil, err
		}

		parentCategoryName[pcn.ParentCategoryId.String] = pcn.ParentCategoryName

	}

	item.ParentCategoryName = parentCategoryName[item.ParentCategoryId.String]

	return item, nil

}
