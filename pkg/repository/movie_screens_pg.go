package repository

import (
	"fmt"
	"ozinshe/schemas"
)

func (r *MoviePostgres) GetScreenshots(movieId int) ([]schemas.Screenshot, error) {
	var screens []schemas.Screenshot

	query := fmt.Sprintf(
		"SELECT s.id, s.link, s.movie_id FROM %s s WHERE s.movie_id = $1",
		movieScreenshotsTable,
	)

	if err := r.db.Select(&screens, query, movieId); err != nil {
		return nil, err
	}

	return screens, nil
}

func (r *MoviePostgres) AddScreenshot(movieId int, link string) (int, error) {
	var id int

	query := fmt.Sprintf(
		"INSERT INTO %s (link, movie_id) VALUES ($1, $2) RETURNING id",
		movieScreenshotsTable,
	)

	row := r.db.QueryRow(query, link, movieId)
	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (r *MoviePostgres) DeleteScreenshot(movieId, screenId int) (string, error) {
	var link string

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1 AND movie_id = $2 RETURNING link",
		movieScreenshotsTable,
	)

	if err := r.db.QueryRow(query, screenId, movieId).Scan(&link); err != nil {
		return "", err
	}

	return link, nil
}
