package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/zepyrshut/tpv/internal/config"
	"github.com/zepyrshut/tpv/internal/util"
)

var appMovie *config.Application
var stMovie *config.AppStatus

func NewMovieHandler(a *config.Application, st *config.AppStatus) {
	appLounge = a
	stLounge = st
}

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		appLounge.Logger.Println(errors.New("invalid id parameter"))
		util.ErrorJSON(w, err)
		return
	}

	movie, err := appLounge.DB.OneMovie(id)
	if err != nil {
		appLounge.Logger.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		appLounge.Logger.Println(err)
		util.ErrorJSON(w, err)
		return
	}
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := appLounge.DB.AllMovies()
	if err != nil {
		appLounge.Logger.Println(err)
		util.ErrorJSON(w, err)
		return
	}

	util.WriteJSON(w, http.StatusOK, movies, "movies")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {

}

func insertMovie(w http.ResponseWriter, r *http.Request) {

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

}

func searchMovie(w http.ResponseWriter, r *http.Request) {

}
