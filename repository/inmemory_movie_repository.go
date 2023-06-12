package repository

import (
	"errors"
	"Movie-API/model"
)

var ErrMovieNotFound = errors.New("FromRepository - movie not found")

type IMovieRespository interface {
	GetMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
	CreateMovie(movie model.Movie) error
	UpdateMovie(id int, movie model.Movie) error
	DeleteMovie(id int) error
	DeleteAllMovies() error
}

type inmemoryMovieRepository struct {
	Movies []model.Movie
}

func NewInMemoryMovieRepository() *inmemoryMovieRepository {
	var movies = []model.Movie{
		{ID:1, Title:"The Shawshank Redemption", ReleaseYear: 1994, Score:9.3},
		{ID:2, Title:"The Godfather", ReleaseYear: 1972, Score:9.2},
		{ID:3, Title:"The Dark Night", ReleaseYear: 2008, Score:9.0},
	}

	return &inmemoryMovieRepository{ Movies : movies }
}

func (i *inmemoryMovieRepository) GetMovies() ([]model.Movie, error) {
	return i.Movies, nil
}

func (i *inmemoryMovieRepository) GetMovie(id int) (model.Movie, error) {
	for _, movie := range i.Movies {
		if movie.ID == id {
			return movie, nil
		}
	}

	return model.Movie{}, ErrMovieNotFound
}

func (i *inmemoryMovieRepository) CreateMovie(movie model.Movie) error {
	movie.ID = len(i.Movies) + 1
	i.Movies = append(i.Movies, movie)

	return nil
}

func (i *inmemoryMovieRepository) UpdateMovie(id int, movie model.Movie) error {
	for k := 0; k < len(i.Movies); k++ {
		if i.Movies[k].ID == id {
			i.Movies[k].Title = movie.Title
			i.Movies[k].ReleaseYear = movie.ReleaseYear
			i.Movies[k].Score = movie.Score
			return nil
		}
	}

	return ErrMovieNotFound
}

func (i *inmemoryMovieRepository) DeleteMovie(id int) error {
	movieExit := false

	var newMovieList []model.Movie
	for _, movie := range i.Movies {
		if movie.ID == id {
			movieExit = true
		} else {
			newMovieList = append(newMovieList, movie)
		}
	}

	if !movieExit {
		return ErrMovieNotFound
	}

	i.Movies = newMovieList

	return nil
}

func (i *inmemoryMovieRepository) DeleteAllMovies() error {
	i.Movies = nil
	return nil
}