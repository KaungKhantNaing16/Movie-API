package service

import (
	"errors"
	"Movie-API/model"
	"Movie-API/repository"
)

var (
	ErrIDIsNotValid = errors.New("Id is not valid")
	ErrTitleIsNotEmpty = errors.New("Moive title can't be empty")
	ErrMovieNotFound = errors.New("The movie can't be found")
)

type IMovieService interface {
	GetMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
	CreateMovie(movie model.Movie) error
	UpdateMovie(id int, movie model.Movie) error
	DeleteMovie(id int) error
	DeleteAllMovies() error
}

type DefaultMovieService struct {
	movieRepo repository.IMovieRespository
}

func NewDefaultMovieService(mRepo repository.IMovieRespository) *DefaultMovieService {
	return &DefaultMovieService { movieRepo : mRepo}
}

func (d *DefaultMovieService) GetMovies() ([]model.Movie, error) {
	return d.movieRepo.GetMovies()
}

func (d *DefaultMovieService) GetMovie(id int) (model.Movie, error) {
	if id <= 0 {
		return model.Movie{}, ErrIDIsNotValid
	}

	movie, err := d.movieRepo.GetMovie(id)
	if err != nil {
		if errors.Is(err, repository.ErrMovieNotFound) {
			return model.Movie{}, ErrMovieNotFound
		}
	}
	return movie, nil
}

func (d *DefaultMovieService) CreateMovie(movie model.Movie) error {
	if movie.Title == "" {
		return ErrTitleIsNotEmpty
	}

	return d.movieRepo.CreateMovie(movie)
}

func (d *DefaultMovieService) UpdateMovie(id int, movie model.Movie) error {
	if id <= 0 {
		return ErrIDIsNotValid
	}

	if movie.Title == "" {
		return ErrTitleIsNotEmpty
	}

	err := d.movieRepo.UpdateMovie(id, movie) 
	if errors.Is(err, repository.ErrMovieNotFound) {
		return ErrMovieNotFound
	}

	return nil
}

func (d *DefaultMovieService) DeleteMovie(id int) error {
	if id <= 0 {
		return ErrIDIsNotValid
	}

	err := d.movieRepo.DeleteMovie(id)
	if errors.Is(err, repository.ErrMovieNotFound) {
		return ErrMovieNotFound
	}

	return nil
}

func (d *DefaultMovieService) DeleteAllMovies() error {
	return d.movieRepo.DeleteAllMovies()
}