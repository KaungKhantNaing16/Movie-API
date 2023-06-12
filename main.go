package main

import (
	"Movie-API/handler"
	"Movie-API/repository"
	"Movie-API/service"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	movieInMemoryRepository := repository.NewInMemoryMovieRepository()
	movieService := service.NewDefaultMovieService(movieInMemoryRepository)
	movieHandler := handler.NewMovieHandler(movieService)
	router := httprouter.New()

	// Define routes
	router.GET("/movies", movieHandler.GetMovies)
	router.GET("/movies/:id", movieHandler.GetMovie)
	router.POST("/movies", movieHandler.CreateMovie)
	router.PATCH("/movies/:id", movieHandler.UpdateMovie)
	router.DELETE("/movies/:id", movieHandler.DeleteMovie)
	router.DELETE("/movies", movieHandler.DeleteAllMovies)

	log.Println("Http server runs on :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}