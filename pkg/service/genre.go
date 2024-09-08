package service

import (
	"ozinshe/pkg/repository"
	"ozinshe/schemas"
)

type GenreService struct {
	repos repository.Genre
}

func NewGenreService(repos repository.Genre) *GenreService {
	return &GenreService{repos: repos}
}

func (s *GenreService) GetAllGenre() ([]schemas.Genre, error) {
	return s.repos.GetAllGenre()
}

func (s *GenreService) AddGenre(genre string) (int, error) {
	return s.repos.AddGenre(genre)
}

func (s *GenreService) UpdateGenre(genreId int, genre string) error {
	return s.repos.UpdateGenre(genreId, genre)
}

func (s *GenreService) DeleteGenre(genreId int) error {
	return s.repos.DeleteGenre(genreId)
}
