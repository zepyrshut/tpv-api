package models

import (
	"context"
	"time"
)

func (m *DBModel) AllLounges() ([]*Lounge, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id_salon, nombre FROM salon
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lounges []*Lounge
	for rows.Next() {
		var lounge Lounge
		err := rows.Scan(
			&lounge.Id,
			&lounge.Name,
		)
		if err != nil {
			return nil, err
		}
		lounges = append(lounges, &lounge)
	}

	return lounges, nil
}
