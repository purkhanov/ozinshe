package service

import (
	"ozinshe/pkg/repository"
	"ozinshe/schemas"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]schemas.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUser(userId int) (schemas.User, error) {
	return s.repo.GetUser(userId)
}

func (s *UserService) UpdateUser(userId int, input schemas.UpdateUserInput) error {
	if input.Password != nil {
		*input.Password = generatePasswordHash(*input.Password)
	}
	return s.repo.UpdateUser(userId, input)
}

func (s *UserService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}

func (s *UserService) GetFavoriteMovies(userId int) ([]schemas.Movie, error) {
	return s.repo.GetFavoriteMovies(userId)
}

func (s *UserService) AddFavoriteMovie(userId, movieId int) (int, error) {
	return s.repo.AddFavoriteMovie(userId, movieId)
}

func (s *UserService) DeleteFavoriteMovie(userId, movieId int) error {
	return s.repo.DeleteFavoriteMovie(userId, movieId)
}

func (s *UserService) GetWatchedMovies(userId int) ([]schemas.Movie, error) {
	return s.repo.GetWatchedMovies(userId)
}

func (s *UserService) AddWatchedMovie(userId, movieId int) (int, error) {
	return s.repo.AddWatchedMovie(userId, movieId)
}

func (s *UserService) DeleteWatchedMovie(userId, movieId int) error {
	return s.repo.DeleteWatchedMovie(userId, movieId)
}
