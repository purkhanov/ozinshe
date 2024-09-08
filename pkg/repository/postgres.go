package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable            = "users"
	moviesTable           = "movies"
	genresTable           = "genres"
	movieScreenshotsTable = "screenshots"
	moviesGenresAssoc     = "movies_genres_assoc"
	favoriteMovies        = "favorite_movies"
	watchedMoviesTable    = "watched_movies"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func (c *Config) dbUrl() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.DBName, c.Password, c.SSLMode,
	)
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.dbUrl())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
