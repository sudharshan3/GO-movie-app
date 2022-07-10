package router

import (
	"github.com/gorilla/mux"
	"github.com/sudharshan3/GO-movie-app/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkWatched).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.DeleteaMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeletemyAllMovies).Methods("DELETE")
	return router
}
