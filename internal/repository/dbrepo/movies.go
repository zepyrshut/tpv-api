package dbrepo

import (
	"context"
	"time"

	"github.com/zepyrshut/tpv/internal/models"
)

func (m *mariaDBRepo) OneMovie(id int) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, release_date, runtime, rating, mpaa_rating,
				created_at, updated_at FROM movies WHERE id = ?
	`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie models.Movie

	err := row.Scan(
		&movie.Id,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.Rating,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	m.appendGenres(&movie)

	return &movie, nil
}

func (m *mariaDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, title, description, year, release_date, runtime, rating, mpaa_rating,
	created_at, updated_at FROM movies ORDER by title
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(
			&movie.Id,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.Rating,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		m.appendGenres(&movie)

		movies = append(movies, &movie)

	}

	return movies, nil
}

func (m *mariaDBRepo) appendGenres(movie *models.Movie) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT
				mg.id, mg.movie_id, mg.genre_id, g.genre_name
			FROM
				movies_genres mg
				left join genres g on (g.id = mg.genre_id)
			WHERE
				mg.movie_id = ?
			`
	rows, _ := m.DB.QueryContext(ctx, query, movie.Id)
	defer rows.Close()

	genres := make(map[int]string)
	for rows.Next() {
		var mg models.MovieGenre
		err := rows.Scan(
			&mg.Id,
			&mg.MovieId,
			&mg.GenreId,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.Id] = mg.Genre.GenreName
	}

	movie.MovieGenre = genres

	return movie, nil
}
