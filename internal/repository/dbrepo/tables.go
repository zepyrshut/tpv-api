package dbrepo

import (
	"context"
	"time"

	"github.com/zepyrshut/tpv-api/internal/models"
)

func (m *mariaDBRepo) AllTablesFromSelectedLounge(id int) ([]models.Table, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT 
				Num_Mesa, Num_Comensales, Disponible, descripcion, izq, top, imagen, width, 
				height, id_tarifa, imagenocupada, id_salon, esbarra 
			  FROM 
			  	mesa 
			  WHERE 
			  	id_salon = ?
	`
	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var table models.Table
		err := rows.Scan(
			&table.NumOfTable,
			&table.NumOfDiners,
			&table.Available,
			&table.Description,
			&table.Left,
			&table.Top,
			&table.Picture,
			&table.Width,
			&table.Height,
			&table.Rate,
			&table.PictureBusy,
			&table.Lounge,
			&table.IsCounter,
		)
		if err != nil {
			return nil, err
		}

		tables = append(tables, table)
	}

	return tables, nil
}
