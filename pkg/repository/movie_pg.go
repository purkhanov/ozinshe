package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"ozinshe/schemas"
	"strings"

	"github.com/jmoiron/sqlx"
)

type MoviePostgres struct {
	db *sqlx.DB
}

func NewMoviePostgres(db *sqlx.DB) *MoviePostgres {
	return &MoviePostgres{db: db}
}

func (r *MoviePostgres) GetAll(limit, offset int) ([]schemas.Movie, error) {
	movies := make([]schemas.Movie, 2)

	query := fmt.Sprintf(
		`
		SELECT 
			m.id, 
			m.name, 
			m.description, 
			m.director, 
			m.producer, 
			m.runtime, 
			m.year, 
			m.stars, 
			m.series, 
			m.seasons, 
			m.video_url, 
			COALESCE(
				array_agg(
					DISTINCT CASE WHEN g.genre IS NOT NULL THEN g.genre ELSE '' END
				), 
				ARRAY[]::text[]
			) AS genres, 
    		COALESCE(
				array_agg(
					DISTINCT CASE WHEN s.link IS NOT NULL THEN s.link ELSE '' END
				), 
				ARRAY[]::text[]
			) AS screenshots 
		FROM %s m
		LEFT JOIN %s mga ON m.id = mga.movie_id 
		LEFT JOIN %s g ON g.id = mga.genre_id 
		LEFT JOIN %s s ON s.movie_id = m.id 
		GROUP BY m.id 
		LIMIT $1 
		OFFSET $2;
		`,
		moviesTable,
		moviesGenresAssoc,
		genresTable,
		movieScreenshotsTable,
	)

	if err := r.db.Select(&movies, query, limit, offset); err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MoviePostgres) SearchByName(serchInput string, limit, offset int) ([]schemas.Movie, error) {
	movies := make([]schemas.Movie, 2)

	query := fmt.Sprintf(
		`
		SELECT 
			m.id, 
			m.name, 
			m.description, 
			m.director, 
			m.producer, 
			m.runtime, 
			m.year, 
			m.stars, 
			m.series, 
			m.seasons, 
			m.video_url, 
			COALESCE(
				array_agg(
					DISTINCT CASE WHEN g.genre IS NOT NULL THEN g.genre ELSE '' END
				), 
				ARRAY[]::text[]
			) AS genres, 
    		COALESCE(
				array_agg(
					DISTINCT CASE WHEN s.link IS NOT NULL THEN s.link ELSE '' END
				), 
				ARRAY[]::text[]
			) AS screenshots 
		FROM %s m
		LEFT JOIN %s mga ON m.id = mga.movie_id 
		LEFT JOIN %s g ON g.id = mga.genre_id
		LEFT JOIN %s s ON s.movie_id = m.id 
		WHERE m.name ILIKE $1 
		GROUP BY m.id 
		LIMIT $2 
		OFFSET $3;
		`,
		moviesTable,
		moviesGenresAssoc,
		genresTable,
		movieScreenshotsTable,
	)
	searchArg := "%" + serchInput + "%"

	if err := r.db.Select(&movies, query, searchArg, limit, offset); err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MoviePostgres) SearchByGenre(genre string, limit, offset int) ([]schemas.Movie, error) {
	movies := make([]schemas.Movie, 2)

	query := fmt.Sprintf(
		`
		SELECT 
			m.id, 
			m.name, 
			m.description, 
			m.director, 
			m.producer, 
			m.runtime, 
			m.year, 
			m.stars, 
			m.series, 
			m.seasons, 
			m.video_url, 
			(
    		    SELECT array_agg(DISTINCT g2.genre) 
    		    FROM movies_genres_assoc mga2 
    		    JOIN genres g2 ON g2.id = mga2.genre_id 
    		    WHERE mga2.movie_id = m.id
    		) AS genres, 
			COALESCE(
				array_agg(
					DISTINCT CASE WHEN s.link IS NOT NULL THEN s.link ELSE '' END
				), 
				ARRAY[]::text[]
			) AS screenshots 
		FROM %s m
		LEFT JOIN %s mga ON m.id = mga.movie_id 
		LEFT JOIN %s g ON g.id = mga.genre_id
		LEFT JOIN %s s ON s.movie_id = m.id 
		WHERE g.genre = $1 
		GROUP BY m.id 
		LIMIT $2 
		OFFSET $3;
		`,
		moviesTable,
		moviesGenresAssoc,
		genresTable,
		movieScreenshotsTable,
	)

	if err := r.db.Select(&movies, query, genre, limit, offset); err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MoviePostgres) GetById(movieId int) (schemas.Movie, error) {
	var movie schemas.Movie

	query := fmt.Sprintf(
		`
		WITH main_movie AS (
			SELECT 
				m.id, 
				m.name, 
				m.description, 
				m.director, 
				m.producer, 
				m.runtime, 
				m.year, 
				m.stars, 
				m.series, 
				m.seasons, 
				m.video_url, 
				COALESCE(
					array_agg(
						DISTINCT CASE WHEN g.genre IS NOT NULL THEN g.genre ELSE '' END
					), 
					ARRAY[]::text[]
				) AS genres, 
				COALESCE(
					array_agg(
						DISTINCT CASE WHEN s.link IS NOT NULL THEN s.link ELSE '' END
					), 
					ARRAY[]::text[]
				) AS screenshots
			FROM %s m
			LEFT JOIN %s mga ON m.id = mga.movie_id 
			LEFT JOIN %s g ON g.id = mga.genre_id
			LEFT JOIN %s s ON s.movie_id = m.id 
			WHERE m.id = $1 
			GROUP BY m.id
		),
		similar_movies AS (
			SELECT 
				m.id, 
				m.name, 
				m.description, 
				m.director, 
				m.producer, 
				m.runtime, 
				m.year, 
				m.stars, 
				m.series, 
				m.seasons, 
				m.video_url, 
				COALESCE(
					array_agg(
						DISTINCT CASE WHEN g.genre IS NOT NULL THEN g.genre ELSE '' END
					), 
					ARRAY[]::text[]
				) AS genres, 
				COALESCE(
					array_agg(
						DISTINCT CASE WHEN s.link IS NOT NULL THEN s.link ELSE '' END
					), 
					ARRAY[]::text[]
				) AS screenshots
			FROM %s m
			LEFT JOIN %s mga ON m.id = mga.movie_id 
			LEFT JOIN %s g ON g.id = mga.genre_id
			LEFT JOIN %s s ON s.movie_id = m.id 
			WHERE m.id != $1 
			AND m.id IN (
				SELECT m2.id
				FROM %s m2
				LEFT JOIN %s mga2 ON m2.id = mga2.movie_id 
				LEFT JOIN %s g2 ON g2.id = mga2.genre_id
				WHERE g2.genre IN (SELECT UNNEST((SELECT genres FROM main_movie))) 
				GROUP BY m2.id
				ORDER BY COUNT(DISTINCT g2.genre) DESC
				LIMIT 5
			)
			GROUP BY m.id
		)
		SELECT 
			main_movie.*,
			COALESCE(
				(
					SELECT array_to_json(array_agg(row_to_json(similar_movies)))
					FROM similar_movies
				), '[]'
			) AS similar_movies
		FROM main_movie
		LEFT JOIN similar_movies ON TRUE
		GROUP BY 
			main_movie.id, main_movie.name, main_movie.description, main_movie.director, 
			main_movie.producer, main_movie.runtime, main_movie.year, main_movie.stars, 
			main_movie.genres, main_movie.screenshots, main_movie.series, main_movie.seasons, 
			main_movie.video_url;
		`,
		moviesTable,
		moviesGenresAssoc,
		genresTable,
		movieScreenshotsTable,
		moviesTable,
		moviesGenresAssoc,
		genresTable,
		movieScreenshotsTable,
		moviesTable,
		moviesGenresAssoc,
		genresTable,
	)

	row := struct {
		schemas.Movie
		SimilarMoviesJson []byte `db:"similar_movies"`
	}{}

	if err := r.db.Get(&row, query, movieId); err != nil {
		return movie, err
	}

	if err := json.Unmarshal(row.SimilarMoviesJson, &row.Movie.SimilarMovies); err != nil {
		return movie, err
	}

	movie = row.Movie
	return movie, nil
}

