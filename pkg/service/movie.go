package service

import (
	"ozinshe/pkg/repository"
	"ozinshe/pkg/utils"
	"ozinshe/schemas"
)

type MovieService struct {
	repos repository.Movie
}

func NewMovieService(repos repository.Movie) *MovieService {
	return &MovieService{repos: repos}
}

func (s *MovieService) GetAll(paginParams *utils.Pagination) (*utils.PaginationMovieResponse, error) {
	offset := (paginParams.PageNum - 1) * paginParams.PerPage

	if paginParams.Total == 0 {
		count, err := s.repos.GetCount()
		if err != nil {
			return nil, err
		}
		paginParams.Total = count
	}

	movies, err := s.repos.GetAll(paginParams.PerPage, offset)
	if err != nil {
		return nil, err
	}

	res := &utils.PaginationMovieResponse{
		Total:      paginParams.Total,
		PerPage:    paginParams.PerPage,
		Pagination: paginParams.Paginate(),
		Data:       movies,
	}

	return res, nil
}

func (s *MovieService) SearchByName(serchInput string, paginParams *utils.Pagination) (*utils.PaginationMovieResponse, error) {
	offset := (paginParams.PageNum - 1) * paginParams.PerPage

	if paginParams.Total == 0 {
		count, err := s.repos.GetCountByName(serchInput)
		if err != nil {
			return nil, err
		}
		paginParams.Total = count
	}

	movies, err := s.repos.SearchByName(serchInput, paginParams.PerPage, offset)
	if err != nil {
		return nil, err
	}

	res := &utils.PaginationMovieResponse{
		Total:      paginParams.Total,
		PerPage:    paginParams.PerPage,
		Pagination: paginParams.Paginate(),
		Data:       movies,
	}

	return res, nil
}

func (s *MovieService) SearchByGenre(genre string, paginParams *utils.Pagination) (*utils.PaginationMovieResponse, error) {
	offset := (paginParams.PageNum - 1) * paginParams.PerPage

	if paginParams.Total == 0 {
		count, err := s.repos.GetCountByGenre(genre)
		if err != nil {
			return nil, err
		}
		paginParams.Total = count
	}

	movies, err := s.repos.SearchByGenre(genre, paginParams.PerPage, offset)
	if err != nil {
		return nil, err
	}

	res := &utils.PaginationMovieResponse{
		Total:      paginParams.Total,
		PerPage:    paginParams.PerPage,
		Pagination: paginParams.Paginate(),
		Data:       movies,
	}

	return res, nil
}

func (s *MovieService) AddMovie(movie schemas.AddMovieInfo) (int, error) {
	return s.repos.AddMovie(movie)
}

func (s *MovieService) GetById(movieId int) (schemas.Movie, error) {
	return s.repos.GetById(movieId)
}

func (s *MovieService) UpdateMovie(movieId int, movieInput schemas.UpdateMovieInfo) error {
	return s.repos.UpdateMovie(movieId, movieInput)
}

func (s *MovieService) DeleteMovie(movieId int) error {
	return s.repos.DeleteMovie(movieId)
}

func (s *MovieService) GetScreenshots(movieId int) ([]schemas.Screenshot, error) {
	return s.repos.GetScreenshots(movieId)
}

func (s *MovieService) AddScreenshot(movieId int, link string) (int, error) {
	return s.repos.AddScreenshot(movieId, link)
}

func (s *MovieService) DeleteScreenshot(movieId, screenId int) (string, error) {
	return s.repos.DeleteScreenshot(movieId, screenId)
}

func (s *MovieService) UploadMovie(movieId int, link string) error {
	return s.repos.UploadMovie(movieId, link)
}
