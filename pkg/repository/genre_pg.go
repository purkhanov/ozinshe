package repository

import (
	"fmt"
	"ozinshe/schemas"

	"github.com/jmoiron/sqlx"
)

type GenrePostgres struct {
	db *sqlx.DB
}

func NewGenrePostgres(db *sqlx.DB) *GenrePostgres {
	return &GenrePostgres{db: db}
}

func (r *GenrePostgres) GetAllGenre() ([]schemas.Genre, error) {
	var genres []schemas.Genre

	query := fmt.Sprintf(
		"SELECT id, genre FROM %s;",
		genresTable,
	)

	err := r.db.Select(&genres, query)
	if err != nil {
		return nil, err
	}

	return genres, nil
}

func (r *GenrePostgres) AddGenre(genre string) (int, error) {
	var id int

	query := fmt.Sprintf(
		"INSERT INTO %s (genre) VALUES ($1) RETURNING;",
		genresTable,
	)

	row := r.db.QueryRow(query, genre)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *GenrePostgres) UpdateGenre(genreId int, genre string) error {
	query := fmt.Sprintf(
		"UPDATE %s SET genre = $1 WHERE id = $2;",
		genresTable,
	)

	_, err := r.db.Exec(query, genre, genreId)
	if err != nil {
		return err
	}

	return nil
}

func (r *GenrePostgres) DeleteGenre(genreId int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;",
		genresTable,
	)

	_, err := r.db.Exec(query, genreId)
	if err != nil {
		return err
	}

	return nil
}
