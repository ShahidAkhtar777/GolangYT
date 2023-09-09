package router

import (
	"github.com/GolangYT/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")
	r.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")

	return r
}
