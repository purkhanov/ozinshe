package service

import (
	"ozinshe/pkg/repository"
	"ozinshe/pkg/utils"
	"ozinshe/schemas"
)

type Authorizhation interface {
	CreateUser(user schemas.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (map[string]any, error)
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
	GetAll(paginParams *utils.Pagination) (*utils.PaginationMovieResponse, error)
	GetById(movieId int) (schemas.Movie, error)
	SearchByName(serchInput string, paginParams *utils.Pagination) (*utils.PaginationMovieResponse, error)
	SearchByGenre(genre string, paginParams *utils.Pagination) (*utils.PaginationMovieResponse, error)
	UpdateMovie(movieId int, movieInput schemas.UpdateMovieInfo) error
	DeleteMovie(movieId int) error

	GetScreenshots(movieId int) ([]schemas.Screenshot, error)
	AddScreenshot(movieId int, link string) (int, error)
	DeleteScreenshot(movieId, screenId int) (string, error)

	UploadMovie(movieId int, link string) error
}

type Genre interface {
	GetAllGenre() ([]schemas.Genre, error)
	AddGenre(genre string) (int, error)
	UpdateGenre(genreId int, genre string) error
	DeleteGenre(genreId int) error
}

type Service struct {
	Authorizhation
	User
	Movie
	Genre

}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorizhation: NewAuthService(repos.Authorizhation),
		User:           NewUserService(repos.User),
		Movie:          NewMovieService(repos.Movie),
		Genre:          NewGenreService(repos.Genre),
	}
}
