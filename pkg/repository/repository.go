package repository

import (
	"ozinshe/schemas"

	"github.com/jmoiron/sqlx"
)

type Authorizhation interface {
	CreateUser(user schemas.User) (int, error)
	GetUser(username, password string) (schemas.User, error)
}

type User interface {
	GetAllUsers() ([]schemas.User, error)
	GetUser(userId int) (schemas.User, error)
	UpdateUser(userId int, input schemas.UpdateUserInput) error
	DeleteUser(userId int) error

	GetFavoriteMovies(userId int) ([]schemas.Movie, error)
	AddFavoriteMovie(userId, movieId int) (int, error)
	DeleteFavoriteMovie(userId, movieId int) error

	GetWatchedMovies(userId int) ([]schemas.Movie, error)
	AddWatchedMovie(userId, movieId int) (int, error)
	DeleteWatchedMovie(userId, movieId int) error
}

type Movie interface {
	AddMovie(movie schemas.AddMovieInfo) (int, error)
	GetAll(limit, offset int) ([]schemas.Movie, error)
	GetById(movieId int) (schemas.Movie, error)
	SearchByName(serchInput string, limit, offset int) ([]schemas.Movie, error)
	SearchByGenre(genre string, limit, offset int) ([]schemas.Movie, error)
	UploadMovie(movieId int, link string) error
	UpdateMovie(movieId int, movieInput schemas.UpdateMovieInfo) error
	DeleteMovie(movieId int) error

	GetCount() (int, error)
	GetCountByName(name string) (int, error)
	GetCountByGenre(genre string) (int, error)

	GetScreenshots(movieId int) ([]schemas.Screenshot, error)
	AddScreenshot(movieId int, link string) (int, error)
	DeleteScreenshot(movieId, screenId int) (string, error)
}

type Genre interface {
	GetAllGenre() ([]schemas.Genre, error)
	AddGenre(genre string) (int, error)
	UpdateGenre(genreId int, genre string) error
	DeleteGenre(genreId int) error
}

type Repository struct {
	Authorizhation
	User
	Movie
	Genre
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorizhation: NewAuthPostgres(db),
		User:           NewUserPostgres(db),
		Movie:          NewMoviePostgres(db),
		Genre:          NewGenrePostgres(db),
	}
}
