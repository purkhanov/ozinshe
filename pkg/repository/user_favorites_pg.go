package repository

import (
	"fmt"
	"ozinshe/schemas"
)

func (r *UserPostgres) GetFavoriteMovies(userId int) ([]schemas.Movie, error) {
	var movies []schemas.Movie

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
		JOIN %s f ON f.movie_id = m.id 
		WHERE f.user_id = $1
		GROUP BY m.id;
		`,
		moviesTable,
		moviesGenresAssoc,
		genresTable,
		movieScreenshotsTable,
		favoriteMovies,
	)

	if err := r.db.Select(&movies, query, userId); err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *UserPostgres) AddFavoriteMovie(userId, movieId int) (int, error) {
	var id int
	query := fmt.Sprintf(
		`
		INSERT INTO %s (user_id, movie_id) 
		SELECT $1, $2 
		WHERE EXISTS (SELECT 1 FROM users WHERE id = $1) 
			AND EXISTS (SELECT 1 FROM movies WHERE id = $2) 
		RETURNING id;
		`,
		favoriteMovies,
	)

	row := r.db.QueryRow(query, userId, movieId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) DeleteFavoriteMovie(userId, movieId int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE user_id = $1 AND movie_id = $2",
		favoriteMovies,
	)
	_, err := r.db.Exec(query, userId, movieId)

	return err
}
