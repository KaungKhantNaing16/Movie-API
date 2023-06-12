package handler

import (
	"encoding/json"
	"errors"
	"Movie-API/model"
	"Movie-API/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"log"
)

type movieHandler struct {
	service service.IMovieService
}

/*
 *	Get all movies list
 *	localhost:8080/movies
 */
func NewMovieHandler(ms service.IMovieService) *movieHandler {
	return &movieHandler{ service:ms }
}

/*
 *	Get movie by id
 *	localhost:8080/movies/:id
 */
func (mh *movieHandler) GetMovies(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	movies, err := mh.service.GetMovies()
	if err != nil {
		http.Error(writer, "Unable to get all movies", http.StatusInternalServerError)
		return
	}

	jsonStr, err := json.Marshal(movies)
	if err != nil {
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Conent-Type", "application/json")
	writer.Write(jsonStr)
}

/*
 *	Get movie by id
 *	localhost:8080/movies/:id
 */
 func (mh *movieHandler) GetMovie(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	movie, err := mh.service.GetMovie(id)

	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		} else if errors.Is(err, service.ErrMovieNotFound) {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonStr, err := json.Marshal(movie)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonStr)
}

/*
 *	Create movie
 *	localhost:8080/movies
 *	-d { "title": "SampelMovie" }
 */
 func (mh *movieHandler) CreateMovie(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var movie model.Movie
	err := json.NewDecoder(request.Body).Decode(&movie)
	log.Println(err)
	if err != nil {
		http.Error(writer, "Error when decoding json", http.StatusInternalServerError)
		return
	}

	err = mh.service.CreateMovie(movie)
	if err != nil {
		if errors.Is(err, service.ErrTitleIsNotEmpty) {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Movie is successfully created."))
}

/*
 *	Update movie by id
 *	localhost:8080/movies/:id
 *	-d { "title": "SampelMovie" }
 */
func (mh *movieHandler) UpdateMovie(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	var movie model.Movie
	err := json.NewDecoder(request.Body).Decode(&movie)
	if err != nil {
		http.Error(writer, "Error when decoding json", http.StatusInternalServerError)
		return
	}

	err = mh.service.UpdateMovie(id, movie)
	log.Println(err)
	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) ||
			errors.Is(err, service.ErrTitleIsNotEmpty) {
				http.Error(writer, err.Error(), http.StatusBadRequest) 
				return
		} else if errors.Is(err, service.ErrMovieNotFound) {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
	writer.Write([]byte("Successfully Updated"))
}

/*
 *	Delete movie by id
 *	localhost:8080/movies/:id
 */
func (mh *movieHandler) DeleteMovie(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	err := mh.service.DeleteMovie(id)

	if err != nil {
		if errors.Is(err, service.ErrIDIsNotValid) || errors.Is(err, service.ErrMovieNotFound) {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
	writer.Write([]byte("Movie has been successfully deleted."))
}

/*
 *	Delete movies list
 *	localhost:8080/movies
 */
func (mh *movieHandler) DeleteAllMovies(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	err := mh.service.DeleteAllMovies()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
	writer.Write([]byte("Movies successfully deleted."))

}