func (r *MoviePostgres) AddMovie(movie schemas.AddMovieInfo) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var movieId int
	query := fmt.Sprintf(
		`
		INSERT INTO %s 
			(name, description, director, producer, runtime, year, 
			stars, series, seasons) 
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id;
		`,
		moviesTable,
	)

	row := tx.QueryRow(
		query, movie.Name, movie.Description, movie.Director, movie.Producer,
		movie.Runtime, movie.Year, movie.Stars, movie.Series, movie.Seasons,
	)

	if err := row.Scan(&movieId); err != nil {
		tx.Rollback()
		return 0, err
	}

	genreQuery := fmt.Sprintf(
		"INSERT INTO %s (movie_id, genre_id) VALUES ($1, $2);",
		moviesGenresAssoc,
	)

	for _, genreId := range movie.Genres {
		_, err := tx.Exec(genreQuery, movieId, genreId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return movieId, nil
}

func (r *MoviePostgres) UploadMovie(movieId int, link string) error {
	query := fmt.Sprintf(
		"UPDATE %s SET video_url = $1 WHERE id = $2",
		moviesTable,
	)

	_, err := r.db.Exec(query, link, movieId)
	if err != nil {
		return err
	}

	return nil
}

func (r *MoviePostgres) UpdateMovie(movieId int, movieInput schemas.UpdateMovieInfo) error {
	movieInputMap := movieInput.ToMap()

	if len(movieInputMap) == 0 {
		return errors.New("there are not inputs")
	}

	querySets := make([]string, 0)
	args := make([]any, 0)
	argId := 1

	for key, val := range movieInputMap {
		querySets = append(querySets, fmt.Sprintf("%s = $%d", key, argId))
		args = append(args, val)
		argId++
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = %d;",
		moviesTable, strings.Join(querySets, ", "), movieId,
	)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *MoviePostgres) DeleteMovie(movieId int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1;",
		moviesTable,
	)

	_, err := r.db.Exec(query, movieId)

	return err
}

func (r *MoviePostgres) GetCount() (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s", moviesTable)

	if err := r.db.Get(&count, query); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *MoviePostgres) GetCountByName(name string) (int, error) {
	var count int
	query := fmt.Sprintf(
		"SELECT COUNT(id) FROM %s WHERE name ILIKE $1",
		moviesTable,
	)

	searchTerm := "%" + name + "%"
	if err := r.db.Get(&count, query, searchTerm); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *MoviePostgres) GetCountByGenre(genre string) (int, error) {
	var count int
	query := fmt.Sprintf(
		`
		SELECT COUNT(m.id) 
		FROM %s m 
		JOIN %s mga ON mga.movie_id = m.id 
		JOIN %s g ON g.id = mga.genre_id 
		WHERE g.genre = $1
		`,
		moviesTable,
		moviesGenresAssoc,
		genresTable,
	)

	if err := r.db.Get(&count, query, genre); err != nil {
		return 0, err
	}
	return count, nil
}
