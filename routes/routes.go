package routes

import (
	"github.com/gorilla/mux"
	"github.com/hxcuber/hollywood-db-go/controllers"
	"net/http"
)

func RegisterRoutes(r *mux.Router) {

	a := r.PathPrefix("/actors").Subrouter()
	a.HandleFunc("", controllers.GetActors).Methods(http.MethodGet)
	a.Handle("/", http.RedirectHandler("/actors", http.StatusPermanentRedirect)).Methods(http.MethodGet)
	a.HandleFunc("/{id}", controllers.GetActor).Methods(http.MethodGet)
	a.HandleFunc("/{id}/movies", controllers.GetMovieByActorID).Methods(http.MethodGet)

	m := r.PathPrefix("/movies").Subrouter()
	m.HandleFunc("", controllers.GetMovies).Methods(http.MethodGet)
	m.Handle("/", http.RedirectHandler("/movies", http.StatusPermanentRedirect)).Methods(http.MethodGet)
	m.HandleFunc("/{id}", controllers.GetMovie).Methods(http.MethodGet)

	http.Handle("/", r)
}
