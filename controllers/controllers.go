package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hxcuber/hollywood-db-go/connect"
	"github.com/hxcuber/hollywood-db-go/models"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

func Init() {
	db = connect.Connect()
}

func GetActors(w http.ResponseWriter, r *http.Request) {
	var actors models.Actors
	result := db.Table("actors").Find(&actors)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprint(w, actors)
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(actors)
	fmt.Fprintf(w, "%d rows found in actors.\n", result.RowsAffected)
	return
}

func GetActor(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor
	vars := mux.Vars(r)
	result := db.Table("actors").First(&actor, vars["id"])
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprint(w, actor)
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(actor)
	fmt.Fprintf(w, "%d rows found in actors.\n", result.RowsAffected)
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	var movies models.Movies
	result := db.Table("movies").Find(&movies)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprint(w, movies)
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(movies)
	fmt.Fprintf(w, "%d rows found in movies.\n", result.RowsAffected)
	return
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	vars := mux.Vars(r)
	result := db.Table("movies").First(&movie, vars["id"])
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprint(w, movie)
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(movie)
	fmt.Fprintf(w, "%d rows found in movies.\n", result.RowsAffected)
}

func GetMovieByActorID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var actor models.Actor
	result := db.Table("actors").First(&actor, vars["id"])
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%s %s has acted in the following movies:\n", actor.FirstName, actor.LastName)

	var movies models.Movies
	result = db.Table("movies").Where("main_actor_id = ?", vars["id"]).Find(&movies)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprint(w, movies)
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(movies)
	fmt.Fprintf(w, "%d rows found in movies.\n", result.RowsAffected)

}
