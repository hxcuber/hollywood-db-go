package routes

import (
	"github.com/gorilla/mux"
	"hollywood/controllers"
	"net/http"
)

func RegisterRoutes(r *mux.Router) {

	a := r.PathPrefix("/actors").Subrouter()
	a.HandleFunc("", controllers.GetActors).Methods(http.MethodGet)
	a.HandleFunc("/", controllers.GetActors).Methods(http.MethodGet)
	a.HandleFunc("/{id}", controllers.GetActor).Methods(http.MethodGet)
	a.HandleFunc("/{id}/movies", controllers.GetMovieByActorID).Methods(http.MethodGet)

	m := r.PathPrefix("/movies").Subrouter()
	m.HandleFunc("", controllers.GetMovies).Methods(http.MethodGet)
	m.HandleFunc("/", controllers.GetMovies).Methods(http.MethodGet)
	m.HandleFunc("/{id}", controllers.GetMovie).Methods(http.MethodGet)

	http.Handle("/", r)
}
